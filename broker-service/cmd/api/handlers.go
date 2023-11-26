package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Broker(response http.ResponseWriter, request *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hello world!",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(out)
}
