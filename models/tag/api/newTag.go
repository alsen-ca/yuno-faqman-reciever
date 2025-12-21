func newTag(w http.ResponseWriter, r *http.Request) {
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

	// Write 'tag' to the database
	fmt.Println("\n--- TAG written ---")
}
