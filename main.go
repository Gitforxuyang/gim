package main

import (
	"gim/handler"
	"gim/server/gnet"
)

func main() {
	handle := handler.NewHandler()
	tcpServer := gnet.NewGNetServer(9003, handle)
	tcpServer.Run()

}
