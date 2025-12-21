// Updates a Tag based on its uuid
func updateTag(w http.ResponseWriter, r *http.Request, uuid string) {
	var tag Tag
	// 404 if no Tag with uuid can be found; returns
	status_code = 404; return

	// Tries to update with the 'body' of the request
	// 400 if invalid body
	status_code = 400; return

	// 20x (202 maybe?) if it could be updated
	status_code = 202; return status_code, tag
}
