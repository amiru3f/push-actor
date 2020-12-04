package consume

import (
	"sync"

	"github.com/streadway/amqp"
)

var once sync.Once
var instance connection

type connection struct {
	connectionString string
	name             string
	conn             *amqp.Connection
	readChannel      *amqp.Channel
	writeChannel     *amqp.Channel
	err              chan error
}

func getConnectionInstance(name string, connectionString string) *connection {

	once.Do(func() {
		instance = connection{}
		instance.connectionString = connectionString
		instance.name = name
	})

	return &instance
}

func (c *connection) Connect(recoveryFunc func()) error {

	if c.conn != nil && !c.conn.IsClosed() {

		if c.readChannel != nil {
			c.readChannel.Close()
		}

		c.conn.Close()
	}

	var err error
	c.conn, err = amqp.Dial(c.connectionString)

	if err != nil {
		return err
	}

	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error))
		recoveryFunc()
	}()

	return nil
}

func (c *connection) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}

	panic("no connection instance")
}

func (c *connection) ReadChannel() (*amqp.Channel, error) {
	var err error

	c.readChannel, err = c.conn.Channel()

	if err != nil {
		return nil, err
	}

	return c.readChannel, nil
}
