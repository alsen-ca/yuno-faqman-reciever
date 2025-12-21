func controller_qas(http) {
	http.HandleFunc("/qa/new", newQA)
	http.HandleFunc(`/qa/get/${question}`, getQas(question))
	http.HandleFunc(`/qa/update/${uuid}`, updateQas(uuid))
	http.HandleFunc(`/qa/delete/${uuid}`, deleteQas(uuid))
}
