package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func api(client *sql.DB) {
	mux := http.NewServeMux()

	routes(mux, client)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Listening on port %s\n", port)
	server.ListenAndServe()
}

func main() {
	// loader .env fil og variablene i filen
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	// henter neon url (connection string) fra .env fil
	NEON_URL := os.Getenv("NEON_URL")

	// lager en instans av databasen med neon url-en
	db, err := sql.Open("postgres", NEON_URL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
	defer db.Close()

	api(db)
}
