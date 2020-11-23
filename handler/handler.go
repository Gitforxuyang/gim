package handler

import (
	"gim/server"
	"gim/server/gnet"
	"gim/utils"
	"sync"
)

type IHandler interface {
	Open(conn server.Conn) error
	Close(conn server.Conn) error
	Shutdown()
	Action(conn server.Conn, gim *gnet.GimProtocol) (*gnet.GimProtocol, error)
}

type handler struct {
	authConnections     sync.Map
	waitAuthConnections sync.Map
	retryList           *utils.RetryList
}

func NewHandler() IHandler {
	h := handler{}
	return &h
}

func (m *handler) Open(conn server.Conn) error {
}

func (m *handler) Close(conn server.Conn) error {
	panic("implement me")
}

func (m *handler) Shutdown() {
	panic("implement me")
}

func (m *handler) Action(conn server.Conn, gim *gnet.GimProtocol) (*gnet.GimProtocol, error) {
	panic("implement me")
}
