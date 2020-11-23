package handler

import (
	"errors"
	"fmt"
	"gim/server"
	"gim/utils"
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
}

func NewHandler() IHandler {
	h := handler{retryList: utils.NewRetryList()}
	h.retrySend()
	h.checkConnectionActive()
	return &h
}

func (m *handler) Open(conn server.Conn) error {
	return nil
}

func (m *handler) Close(conn server.Conn) error {
	if conn.GetUid() == 0 {
		m.waitAuthConnections.Delete(conn.GetRemoteAddr())
		//从redis中删除状态
	} else {
		m.authConnections.Delete(conn.GetUid())
	}
	return nil
}

func (m *handler) Shutdown() {
}

func (m *handler) Action(conn server.Conn, gim *server.GimProtocol) (res *server.GimProtocol, err error) {
	utils.PrintStrcut(gim)
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("action panic err: %s", e))
		}
		if err != nil {
			fmt.Println(err)
		}
	}()
	switch gim.CmdId {
	case server.CmdId_Ping:
		conn.SetPingAt(utils.NowMillisecond())
		res = server.MakePong(conn)
	case server.CmdId_AuthReq:
	case server.CmdId_LogoutReq:
	case server.CmdId_FetchMessageReq:
	case server.CmdId_NotifyAck:

	case server.CmdId_SendMessageReq:
	case server.CmdId_SyncLastIdReq:
	case server.CmdId_SyncMessageReq:
	default:
		err = errors.New("非法cmdId")
		return
	}
	return
}

func (m *handler) retrySend() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("retrySend panic:", e)
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
				m.Close(conn)
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
				m.Close(conn)
			}
			return true
		})
	}
}
