package server

//tcp跟websocket的connection的定义
type Conn interface {
	SetPingAt(t int64)
	GetPingAt() int64
	Close()
	GetRemoteAddr() string
	SetUid(uid int64)
	GetUid() int64
	GetVersion() uint8
	SetVersion(version uint8)
	SetToken(token string)
	GetToken() string
	Write(gim *GimProtocol) error
}
