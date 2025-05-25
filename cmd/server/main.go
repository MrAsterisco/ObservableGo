package main

import (
	"log"
	"net/http"
	"os"

	"io.github.mrasterisco/observablego/internal/db"
	"io.github.mrasterisco/observablego/internal/user"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}
	dbConn, err := db.Connect(dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer dbConn.Close()

	if err := db.Migrate(dbConn); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	repo := user.NewRepository(dbConn)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	http.HandleFunc("/users", handler.Users)
	http.HandleFunc("/users/", handler.UserByID)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
