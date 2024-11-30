package main

import (
	"log-service/data"
	"net/http"
)

type JsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))

	var payload JsonPayload
	if err := app.readJSON(w, r, &payload); err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	entry := data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	}

	if err := app.Models.LogEntry.Insert(entry, r.Context()); err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "logged!",
	})
}
