// Recieves the question for a QA and returns the whole QA (including uuid)
func getQa(w http.ResponseWriter, r *http.Request, question string) {
	// Takes the question and does a 1:1 search for it on the QA table
	// SELECT * from qa where question like `question`
	// qa = Qa.where(question: question)

	var qa Qa
	
	// 404 if could not find the answer
	// 200 if answers.

	// Returns 200 and the answers
}
