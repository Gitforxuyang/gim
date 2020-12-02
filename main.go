package main

import (
	"fmt"
	"gim/conf"
	"gim/handler"
	"gim/infra/grpc"
	"gim/infra/rabbit"
	redis2 "gim/infra/redis"
	"gim/server/gnet"
	"gim/server/ws"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := conf.InitConfig()
	level, _ := logrus.ParseLevel(config.LogLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.999"})
	logrus.SetLevel(level)
	redis := redis2.InitClient(config)
	imClient := grpc.InitClient(config)
	queue := rabbit.InitClient(config)
	handle := handler.NewHandler(redis, imClient,queue)
	tcpServer := gnet.NewGNetServer(9003, handle)
	go tcpServer.Run()
	wsServer := ws.NewWsServer(handle, 9004)
	go wsServer.Run()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	fmt.Println("接收到信号关闭:", s)
	handle.Shutdown()
	fmt.Println("关闭服务完成")
}
