package qa

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleUpdate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
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

    payload, err := decodeQaPayload(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    err = service.UpdateQa(ctx, client, *sel.ID, payload)


    if err == service.ErrDuplicateQa {
        httpx.WriteError(w, http.StatusConflict, err.Error())
        return
    }
    if err != nil {
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
        return
    }

    mapWriteResult(w, err)
}
