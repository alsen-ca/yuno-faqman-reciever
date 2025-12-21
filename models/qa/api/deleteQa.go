// Deletes a Qa based on its uuid
func deleteQa(w http.ResponseWriter, r *http.Request, uuid string) {
	var qa Qa
	
	// Tries to delete a qa with that uuid
	// 404 if not exists
	// 20x (204 I think?) if it could be deleted
}
