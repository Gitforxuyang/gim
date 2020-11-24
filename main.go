package main

import (
	"fmt"
	"gim/conf"
	"gim/handler"
	redis2 "gim/infra/redis"
	"gim/server/gnet"
	"gim/server/ws"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := conf.InitConfig()
	redis := redis2.InitClient(config)
	handle := handler.NewHandler(redis)
	tcpServer := gnet.NewGNetServer(9003, handle)
	go tcpServer.Run()
	wsServer := ws.NewWsServer(handle, 9004)
	go wsServer.Run()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	fmt.Println("接收到信号关闭:", s)
}
