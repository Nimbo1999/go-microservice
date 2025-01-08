package main

import (
	"broker/config"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "80"

type Config struct {
	env      *config.EnvVariables
	RabbitMQ *amqp.Connection
}

func main() {
	env := config.NewEnvVariables()
	rabbitMq, err := connect(env)
	if err != nil {
		log.Fatalln(err)
	}
	app := Config{
		env:      env,
		RabbitMQ: rabbitMq,
	}
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	log.Printf("Starting server on port :%s\n", webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func connect(env *config.EnvVariables) (*amqp.Connection, error) {
	var counts int
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial(env.AMQP_URL)
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
