package handler

import (
	"fmt"
	"gim/proto"
	"gim/server"
	"github.com/gogo/protobuf/proto"
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
	}
	return _mapping
}
