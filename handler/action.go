package handler

import (
	"errors"
	"gim/proto"
	"gim/server"
	"gim/infra/utils"
	"github.com/gogo/protobuf/proto"
)

func (m *handler) ping(conn server.Conn, message proto.Message) (res proto.Message, err error) {
	//req := message.(*gim.Ping)

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
