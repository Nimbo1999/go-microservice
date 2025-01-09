package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (server *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	entry := data.LogEntry{}
	if err := entry.Insert(data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	}, context.TODO()); err != nil {
		log.Panicln("[ERROR]:", err)
		return err
	}
	*resp = fmt.Sprintf("Processed payload via RPC: %s", payload.Name)
	return nil
}
