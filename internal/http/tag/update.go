package tag

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleUpdate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    sel, err := resolveSelector(r)
    if sel.ID == nil {
        httpx.WriteError(w, http.StatusBadRequest, "id is required")
        return
    }
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    payload, err := decodeTagPayload(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    err = service.UpdateTag(ctx, client, *sel.ID, payload)


    mapWriteResult(w, err)
}
