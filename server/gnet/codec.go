package gnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/panjf2000/gnet"
)

type GimProtocol struct {
	Version uint8
	CmdId   uint8
	BodyLen uint16
	Data    []byte
}

const (
	HEAD_LEN = 4
)

func (m *GimProtocol) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
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

func (m *GimProtocol) Decode(c gnet.Conn) ([]byte, error) {
	fmt.Println("decode")
	g := &GimProtocol{}
	len, buf := c.ReadN(HEAD_LEN)
	if len == 0 || len != HEAD_LEN {
		return nil, errors.New("没有更多数据了")
	}
	headBuf := bytes.NewBuffer(buf)
	binary.Read(headBuf, binary.BigEndian, &g.Version)
	binary.Read(headBuf, binary.BigEndian, &g.CmdId)
	binary.Read(headBuf, binary.BigEndian, &g.BodyLen)
	if !isCorrectCmdId(g.CmdId) {
		c.ResetBuffer()
		return nil, errors.New("错误的cmdId")
	}
	protocolLen := int(g.BodyLen + HEAD_LEN)
	if g.BodyLen > 0 {
		len, data := c.ReadN(protocolLen)
		if len != protocolLen {
			return nil, errors.New("body数据不全")
		}
		g.Data = data[HEAD_LEN:]
	}
	c.ShiftN(protocolLen)
	protocolBuf := gimToByte(g)
	return protocolBuf, nil
}

func isCorrectCmdId(cmdId uint8) bool {
	if cmdId == 1 {
		return true
	}
	return false
}

func gimToByte(g *GimProtocol) ([]byte) {
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

func byteToGim(buf []byte) (*GimProtocol) {
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
