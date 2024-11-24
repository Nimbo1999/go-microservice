package main

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrinvalidCredentials error = errors.New("invalid credentials")

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJSON(w, r, &requestPayload); err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, ErrinvalidCredentials, http.StatusBadRequest)
		return
	}

	if valid, err := user.PasswordMatches(requestPayload.Password); err != nil || !valid {
		app.errorJSON(w, ErrinvalidCredentials, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in as user %s", user.Email),
		Data:    user,
	})
}
