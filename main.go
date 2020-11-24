package main

import (
	"gim/conf"
	"gim/handler"
	redis2 "gim/infra/redis"
	"gim/server/gnet"
)

func main() {
	config := conf.InitConfig()
	redis := redis2.InitClient(config)
	handle := handler.NewHandler(redis)
	tcpServer := gnet.NewGNetServer(9003, handle)
	tcpServer.Run()

}
