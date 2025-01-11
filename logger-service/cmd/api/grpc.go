package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogsService struct {
	logs.UnimplementedLogServiceServer
	Model data.Models
}

func (*LogsService) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	entry := data.LogEntry{
		Name: input.GetName(),
		Data: input.GetData(),
	}
	if err := entry.Insert(entry, ctx); err != nil {
		log.Println("[ERROR]:", err)
		return &logs.LogResponse{Result: "failed to log!"}, err
	}
	return &logs.LogResponse{Result: "logged with gRPC!"}, nil
}

func (app *Config) gRPCListen() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Println("[ERROR]:", err)
		return err
	}

	server := grpc.NewServer()
	logs.RegisterLogServiceServer(server, &LogsService{Model: app.Models})
	log.Println("gRPC server started on port", grpcPort)
	return server.Serve(listener)
}
