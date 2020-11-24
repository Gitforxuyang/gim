package handler

import (
	"errors"
	"gim/server"
	"gim/utils"
)

func (m *handler) ping(conn server.Conn, gim *server.GimProtocol) (res *server.GimProtocol, err error) {
	conn.SetPingAt(utils.NowMillisecond())
	if conn.GetUid() != 0 {
		ok := m.expireUserOnlineStatus(conn.GetUid())
		if !ok {
			m.Close(conn)
			err = errors.New("在ping时发现用户状态已过期，关闭连接")
			return
		}
	}
	res = server.MakePong(conn)
	return res, nil
}
