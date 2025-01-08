package main

import (
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close()

	consumer, err := event.NewConsumer(connection)
	if err != nil {
		log.Fatalln(err)
	}
	if err := consumer.Listen([]string{"log.DEBUG", "log.INFO", "log.ERROR"}); err != nil {
		log.Fatalln(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial(os.Getenv("AMQP_URL"))
		if err != nil {
			log.Println("RabbitMQ is not ready yet...")
			log.Println(err)
			counts++
		} else {
			connection = c
			log.Println("Connected to amqp!")
			break
		}
		if counts > 5 {
			log.Println(err)
			return nil, err
		}
		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backoff)
	}
	return connection, nil
}
