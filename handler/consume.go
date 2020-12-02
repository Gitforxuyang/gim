package handler

import (
	"gim/server"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func (m *handler) runConsume() {
	m.queue.InitQueue(m.node)
	m.queue.Consume(m.consume)
}

type QueueHeader struct {
	Type  int32 `json:"type"`  //消息类别  单聊 群聊 聊天室
	CmdId uint8 `json:"cmdId"` //操作码id
	MsgId int64 `json:"msgId"` //消息id
	To    int64 `json:"to"`    //发送目标
	UUID  int64 `json:"uuid"`  //如果指定连接id则不为0
}

func (m *handler) consume(msg amqp.Delivery) {
	header := QueueHeader{}
	mapstructure.Decode(msg.Headers, &header)
	switch header.CmdId {
	case server.CmdId_Notify:
		if header.Type == 2 || header.Type == 3 {
			m.groupSend(&header, msg.Body)
			return
		}
	case server.CmdId_KickOut:
	case server.CmdId_Broadcast:
		m.allSend(&header, msg.Body)
		return
	}
	//省下的都是明确指定发给谁
	conn, ok := m.authConnections.Load(header.To)
	if !ok {
		logrus.Errorln("consume 指定的接收方不存在:", header)
		return
	}
	c, _ := conn.(server.Conn)
	if header.UUID != 0 && c.GetUUID() != header.UUID {
		logrus.Errorln("consume 指定的接收方的uuid对不上:", header)
		return
	}
	write(c, header.CmdId, msg.Body)
}

func (m *handler) groupSend(header *QueueHeader, body []byte) {
	//找到需要群发的所有人
}

func (m *handler) allSend(header *QueueHeader, body []byte) {
	//所有人都发
}
