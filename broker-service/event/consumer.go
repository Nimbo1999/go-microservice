package event

import (
	"broker/config"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	if err := consumer.setup(); err != nil {
		return Consumer{}, err
	}
	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen(topics []string, env *config.EnvVariables) error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	q, err := declareRandomQueue(channel)
	if err != nil {
		return err
	}
	for _, topic := range topics {
		if err := channel.QueueBind(q.Name, topic, "logs_topic", false, nil); err != nil {
			return err
		}
	}
	messages, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	forever := make(chan error)

	go func(variables *config.EnvVariables) {
		for d := range messages {
			var payload Payload
			if err := json.Unmarshal(d.Body, &payload); err != nil {
				log.Println(err)
				forever <- err
			} else {
				go handlePayload(payload, variables)
			}
		}
	}(env)

	log.Printf("Waiting for message on [Exchange, Queue] [logs_topic, %s]\n", q.Name)
	err = <-forever
	log.Fatalln(err)
	return err
}

func handlePayload(payload Payload, env *config.EnvVariables) {
	switch payload.Name {
	case "log", "event":
		err := logEvent(payload, env)
		if err != nil {
			log.Println(err)
		}
	// case "auth":
	default:
		err := logEvent(payload, env)
		if err != nil {
			log.Println(err)
		}
	}
}

func logEvent(payload Payload, env *config.EnvVariables) error {
	jsonData, _ := json.MarshalIndent(payload, "", "\t")
	request, err := http.NewRequest("POST", env.LogServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[ERROR]: %v\n", err.Error())
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("[ERROR]: %v\n", err.Error())
		return err
	}
	defer response.Body.Close()
	if response.StatusCode > 399 {
		err = errors.New("received an error response code from logger service")
		log.Printf("[ERROR]: %v\n", err.Error())
		return err
	}
	return nil
}
