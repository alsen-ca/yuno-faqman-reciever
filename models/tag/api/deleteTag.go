// Deletes a Tag based on its uuid
func deleteTag(w http.ResponseWriter, r *http.Request, uuid string) {
	var tag Tag
	
	// Tries to delete a Tag with that uuid
	// 404 if not exists
	// 20x (204 I think?) if it could be deleted
}
