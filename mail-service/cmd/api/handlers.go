package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}
	var payload mailMessage
	if err := app.readJSON(w, r, &payload); err != nil {
		log.Println("[ERROR]:", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	msg := Message{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Data:    payload.Message,
	}
	log.Println("Message:")
	log.Println(msg)
	if err := app.Mailer.SendSMTPMessage(msg); err != nil {
		log.Println("[ERROR]:", err)
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	if err := app.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("email message sent to %s", msg.To),
	}); err != nil {
		log.Println("[ERROR]:", err)
		app.errorJSON(w, err, http.StatusInternalServerError)
	}
}
