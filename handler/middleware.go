package handler

import (
	err2 "gim/infra/err"
	"gim/server"
	"github.com/gogo/protobuf/proto"
)

type handlerFunc func(conn server.Conn, message proto.Message) (res proto.Message, err error)

//验证登陆态
func auth(f handlerFunc) handlerFunc {
	return func(conn server.Conn, message proto.Message) (res proto.Message, err error) {
		if conn.GetUid() == 0 {
			return nil, err2.UserNotLoginError
		}
		return f(conn, message)
	}
}
