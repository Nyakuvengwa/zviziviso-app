package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	repository "zviziviso-app/api/internal/db"
	"zviziviso-app/api/internal/handlers"
	"zviziviso-app/api/internal/http/middleware"
	"zviziviso-app/api/internal/routing"
	"zviziviso-app/api/internal/services"

	"github.com/jackc/pgx/v5"
)

func main() {
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

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Hello World")
	})

	router.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	handler := handlers.NewCountryHandler(services.NewCountryService(*dbQueries))
	countryRoutes := routing.NewCountryRoutes(handler)
	countryRoutes.SetupCountriesRoutes(router)

	stack := middleware.CreateStack(
		middleware.Logging,
	)
	server := http.Server{
		Addr:    ":7653",
		Handler: stack(router),
	}

	log.Fatal(server.ListenAndServe())
}
