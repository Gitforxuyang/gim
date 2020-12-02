package rabbit

import (
	"fmt"
	"gim/conf"
	"gim/infra/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Queue struct {
	c     *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
	node  string
}

type ConsumeHandle func(msg amqp.Delivery)

func (m *Queue) Consume(h ConsumeHandle) {
	for i := 0; i < 2; i++ {
		go func(index int) {
			defer func() {
				if e := recover(); e != nil {
					logrus.Errorln("consume panic err:", e)
				}
			}()
			ch, _ := m.c.Channel()
			ch.Qos(500, 0, false)
			msgs, err := ch.Consume(m.queue.Name, fmt.Sprintf("gim_node.%s-%d", m.node, index), true, false, false, false, nil)
			for {
				utils.Must(err)
				msg := <-msgs
				h(msg)
			}
		}(i)
	}
}
func (m *Queue) InitQueue(node string) {
	queue := fmt.Sprintf("gim_node_%s", node)
	q, err := m.ch.QueueDeclare(queue, true, true, false, false, nil)
	utils.Must(err)
	err = m.ch.QueueBind(queue, fmt.Sprintf("gim_node.%s", node), "im.topic", false, nil)
	utils.Must(err)
	m.queue = q
	m.node = node
}
func InitClient(config *conf.Config) *Queue {
	c, err := amqp.Dial(config.Rabbit.Addr)
	utils.Must(err)
	ch, err := c.Channel()
	utils.Must(err)
	mq := Queue{}
	mq.c = c
	mq.ch = ch
	return &mq
}
