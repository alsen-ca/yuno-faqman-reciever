package thema

import (
    "net/http"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleGet(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()
    idStr := r.URL.Query().Get("id")
    title := r.URL.Query().Get("title")

    switch {
    case idStr != "":
        id, err := uuid.Parse(idStr)
        if err != nil {
            httpx.WriteError(w, http.StatusBadRequest, "invalid uuid")
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
            httpx.WriteError(w, http.StatusInternalServerError, "internal error")
            return
        }
        httpx.WriteJSON(w, http.StatusOK, themas)
    }
}
