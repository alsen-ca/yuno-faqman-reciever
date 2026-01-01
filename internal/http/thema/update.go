package thema

import (
    "net/http"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleUpdate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    id, oldTitle, err := resolveSelector(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    payload, err := decodeTitlePayload(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    if id != uuid.Nil {
        err = service.UpdateThemaTitle(ctx, client, id, payload.Title)
    } else {
        err = service.UpdateThemaTitleByTitle(ctx, client, oldTitle, payload.Title)
    }

    mapWriteResult(w, err)
}
