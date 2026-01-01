package thema

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"
)

func handler(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            handleGet(w, r, client)
        case http.MethodPost:
            handleCreate(w, r, client)
        case http.MethodPut:
            handleUpdate(w, r, client)
        case http.MethodDelete:
            handleDelete(w, r, client)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    }
}
