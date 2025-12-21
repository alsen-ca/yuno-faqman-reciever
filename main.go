package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
// Maybe import:
// microservice.go
// db.go
// models/*/controller.go

func main() {
	controller_themas(http)
	controller_tags(http)
	controller_qas(http)

	microservice()
	connection()
	// run_algorithm()
}
