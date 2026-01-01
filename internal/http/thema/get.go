package thema

import (
    "encoding/json"
    "net/http"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
)

func handleGet(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()
    idStr := r.URL.Query().Get("id")
    title := r.URL.Query().Get("title")

    switch {
    case idStr != "":
        id, err := uuid.Parse(idStr)
        if err != nil {
            http.Error(w, "invalid uuid", http.StatusBadRequest)
            return
        }

        thema, err := service.GetThemaByID(ctx, client, id)
        respondSingle(w, thema, err)

    case title != "":
        thema, err := service.GetThemaByTitle(ctx, client, title)
        respondSingle(w, thema, err)

    default:
        themas, err := service.ListThemas(ctx, client)
        if err != nil {
            http.Error(w, "internal error", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(themas)
    }
}
