package test

import (
	"context"
	"encoding/hex"
	"fmt"
	"gim/proto"
	"github.com/golang/protobuf/proto"
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
func TestProto(t *testing.T) {
	//authReq := gim.AuthReq{Uid:1,Token:"1"}
	//fmt.Println(len(buf))
	req := gim.SendMessageReq{}
	req.Seq = time.Now().Unix()
	req.Content = "123"
	req.To = 2
	req.From = 1
	req.Action = 1
	req.Type = 1
	buf,_ := proto.Marshal(&req)
	fmt.Println(hex.EncodeToString(buf))
}
