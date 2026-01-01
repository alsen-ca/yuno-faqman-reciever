package main

import (
	"log"
	"net/http"
	"context"

	"yuno-faqman-reciever/internal/db"
	"yuno-faqman-reciever/internal/http/thema"
)

func main() {
	// Connect to Mongo
	client, err := db.ConnectMongo("mongodb://faqman-db:27017")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to database was successfull")

	// Add Indexes to Mongo
	ctx := context.Background()
	if err := db.EnsureThemaIndexes(ctx, client); err != nil {
		log.Fatal(err)
	}
	log.Println("Indexes to database added successfully")


	// Start router
	mux := http.NewServeMux()
	thema.RegisterRoutes(mux, client)

	// Start server
	log.Println("Listening on :8221")
	err = http.ListenAndServe(":8221", mux)
	if err != nil {
		log.Fatal(err)
	}
}
