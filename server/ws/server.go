package ws

import (
	"fmt"
	"gim/handler"
	"gim/server"
	"github.com/gorilla/websocket"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/sirupsen/logrus"
	"net/http"
)

type wsServer struct {
	port     int32
	handler  handler.IHandler
	workPool *goroutine.Pool
}

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (m *wsServer) echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	conn := &conn{c: c, remoteAddr: c.RemoteAddr().String()}
	defer m.handler.Close(conn)
	err = m.handler.Open(conn)
	if err != nil {
		logrus.Errorln("因为连接建立错误，关闭连接 :", conn.remoteAddr)
		conn.Close()
		return
	}
	for {
		_, data, err := c.ReadMessage()
		if err != nil {
			logrus.Errorln("读取消息是出错:", err)
			continue
		}
		gim := server.ByteToGim(data)
		if !server.IsCorrectCmdId(gim.CmdId) {
			logrus.Errorln("非法的cmdId")
			continue
		}
		if gim.BodyLen > server.MAX_BODY_LEN {
			logrus.Errorln("超过body体最大范围")
			continue
		}
		conn.SetVersion(gim.Version)
		m.workPool.Submit(func() {
			m.handler.Action(conn, gim)
		})
		//if err != nil {
		//	logrus.Errorln("handler出现error，直接忽略:", err)
		//} else {
		//	conn.Write(res)
		//}

	}
}

func (m *wsServer) Run() {
	logrus.Infoln("ws server run")
	http.HandleFunc("/echo", m.echo)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", m.port), nil)
}

func NewWsServer(handler handler.IHandler, port int32) server.Server {
	s := wsServer{handler: handler, port: port, workPool: goroutine.Default()}
	return &s
}
