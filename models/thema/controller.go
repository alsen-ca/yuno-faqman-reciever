func controller_themas(http) {
	http.HandleFunc("/thema/new", newThema)
	http.HandleFunc(`/thema/get/${title}`, getThema(title))
	http.HandleFunc(`/thema/update/${uuid}`, updateThema(uuid))
	http.HandleFunc(`/thema/delete/${uuid}`, deleteThema(uuid))
}
