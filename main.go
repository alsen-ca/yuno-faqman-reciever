package main

import (
	"log"
	"net/http"

	"yuno-faqman-reciever/internal/db"
	httpHandlers "yuno-faqman-reciever/internal/http"
)

func main() {
	// Connect to Mongo
	client, err := db.ConnectMongo("mongodb://faqman-db:27017")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to database was successfull")

	// Start router
	mux := http.NewServeMux()
	httpHandlers.RegisterThemaRoutes(mux, client)

	// Start server
	log.Println("Listening on :8221")
	err = http.ListenAndServe(":8221", mux)
	if err != nil {
		log.Fatal(err)
	}
}
