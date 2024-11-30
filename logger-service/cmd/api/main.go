package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	webPort  string
	rpcPort  string
	mongoUrl string
	grpcPort string
	client   *mongo.Client
)

type Config struct {
}

func init() {
	webPort = os.Getenv("APP_WEB_PORT")
	rpcPort = os.Getenv("APP_RPC_PORT")
	mongoUrl = os.Getenv("APP_MONGO_URL")
	grpcPort = os.Getenv("APP_GRPC_PORT")
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
}

func connectToMongo(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	return mongo.Connect(ctx, clientOptions)
}
