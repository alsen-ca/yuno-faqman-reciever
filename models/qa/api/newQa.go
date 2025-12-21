func newQA(w http.ResponseWriter, r *http.Request) {
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
