func controller_qas(http) {
	http.HandleFunc("/qa/new", newQA)
	http.HandleFunc(`/qa/get/${question}`, getTag(question))
	http.HandleFunc(`/qa/update/${uuid}`, updateTag(uuid))
	http.HandleFunc(`/qa/delete/${uuid}`, deleteTag(uuid))
}
