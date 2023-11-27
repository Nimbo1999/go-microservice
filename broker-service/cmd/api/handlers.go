package main

import (
	"net/http"
)

func (app *Config) Broker(response http.ResponseWriter, request *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hello world!",
	}

	app.writeJSON(response, http.StatusOK, payload)
}
