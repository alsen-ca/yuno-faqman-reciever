package http

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"
)

func RegisterThemaRoutes(mux *http.ServeMux, client *mongo.Client) {
    mux.HandleFunc("/thema/new", newThemaHandler(client))
}

func newThemaHandler(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if r.Method != http.MethodPost {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK\n"))
    }
}
