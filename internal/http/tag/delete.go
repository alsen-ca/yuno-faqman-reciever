package tag

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleDelete(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    // TODO
    id, err := resolveSelector(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    // TODO
    err = service.DeleteTag(ctx, client, id)
    mapWriteResult(w, err)
}
