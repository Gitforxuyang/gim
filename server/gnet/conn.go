package gnet

import "github.com/panjf2000/gnet"

type conn struct {
	c          gnet.Conn
	remoteAddr string
	uid        int64
	pingAt     int64
}

func (m *conn) SetPingAt(t int64) {
	m.pingAt = t
}

func (m *conn) GetPingAt() int64 {
	return m.pingAt
}

func (m *conn) Close() {
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
