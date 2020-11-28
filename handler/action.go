package handler

import (
	"errors"
	err2 "gim/infra/err"
	"gim/infra/utils"
	"gim/proto"
	"gim/server"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
)

func (m *handler) ping(conn server.Conn, message proto.Message) (res proto.Message, err error) {
	//req := message.(*gim.Ping)
	logrus.Debugln("ping uuid:", conn.GetUUID(), "uid:", conn.GetUid())
	conn.SetPingAt(utils.NowMillisecond())
	if conn.GetUid() != 0 {
		ok := m.expireUserOnlineStatus(conn.GetUid())
		if !ok {
			m.Close(conn)
			err = errors.New("在ping时发现用户状态已过期，关闭连接")
			return nil, err
		}
	}
	res = &gim.Pong{}
	return res, nil
}

func (m *handler) auth(conn server.Conn, msg proto.Message) (res proto.Message, err error) {
	req := msg.(*gim.AuthReq)
	logrus.Infoln("auth uuid:", conn.GetUUID(), "uid:", conn.GetUid(), "req:", req.String())
	if req.Token == "" || req.Uid == 0 {
		return nil, err2.ParamError
	}
	//登录动作
	return res, nil
}
