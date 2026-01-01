package thema

import (
    "encoding/json"
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
)

func handleCreate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    payload, err := decodeTitlePayload(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    thema, err := service.CreateThema(ctx, client, payload.Title)
    if err == service.ErrDuplicateTitle {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(thema)
}
