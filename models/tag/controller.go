func controller_tags(http) {
	http.HandleFunc("/tag/new", newTag)
	http.HandleFunc(`/tag/get/${lang_tag}`, getTag(lang_tag))
	http.HandleFunc("tag/all"), allTags(which)
	http.HandleFunc(`/tag/update/${uuid}`, updateTag(uuid))
	http.HandleFunc(`/tag/delete/${uuid}`, deleteTag(uuid))
}
