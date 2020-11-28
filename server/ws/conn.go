package ws

import (
	"gim/server"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type conn struct {
	c          *websocket.Conn
	remoteAddr string
	uid        int64
	pingAt     int64
	version    uint8
	token      string
	uuid       int64
}

func (m *conn) SetUUID(uuid int64) {
	m.uuid = uuid
}

func (m *conn) GetUUID() int64 {
	return m.uuid
}

func (m *conn) SetPingAt(t int64) {
	m.pingAt = t
}

func (m *conn) GetPingAt() int64 {
	return m.pingAt
}

func (m *conn) Close() {
	logrus.Debugln("连接被server主动关闭 uuid:", m.uuid, "remote:", m.remoteAddr)
	m.c.Close()
}

func (m *conn) GetRemoteAddr() string {
	return m.remoteAddr
}

func (m *conn) SetUid(uid int64) {
	m.uid = uid
}

func (m *conn) GetUid() int64 {
	return m.uid
}

func (m *conn) GetVersion() uint8 {
	return m.version
}

func (m *conn) SetVersion(version uint8) {
	m.version = version
}

func (m *conn) SetToken(token string) {
	m.token = token
}

func (m *conn) GetToken() string {
	return m.token
}

func (m *conn) Write(gim *server.GimProtocol) error {
	return m.c.WriteMessage(1, server.GimToByte(gim))
}
