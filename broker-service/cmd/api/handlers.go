package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
