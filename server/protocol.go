package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type GimProtocol struct {
	Version uint8
	CmdId   uint8
	BodyLen uint16
	Data    []byte
}

const (
	//基本功能
	CmdId_Ping uint8 = 1
	CmdId_Pong uint8 = 2

	//基础项
	CmdId_AuthReq     uint8 = 21 //认证请求
	CmdId_AuthResp    uint8 = 22 //认证返回
	CmdId_LogoutReq   uint8 = 23 //退出
	CmdId_LoutoutResp uint8 = 23
	CmdId_KickOut     uint8 = 24 //踢出

	//偏业务逻辑项
	CmdId_SendMessageReq   uint8 = 101 //发送消息
	CmdId_SendMessageResp  uint8 = 102 //发送消息
	CmdId_Notify           uint8 = 103
	CmdId_NotifyAck        uint8 = 104
	CmdId_SyncMessageReq   uint8 = 105
	CmdId_SyncMessageResp  uint8 = 106
	CmdId_FetchMessageReq  uint8 = 107
	CmdId_FetchMessageResp uint8 = 108
	CmdId_SyncLastIdReq    uint8 = 109
	CmdId_SyncLastIdResp   uint8 = 110
)

const (
	HEAD_LEN = 4
	//单个请求消息最大6k
	MAX_BODY_LEN = 6000
)

//func MakePong(conn Conn) *GimProtocol {
//	gim := GimProtocol{CmdId: CmdId_Pong, Version: conn.GetVersion()}
//	return &gim
//}

func GimToByte(g *GimProtocol) ([]byte) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, g.Version)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Write(buf, binary.BigEndian, g.CmdId)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Write(buf, binary.BigEndian, g.BodyLen)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Write(buf, binary.BigEndian, g.Data)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}

func ByteToGim(buf []byte) (*GimProtocol) {
	g := &GimProtocol{}
	buffer := bytes.NewBuffer(buf)
	err := binary.Read(buffer, binary.BigEndian, &g.Version)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Read(buffer, binary.BigEndian, &g.CmdId)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Read(buffer, binary.BigEndian, &g.BodyLen)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Read(buffer, binary.BigEndian, &g.Data)
	if err != nil {
		fmt.Println(err)
	}
	return g
}
func IsCorrectCmdId(cmdId uint8) bool {
	if cmdId == 1 {
		return true
	}
	return false
}
