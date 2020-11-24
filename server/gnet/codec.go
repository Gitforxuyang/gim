package gnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"gim/server"
	"github.com/panjf2000/gnet"
)

const (
	HEAD_LEN = 4
	//单个请求消息最大6k
	MAX_BODY_LEN = 6000
)

type codec struct {
}

func (m *codec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	gim := byteToGim(buf)
	buffer := &bytes.Buffer{}
	if err := binary.Write(buffer, binary.BigEndian, gim.Version); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gim.CmdId); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gim.BodyLen); err != nil {
		return nil, err
	}
	if gim.BodyLen > 0 {
		if err := binary.Write(buffer, binary.BigEndian, gim.Data); err != nil {
			return nil, err
		}
	}
	return buffer.Bytes(), nil
}

func (m *codec) Decode(c gnet.Conn) (protocolBuf []byte, err error) {
	g := &server.GimProtocol{}
	defer func() {
		if err != nil && err.Error() != "没有更多数据了" {
			fmt.Println(err)
		}
	}()
	len, buf := c.ReadN(HEAD_LEN)
	if len == 0 || len != HEAD_LEN {
		err = errors.New("没有更多数据了")
		return
	}
	headBuf := bytes.NewBuffer(buf)
	binary.Read(headBuf, binary.BigEndian, &g.Version)
	binary.Read(headBuf, binary.BigEndian, &g.CmdId)
	binary.Read(headBuf, binary.BigEndian, &g.BodyLen)
	if !isCorrectCmdId(g.CmdId) {
		c.ResetBuffer()
		err = errors.New("错误的cmdId")
		return
	}
	if g.BodyLen > MAX_BODY_LEN {
		c.ResetBuffer()
		err = errors.New("消息超过最大体积")
		return
	}
	protocolLen := int(g.BodyLen + HEAD_LEN)
	if g.BodyLen > 0 {
		len, data := c.ReadN(protocolLen)
		if len != protocolLen {
			err = errors.New("body数据不全")
			return
		}
		g.Data = data[HEAD_LEN:]
	}
	c.ShiftN(protocolLen)
	protocolBuf = gimToByte(g)
	return protocolBuf, nil
}

func isCorrectCmdId(cmdId uint8) bool {
	if cmdId == 1 {
		return true
	}
	return false
}

func gimToByte(g *server.GimProtocol) ([]byte) {
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

func byteToGim(buf []byte) (*server.GimProtocol) {
	g := &server.GimProtocol{}
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
