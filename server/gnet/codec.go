package gnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"gim/server"
	"github.com/panjf2000/gnet"
	"github.com/sirupsen/logrus"
)

type codec struct {
}

func (m *codec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	gim := server.ByteToGim(buf)
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
			logrus.Errorln(err)
		}
	}()
	len, buf := c.ReadN(server.HEAD_LEN)
	if len == 0 || len != server.HEAD_LEN {
		err = errors.New("没有更多数据了")
		return
	}
	headBuf := bytes.NewBuffer(buf)
	binary.Read(headBuf, binary.BigEndian, &g.Version)
	binary.Read(headBuf, binary.BigEndian, &g.CmdId)
	binary.Read(headBuf, binary.BigEndian, &g.BodyLen)
	if !server.IsCorrectCmdId(g.CmdId) {
		c.ResetBuffer()
		err = errors.New("错误的cmdId")
		return
	}
	if g.BodyLen > server.MAX_BODY_LEN {
		c.ResetBuffer()
		err = errors.New("消息超过最大体积")
		return
	}
	protocolLen := int(g.BodyLen + server.HEAD_LEN)
	if g.BodyLen > 0 {
		len, data := c.ReadN(protocolLen)
		if len != protocolLen {
			err = errors.New("body数据不全")
			return
		}
		g.Data = data[server.HEAD_LEN:]
	}
	c.ShiftN(protocolLen)
	protocolBuf = server.GimToByte(g)
	return protocolBuf, nil
}
