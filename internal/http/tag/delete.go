package tag

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleDelete(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    sel, err := resolveSelector(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    if sel.ID == nil {
        httpx.WriteError(w, http.StatusBadRequest, "id is required")
        return
    }

    err = service.DeleteTag(ctx, client, *sel.ID)
    mapWriteResult(w, err)
}
