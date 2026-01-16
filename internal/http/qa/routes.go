package qa

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(mux *http.ServeMux, client *mongo.Client) {
    mux.HandleFunc("/qa", handler(client))
}
