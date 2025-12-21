import ("thema")

func newThema(w http.ResponseWriter, r *http.Request) {
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

	// Write 'thema' to the database
	fmt.Println("\n--- THEMA written ---")
}