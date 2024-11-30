package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	webPort        string
	rpcPort        string
	mongoUrl       string
	grpcPort       string
	allowedOrigins string
	client         *mongo.Client
)

type Config struct {
	Models data.Models
}

func init() {
	webPort = os.Getenv("WEB_PORT")
	rpcPort = os.Getenv("RPC_PORT")
	mongoUrl = os.Getenv("MONGO_URL")
	grpcPort = os.Getenv("GRPC_PORT")
	allowedOrigins = os.Getenv("ALLOWED_ORIGIN")
}

func main() {
	mongoConnectionTimeout, cancelConnectionContext := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelConnectionContext()
	mongoClient, err := connectToMongo(mongoConnectionTimeout)
	if err != nil {
		log.Panic(err)
	}

	ctx, disconnectCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer disconnectCancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	app := Config{
		Models: data.New(mongoClient),
	}
	if err := app.Server(); err != nil {
		log.Panicln(err)
	}
}

func (app *Config) Server() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(allowedOrigins),
	}
	return srv.ListenAndServe()
}

func connectToMongo(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	return mongo.Connect(ctx, clientOptions)
}
