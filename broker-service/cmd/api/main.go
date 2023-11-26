package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	log.Printf("Starting server on port :%s\n", webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
