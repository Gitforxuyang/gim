package handler

import (
	"fmt"
	err2 "gim/infra/err"
	"gim/infra/utils"
	"gim/proto"
	"gim/server"
	"github.com/go-redis/redis/v7"
	"github.com/gogo/protobuf/proto"
	"strconv"
	"time"
)

const (
	//在线状态2分半就过期
	ONLINE_TIME_OUT = 150
)

//设置用户在线状态 返回是否正常，如果返回false，则代表用户状态已经失效了，需要关闭连接
func (m *handler) expireUserOnlineStatus(uid int64) bool {
	ok, _ := m.redis.Expire(userOnlineStatusKey(uid), ONLINE_TIME_OUT).Result()
	return ok
}

func userOnlineStatusKey(uid int64) string {
	return fmt.Sprintf("u:%d", uid)
}

//保存用户在线状态
func (m *handler) saveUserOnlineStatus(uid int64, uuid int64) error {
	return m.redis.HSet(userOnlineStatusKey(uid), "node", m.node, "uuid", strconv.Itoa(int(uuid))).Err()
}
func (m *handler) getUserOnlineStatus(uid int64) (string, int64, error) {
	data, err := m.redis.HGetAll(userOnlineStatusKey(uid)).Result()
	if err != nil {
		return "", 0, err
	}
	if len(data) == 0 {
		return "", 0, nil
	}
	uuid, err := strconv.ParseInt(data["uuid"], 10, 64)
	if err != nil {
		return "", 0, err
	}
	return data["node"], uuid, nil
}

//保存本地在线状态
func (m *handler) saveLocalOnlineStatus(conn server.Conn, uid int64, uuid int64) error {
	conn.SetUid(uid)
	m.authConnections.Store(uid, conn)
	m.waitAuthConnections.Delete(conn.GetRemoteAddr())
	return nil
}

//清理用户在线状态
func (m *handler) clearUserOnlineStatus(uid, uuid int64) error {
	str, err := m.redis.HGet(userOnlineStatusKey(uid), "uuid").Result()
	if err != nil {
		return err
	}
	if str == strconv.Itoa(int(uuid)) {
		return m.redis.Del(userOnlineStatusKey(uid)).Err()
	}
	return nil
}

func (m *handler) clearLocalOnlineStatus(uid, uuid int64) error {
	value, ok := m.authConnections.Load(uid)
	//如果没有这个用户的连接，则直接返回
	if !ok {
		return nil
	}
	conn, _ := value.(server.Conn)
	if conn.GetUUID() == uuid {
		m.authConnections.Delete(uid)
		m.waitAuthConnections.Store(uid, conn)
		conn.Clear()
	}
	return nil
}

type mapping struct {
	msg     func() proto.Message
	handler handlerFunc
	cmdId   uint8
}

func (m *handler) handleMapping(cmdId uint8) *mapping {
	_mapping := &mapping{}
	switch cmdId {
	case server.CmdId_Ping:
		_mapping.msg = func() proto.Message {
			return &gim.Ping{}
		}
		_mapping.handler = m.ping
		_mapping.cmdId = server.CmdId_Pong
	case server.CmdId_AuthReq:
		_mapping.msg = func() proto.Message {
			return &gim.AuthReq{}
		}
		_mapping.cmdId = server.CmdId_AuthResp
		_mapping.handler = m.auth
	case server.CmdId_SendMessageReq:
		_mapping.msg = func() proto.Message {
			return &gim.SendMessageReq{}
		}
		_mapping.cmdId = server.CmdId_SendMessageResp
		_mapping.handler = auth(m.sendMsg)
	}

	return _mapping
}

type redlock struct {
	value string
	uid   int64
	redis *redis.Client
}

var (
	releaseScript = redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end`)
)

func _lockKey(uid int64) string {
	return fmt.Sprintf("lock:gim:u:%d", uid)
}
func (m *handler) lock(uid int64) (*redlock, error) {
	value := strconv.Itoa(int(utils.GetSnowflakeId()))
	success, err := m.redis.SetNX(_lockKey(uid), value, time.Second*3).Result()
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, err2.LockExistsError
	}
	return &redlock{value: value, uid: uid, redis: m.redis}, nil
}

func (m *redlock) unlock() {
	releaseScript.Eval(m.redis, []string{_lockKey(m.uid)}, m.value)
}
