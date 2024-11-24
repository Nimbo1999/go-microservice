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

var webPort string

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func init() {
	webPort = os.Getenv("PORT")
}

func main() {
	log.Println("Starting Authentication service")

	// TODO: connect to DB
	log.Println("Connection to the db")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db, err := connectToDB(ctx)
	if err != nil {
		log.Println("Could not connect to the DB")
		log.Panic(err)
	}
	defer db.Close()
	log.Println("Connected to the Database successfuly")

	// Set up Config
	app := Config{
		DB:     db,
		Models: data.New(db),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string, ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}
	return db, err
}

func connectToDB(ctx context.Context) (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	return openDB(dsn, ctx)
}
