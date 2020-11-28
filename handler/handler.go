package handler

import (
	"errors"
	"fmt"
	err2 "gim/infra/err"
	"gim/infra/utils"
	gim2 "gim/proto"
	"gim/server"
	"github.com/go-redis/redis/v7"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type IHandler interface {
	Open(conn server.Conn) error
	Close(conn server.Conn) error
	Shutdown()
	Action(conn server.Conn, gim *server.GimProtocol) (*server.GimProtocol, error)
}

const (
	//两分钟内不活跃就被踢出
	CONN_ACTIVE_TIME = 120 * 1000
)

type handler struct {
	authConnections     sync.Map
	waitAuthConnections sync.Map
	retryList           *utils.RetryList
	redis               *redis.Client
}

func NewHandler(redis *redis.Client) IHandler {
	h := handler{retryList: utils.NewRetryList()}
	go h.retrySend()
	go h.checkConnectionActive()
	h.redis = redis
	return &h
}

func (m *handler) Open(conn server.Conn) error {
	conn.SetPingAt(utils.NowMillisecond())
	conn.SetUUID(utils.GetSnowflakeId())
	logrus.Debugln("连接建立 remote:", conn.GetRemoteAddr(), "uuid:", conn.GetUUID())
	m.waitAuthConnections.Store(conn.GetRemoteAddr(), conn)
	return nil
}

func (m *handler) Close(conn server.Conn) error {
	logrus.Debugln("连接关闭 remote:", conn.GetRemoteAddr(), "uuid:", conn.GetUUID())
	if conn.GetUid() == 0 {
		m.waitAuthConnections.Delete(conn.GetRemoteAddr())
		//从redis中删除状态
	} else {
		m.authConnections.Delete(conn.GetUid())
	}
	return nil
}

func (m *handler) _close(conn server.Conn) error {
	if conn.GetUid() == 0 {
		m.waitAuthConnections.Delete(conn.GetRemoteAddr())
	} else {
		m.authConnections.Delete(conn.GetUid())
		//从redis中删除状态
		m.redis.Del(userOnlineStatusKey(conn.GetUid()))
	}
	conn.Close()
	return nil
}

//关闭服务，下线所有client
func (m *handler) Shutdown() {
	m.waitAuthConnections.Range(func(key, value interface{}) bool {
		c, ok := value.(server.Conn)
		if ok {
			m._close(c)
		}
		return true
	})
	m.authConnections.Range(func(key, value interface{}) bool {
		c, ok := value.(server.Conn)
		if ok {
			m._close(c)
		}
		return true
	})
}

func (m *handler) Action(conn server.Conn, gim *server.GimProtocol) (res *server.GimProtocol, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("action panic err: %s", e))
		}
		if err != nil {
			logrus.Errorln("action error: ", err)
		}
	}()
	logrus.Debugln("action version:", gim.Version, "cmdId:", gim.CmdId)
	_mapping := m.handleMapping(gim.CmdId)
	msg := _mapping.msg()
	err = proto.Unmarshal(gim.Data, msg)
	if err != nil {
		return nil, err
	}
	resMsg, err := _mapping.handler(conn, msg)
	if err != nil {
		logrus.Errorln("action出现错误,", err.Error())
		if e, ok := err.(*err2.GimError); ok {
			resMsg = &gim2.BaseResp{Code: e.Code, Msg: e.Msg}
		} else {
			resMsg = &gim2.BaseResp{Code: err2.UnknownError.Code, Msg: err2.UnknownError.Msg}
		}
	}

	res = &server.GimProtocol{}
	res.Version = conn.GetVersion()
	res.CmdId = _mapping.cmdId
	res.Data, err = proto.Marshal(resMsg)
	if err != nil {
		return nil, err
	}
	res.BodyLen = uint16(len(res.Data))
	return
}

func (m *handler) retrySend() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("retrySend panic:", e)
			m.retrySend()
		}
	}()
	for {
		list, err := m.retryList.GetWaitRetryMsg()
		if err != nil {
			fmt.Println(err)
			break
		}
		for _, v := range list {
			conn, ok := m.authConnections.Load(v.Uid)
			//当某些场景下，需要发送的连接已经不存在了（比如用户下线了）,则删除消息
			if !ok {
				m.retryList.RemoveRetryMsg(v.MsgId)
				continue
			}
			c, ok := conn.(server.Conn)
			if !ok {
				m.retryList.RemoveRetryMsg(v.MsgId)
				fmt.Println("连接转换错误")
				continue
			}
			err = c.Write(v.Msg)
			if err != nil {
				fmt.Println(err)
			}
		}
		//每隔100ms循环检查一次是否有需要发送的消息
		time.Sleep(time.Millisecond * 100)
	}
}

//检查连接活跃
func (m *handler) checkConnectionActive() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("CheckConnectionActive panic:", e)
			m.checkConnectionActive()
		}
	}()

	for {
		time.Sleep(time.Second)
		num := 100
		m.authConnections.Range(func(key, value interface{}) bool {
			num++
			if num > 100 {
				time.Sleep(time.Millisecond * 10)
				num = 0
			}
			conn, ok := value.(server.Conn)
			if !ok {
				m.authConnections.Delete(key)
				fmt.Println("CheckConnectionActive 连接有问题")
				return true
			}
			//如果最近一次活跃时间在规定时间前，则关闭
			if conn.GetPingAt()+CONN_ACTIVE_TIME < utils.NowMillisecond() {
				m._close(conn)
			}
			return true
		})
		m.waitAuthConnections.Range(func(key, value interface{}) bool {
			num++
			if num > 100 {
				time.Sleep(time.Millisecond * 10)
				num = 0
			}
			conn, ok := value.(server.Conn)
			if !ok {
				m.waitAuthConnections.Delete(key)
				fmt.Println("CheckConnectionActive 连接有问题")
				return true
			}
			//如果最近一次活跃时间在规定时间前，则关闭
			if conn.GetPingAt()+CONN_ACTIVE_TIME < utils.NowMillisecond() {
				m._close(conn)
			}
			return true
		})
	}
}
