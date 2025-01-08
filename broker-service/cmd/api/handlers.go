package main

import (
	"broker/event"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(response http.ResponseWriter, request *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hello world!",
	}

	app.writeJSON(response, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(response http.ResponseWriter, request *http.Request) {
	var payload RequestPayload

	if err := app.readJSON(response, request, &payload); err != nil {
		app.errorJSON(response, err, http.StatusBadRequest)
		return
	}

	switch payload.Action {
	case "auth":
		app.authenticate(response, payload.Auth)
	case "mail":
		app.sendMail(response, payload.Mail)
	case "log":
		app.logItemViaRabbitMQ(response, payload.Log)
	default:
		app.errorJSON(response, errors.New("unknown action"), http.StatusBadRequest)
	}
}

func (app *Config) authenticate(w http.ResponseWriter, p AuthPayload) {
	jsonData, _ := json.MarshalIndent(p, "", "\t")
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/authenticate", app.env.AuthServiceUrl), bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"), response.StatusCode)
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"), http.StatusInternalServerError)
		return
	}

	var dataFromService jsonResponse
	if err = json.NewDecoder(response.Body).Decode(&dataFromService); err != nil {
		app.errorJSON(w, errors.New("error while parsing the json data"), http.StatusInternalServerError)
		return
	}

	if dataFromService.Error {
		app.errorJSON(w, errors.New(dataFromService.Message), http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = dataFromService.Data
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) sendMail(w http.ResponseWriter, p MailPayload) {
	jsonData, _ := json.MarshalIndent(p, "", "\t")
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/send", app.env.MailServiceUrl), bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	if response.StatusCode > 300 {
		app.errorJSON(w, errors.New("unabled to communicate with the mail service"), response.StatusCode)
		return
	}
	var payload jsonResponse
	payload.Error = false
	payload.Message = fmt.Sprintf("Message sent to %s", p.To)
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) logItem(w http.ResponseWriter, entry LogPayload) {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	request, err := http.NewRequest("POST", app.env.LogServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[ERROR]: %v\n", err.Error())
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("[ERROR]: %v\n", err.Error())
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	if response.StatusCode > 399 {
		err = errors.New("received an error response code from logger service")
		log.Printf("[ERROR]: %v\n", err.Error())
		app.errorJSON(w, err, response.StatusCode)
		return
	}
	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged!"
	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) logItemViaRabbitMQ(w http.ResponseWriter, entry LogPayload) {
	if err := app.pushToQueue(entry.Name, entry.Data); err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged with RabbitMQ"
	if err := app.writeJSON(w, http.StatusOK, payload); err != nil {
		log.Println(err)
		return
	}
}

func (app *Config) pushToQueue(name, message string) error {
	emitter, err := event.NewEventEmitter(app.RabbitMQ)
	if err != nil {
		return err
	}

	payload := LogPayload{
		Name: name,
		Data: message,
	}

	j, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println(err)
		return err
	}

	return emitter.Push(string(j), "log.INFO")
}
