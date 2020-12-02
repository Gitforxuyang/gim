package handler

import (
	"fmt"
	"github.com/streadway/amqp"
)

func (m *handler) runConsume() {
	m.queue.InitQueue(m.node)
	m.queue.Consume(m.consume)
}

func (m *handler) consume(msg amqp.Delivery) {
	fmt.Println("handler consume:", msg)
}
