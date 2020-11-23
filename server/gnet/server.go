package gnet

import (
	"fmt"
	"gim/server"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"time"
)

type gnetServer struct {
	port     int32
	workPool *goroutine.Pool
}

func (m *gnetServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	fmt.Println("初始化完成")
	return 0
}

func (m *gnetServer) OnShutdown(server gnet.Server) {
	fmt.Println("服务关闭")
}

func (m *gnetServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("连接建立：", c.RemoteAddr())
	return nil, 0
}

func (m *gnetServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println("连接关闭：", c.RemoteAddr())
	return 0
}

func (m *gnetServer) PreWrite() {
}

func (m *gnetServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	//gim := byteToGim(frame)
	m.workPool.Submit(func() {
		c.AsyncWrite(frame)
	})
	return
}

func (m *gnetServer) Tick() (delay time.Duration, action gnet.Action) {
	return time.Second * 100, 0
}

func NewGNetServer(port int32) server.Server {
	s := &gnetServer{port: port, workPool: goroutine.Default()}
	return s
}

func (m *gnetServer) Run() {
	gnet.Serve(m,
		fmt.Sprintf("tcp://0.0.0.0:%d", m.port),
		gnet.WithMulticore(true),
		gnet.WithTCPKeepAlive(time.Minute*5),
		gnet.WithCodec(&GimProtocol{}),
		gnet.WithTicker(false),
	)
}
