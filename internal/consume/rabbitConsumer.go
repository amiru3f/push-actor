package consume

import (
	"errors"
	"fmt"
	"time"

	"github.com/albb-b2b/push2b/internal/dispatch"
	"github.com/albb-b2b/push2b/pkg"
)

type rabbitConsumer struct {
	config     pkg.RabbitConfig
	dispatcher dispatch.Dispatcher
}

func NewRabbitConsumer(config pkg.RabbitConfig, dispatcher dispatch.Dispatcher) Consumer {

	if dispatcher == nil {
		panic("dispatcher can not be nil")
	}

	return &rabbitConsumer{config, dispatcher}
}

func (c rabbitConsumer) Consume() error {

	//topology and connection recovery
	//this will be called after connection break or connection refuse.
	defer func() {
		count := time.Duration(c.config.RetrySeconds)
		time.Sleep(time.Second * count)
		c.Consume()
	}()

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d", c.config.User, c.config.Password, c.config.Host, c.config.Port)
	conn := getConnectionInstance("test", connectionString)
	err := conn.Connect(consumingInterrupted)

	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.ReadChannel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"altrabo.event.mail", // name
		"topic",              // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)

	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,                 // queue name
		"altrabo.event.mail.*", // routing key
		"altrabo.event.mail",   // exchange
		false,
		nil)

	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan error)

	go func() {
		for d := range msgs {
			work := pkg.Work{Header: d.RoutingKey, Payload: string(d.Body)}
			go c.dispatcher.Received(work)
		}

		forever <- errors.New("connection/channel closed")
	}()

	return <-forever
}
