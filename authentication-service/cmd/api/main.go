package main

import (
	"authentication/data"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	fmt.Println("Starting the authentication service!")

	// TODO: connect to DB
	conn, err := openDB(os.Getenv("DB_DNS"))
	if err != nil {
		log.Panicln(err)
	}

	// set up config
	config := Config{
		DB:     conn,
		Models: data.New(conn),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: config.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func openDB(dns string) (*sql.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, err
}
