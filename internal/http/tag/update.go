package tag

import (
    "net/http"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleUpdate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    // TODO
    id, oldTitle, err := resolveSelector(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    payload, err := decodePayload(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    if id != uuid.Nil {
        err = service.UpdateTagTitle(ctx, client, id, payload.Title)
    }

    mapWriteResult(w, err)
}
