package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Thema struct {
	Title string `json:"title"`
}

func handleNewThema(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var thema Thema

	err := json.NewDecoder(r.Body).Decode(&thema)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	fmt.Println("\n--- THEMA RECEIVED ---")
	fmt.Printf("Title: %s\n", thema.Title)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func main() {
	http.HandleFunc("/thema", handleNewThema)

	address := "127.0.0.1:3200"
	fmt.Println("Listening on http://" + address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
