package thema

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/httpx"
)

func RegisterRoutes(mux *http.ServeMux, client *mongo.Client) {
    mux.HandleFunc("/thema", handler(client))
}

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
            httpx.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
        }
    }
}
