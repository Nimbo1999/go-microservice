package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connection *amqp.Connection
}

func (emitter *Emitter) setup() error {
	channel, err := emitter.connection.Channel()
	if err != nil {
		log.Println(err)
		return err
	}
	defer channel.Close()
	return declareExchange(channel)
}

func (emitter *Emitter) Push(event, severity string) error {
	channel, err := emitter.connection.Channel()
	if err != nil {
		log.Println(err)
		return err
	}
	defer channel.Close()
	log.Println("Pushing to channel")
	return channel.Publish("logs_topic", severity, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(event),
	})
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}
	if err := emitter.setup(); err != nil {
		log.Println(err)
		return Emitter{}, err
	}
	return emitter, nil
}
