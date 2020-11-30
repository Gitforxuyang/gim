package gnet

import (
	"fmt"
	"gim/handler"
	"gim/server"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/sirupsen/logrus"
	"time"
)

type gnetServer struct {
	port     int32
	workPool *goroutine.Pool
	handler  handler.IHandler
}

func (m *gnetServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	logrus.Infoln("初始化完成")
	return 0
}

func (m *gnetServer) OnShutdown(server gnet.Server) {
	//logrus.Errorln("服务开始关闭")
	//m.handler.Shutdown()
	//logrus.Errorln("服务关闭完成")
}

func (m *gnetServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	//logrus.Errorln("连接建立：", c.RemoteAddr())
	con := conn{remoteAddr: c.RemoteAddr().String(), c: c}
	err := m.handler.Open(&con)
	if err != nil {
		logrus.Errorln("gnet因为连接建立错误，关闭连接 :", con.remoteAddr)
		return nil, gnet.Close
	}
	c.SetContext(&con)
	return nil, 0
}

func (m *gnetServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	//logrus.Errorln("连接关闭：", c.RemoteAddr(), "err:", err)
	con := c.Context().(server.Conn)
	err = m.handler.Close(con)
	if err != nil {
		logrus.Errorln("关闭时报错: ", err)
	}
	return 0
}

func (m *gnetServer) PreWrite() {
}

func (m *gnetServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	gim := server.ByteToGim(frame)
	con := c.Context().(server.Conn)
	con.SetVersion(gim.Version)
	m.workPool.Submit(func() {
		m.handler.Action(con, gim)
		//if err != nil {
		//	logrus.Errorln("handler出现error，直接忽略:", err)
		//} else {
		//	c.AsyncWrite(server.GimToByte(res))
		//}
	})
	return
}

func (m *gnetServer) Tick() (delay time.Duration, action gnet.Action) {
	return time.Second * 100, 0
}

func NewGNetServer(port int32, handler handler.IHandler) server.Server {
	s := &gnetServer{port: port, workPool: goroutine.Default(), handler: handler}
	return s
}

func (m *gnetServer) Run() {
	logrus.Infoln("gnet server run")
	err := gnet.Serve(m,
		fmt.Sprintf("tcp://0.0.0.0:%d", m.port),
		gnet.WithMulticore(true),
		gnet.WithTCPKeepAlive(time.Minute*5),
		gnet.WithCodec(&codec{}),
		gnet.WithTicker(false),
	)
	if err != nil {
		panic(err)
	}
}
