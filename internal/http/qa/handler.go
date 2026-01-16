package qa

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/httpx"
)

func handler(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            handleCreate(w, r, client)
        default:
            httpx.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
        }
    }
}
