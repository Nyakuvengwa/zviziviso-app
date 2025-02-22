package main

import (
	"context"
	"fmt"
	"log"
	"os"
	repository "zviziviso-app/internal/db"

	"github.com/jackc/pgx/v5"
)

func main() {

	appConfig := applicationConfig{
		address: os.Getenv("ZVIZIVISO_ADDRESS"),
	}

	fmt.Println(appConfig)
	ctx := context.Background()

	connectionString := os.Getenv("ZVIZIVISO_DB_CONNECTION_STRING")

	if connectionString == "" {
		log.Fatal("Connection string is not set")
	}
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close(ctx)

	dbQueries := repository.New(conn)
	app := Application{
		config:  appConfig,
		service: dbQueries,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
