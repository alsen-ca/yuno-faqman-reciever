package tag

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

// TODO
func handleCreate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    payload, err := decodeTagPayload(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    tag, err := service.CreateTag(ctx, client, payload.ToDomain())
    if err == service.ErrDuplicateTag {
        httpx.WriteError(w, http.StatusConflict, err.Error())
        return
    }
    if err != nil {
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
        return
    }

    httpx.WriteJSON(w, http.StatusCreated, tag)
}
