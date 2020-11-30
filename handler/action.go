package handler

import (
	"context"
	"errors"
	err2 "gim/infra/err"
	"gim/infra/utils"
	"gim/proto"
	"gim/proto/im"
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
	ctx := context.TODO()
	//登录动作
	_, err = m.imClient.Auth(ctx, &im.AuthReq{Token: req.Token, Uid: req.Uid, Uuid: conn.GetUUID()})
	if err != nil {
		return nil, err
	}
	lock, err := m.lock(req.Uid)
	if err != nil {
		return nil, err
	}
	defer lock.unlock()
	node, uuid, err := m.getUserOnlineStatus(req.Uid)
	if err != nil {
		return nil, err
	}
	//如果这个有已经登陆的节点,则要踢出老节点
	if node != "" {
		if node == m.node {
			m.clearLocalOnlineStatus(req.Uid, uuid)
		}
		err = m.clearUserOnlineStatus(req.Uid, uuid)
		if err != nil {
			return nil, err
		}
	}
	err = m.saveLocalOnlineStatus(conn, req.Uid, conn.GetUUID())
	if err != nil {
		return nil, err
	}
	err = m.saveUserOnlineStatus(req.Uid, conn.GetUUID())
	if err != nil {
		return nil, err
	}
	return &gim.AuthRes{}, nil
}

func (m *handler) logout(conn server.Conn, msg proto.Message) (res proto.Message, err error) {
	req := msg.(*gim.LogoutReq)
	logrus.Infoln("logout uuid:", conn.GetUUID(), "uid:", conn.GetUid(), "req:", req.String())
	return res, nil
}

func (m *handler) sendMsg(conn server.Conn, msg proto.Message) (res proto.Message, err error) {
	req := msg.(*gim.SendMessageReq)
	ctx := context.TODO()
	logrus.Infoln("sendMsg uuid:", conn.GetUUID(), "uid:", conn.GetUid(), "req:", req.String())
	if req.Seq == 0 || req.To == 0 || req.From == 0 || req.Type == 0 || req.Action == 0 {
		return nil, err2.ParamError
	}
	data, err := m.imClient.SendMsg(ctx,
		&im.SendMsgReq{Seq: req.Seq, Type: req.Type, Action: im.MessageAction(req.Action), From: req.From, To: req.To, Content: req.Content})
	if err != nil {
		return nil, err
	}
	return &gim.SendMessageResp{Seq: data.Seq, MsgId: data.MsgId}, nil
}

func (m *handler) notifyAck(conn server.Conn, msg proto.Message) (res proto.Message, err error) {
	req := msg.(*gim.NotifyAck)
	logrus.Infoln("notifyAck uuid:", conn.GetUUID(), "uid:", conn.GetUid(), "req:", req.String())

	//登录动作
	return res, nil
}

func (m *handler) syncMsg(conn server.Conn, msg proto.Message) (res proto.Message, err error) {
	req := msg.(*gim.SyncMessageReq)
	logrus.Infoln("syncMsg uuid:", conn.GetUUID(), "uid:", conn.GetUid(), "req:", req.String())
	//登录动作
	return res, nil
}

func (m *handler) syncClientSeqId(conn server.Conn, msg proto.Message) (res proto.Message, err error) {
	req := msg.(*gim.SyncLastIdReq)
	logrus.Infoln("syncClientSeqId uuid:", conn.GetUUID(), "uid:", conn.GetUid(), "req:", req.String())

	//登录动作
	return res, nil
}

func (m *handler) fetchMsg(conn server.Conn, msg proto.Message) (res proto.Message, err error) {
	req := msg.(*gim.FetchMessageReq)
	logrus.Infoln("fetchMsg uuid:", conn.GetUUID(), "uid:", conn.GetUid(), "req:", req.String())
	return res, nil
}

func (m *handler) _login() {

}
