// Make this application work as a microservice and listen to data for API
func microservice() {
	address := "127.0.0.1:8221"
	fmt.Println("Listening on http://" + address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
