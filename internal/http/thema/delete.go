package thema

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
)

func handleDelete(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    id, title, err := resolveSelector(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = service.DeleteThema(ctx, client, id, title)
    mapWriteResult(w, err)
}
