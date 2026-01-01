package thema

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleDelete(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    id, title, err := resolveSelector(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    err = service.DeleteThema(ctx, client, id, title)
    mapWriteResult(w, err)
}
