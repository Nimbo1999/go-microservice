package main

import (
	"broker/config"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	env *config.EnvVariables
}

func main() {
	app := Config{
		env: config.NewEnvVariables(),
	}
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	log.Printf("Starting server on port :%s\n", webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
