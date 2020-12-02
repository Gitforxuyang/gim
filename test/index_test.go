package test

import (
	"context"
	"fmt"
	"gim/proto"
	"gim/proto/im"
	"gim/server"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"net"
	"os"
	"testing"
	"time"
)

var (
	ctx context.Context
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	exitCode := m.Run()
	os.Exit(exitCode)
}

//func TestProto(t *testing.T) {
//	//authReq := gim.AuthReq{Uid:1,Token:"1"}
//	//fmt.Println(len(buf))
//	req := gim.SendMessageReq{}
//	req.Seq = time.Now().Unix()
//	req.Content = "123"
//	req.To = 2
//	req.From = 1
//	req.Action = 1
//	req.Type = 1
//	buf,_ := proto.Marshal(&req)
//	fmt.Println(hex.EncodeToString(buf))
//}
//
func send(c net.Conn, cmdId uint8, msg proto.Message) {
	gim := server.GimProtocol{}
	gim.CmdId = cmdId
	gim.Version = 1
	fmt.Println(msg)
	gim.Data, _ = proto.Marshal(msg)
	gim.BodyLen = uint16(len(gim.Data))
	c.Write(server.GimToByte(&gim))
}
func TestClient(t *testing.T) {
	client, err := net.Dial("tcp", "localhost:9003")
	assert.NoError(t, err)
	authReq := gim.AuthReq{Uid: 1, Token: "1"}
	send(client, server.CmdId_AuthReq, &authReq)
	time.Sleep(time.Millisecond * 500)
	msg := im.SendMsgReq{Seq: time.Now().Unix(), Type: 1, Action: 1, From: 1, To: 1, Content: "123"}
	send(client, server.CmdId_SendMessageReq, &msg)
}
