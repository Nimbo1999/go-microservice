package main

import (
	"fmt"
	"log"
	"mail-service/config"
	"net/http"
)

type Config struct {
	env    *config.EnvVariables
	Mailer *Mail
}

func main() {
	env := config.NewEnvVariables()
	app := &Config{
		env:    env,
		Mailer: NewMail(env),
	}

	log.Printf("Web server initializing on port %s\n", app.env.WebPort)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", app.env.WebPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
