package thema

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleCreate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    payload, err := decodeTitlePayload(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    thema, err := service.CreateThema(ctx, client, payload.Title)
    if err == service.ErrDuplicateTitle {
        httpx.WriteError(w, http.StatusConflict, err.Error())
        return
    }
    if err != nil {
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
        return
    }

    httpx.WriteJSON(w, http.StatusCreated, thema)
}
