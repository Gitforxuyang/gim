package server

//tcp跟websocket的connection的定义
type Conn interface {
	SetPingAt(t int64)
	GetPingAt() int64
	Close()
	GetRemoteAddr() string
	SetUid(uid int64)
	GetUid() int64
}
