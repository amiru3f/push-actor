package consume

import (
	"errors"
	"fmt"

	"github.com/albb-b2b/push2b/internal/dispatch"
	"github.com/albb-b2b/push2b/pkg"
)

type rabbitConsumer struct {
	host       string
	port       int
	username   string
	password   string
	dispatcher dispatch.Dispatcher
}

func NewRabbitConsumer(config pkg.RabbitConfig, dispatcher dispatch.Dispatcher) Consumer {

	if dispatcher == nil {
		panic("dispatcher can not be nil")
	}

	return &rabbitConsumer{config.Host, config.Port, config.User, config.Password, dispatcher}
}

func (c rabbitConsumer) Consume() error {

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d", c.username, c.password, c.host, c.port)
	conn := getConnectionInstance("test", connectionString)
	err := conn.Connect(closedEvent)

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

	//topology and connection recovery
	defer c.Consume()
	return <-forever
}
