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

type Tag struct {
	EnOg string `json:"en_og"`;
	DeTranslation string `json:"de_trans"`;
	EsTranslation string `json:"es_trans"`
}

type Qa struct {
	Question string `json:"question"`;
	Answer string `json:"answer"`;
	Language string `json:"lang"`
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

func handleNewTag(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var tag Tag

	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	fmt.Println("\n--- TAG RECEIVED ---")
	fmt.Printf("Tag English Original: %s\n", tag.EnOg)
	fmt.Printf("Tag German Translation: %s\n", tag.DeTranslation)
	fmt.Printf("Tag Spanish Translation: %s\n", tag.EsTranslation)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func handleNewQA(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var qa Qa

	err := json.NewDecoder(r.Body).Decode(&qa)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	fmt.Println("\n--- QA RECEIVED ---")
	fmt.Printf("Question: %s\n", qa.Question)
	fmt.Printf("Answer: %s\n", qa.Answer)
	fmt.Printf("QA Language: %s\n", qa.Language)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func main() {
	http.HandleFunc("/thema", handleNewThema)
	http.HandleFunc("/tag", handleNewTag)
	http.HandleFunc("/qa", handleNewQA)

	address := "127.0.0.1:3200"
	fmt.Println("Listening on http://" + address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
