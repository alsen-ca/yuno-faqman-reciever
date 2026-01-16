package main

import (
	"log"
	"net/http"
	"context"

	"yuno-faqman-reciever/internal/db"
	"yuno-faqman-reciever/internal/http/thema"
	"yuno-faqman-reciever/internal/http/tag"
	"yuno-faqman-reciever/internal/http/qa"
	"yuno-faqman-reciever/internal/middleware"
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
	if err := db.EnsureTagIndexes(ctx, client); err != nil {
		log.Fatal(err)
	}
	if err := db.EnsureQaIndexes(ctx, client); err != nil {
		log.Fatal(err)
	}
	log.Println("Indexes to database added successfully")


	// Start routers
	mux := http.NewServeMux()
	thema.RegisterRoutes(mux, client)
	tag.RegisterRoutes(mux, client)
	qa.RegisterRoutes(mux, client)

	// Start server
	handler := middleware.Logging(mux)
	log.Println("Listening on :8221")
	log.Fatal(http.ListenAndServe(":8221", handler))
}
