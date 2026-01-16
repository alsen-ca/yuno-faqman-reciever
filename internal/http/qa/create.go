package qa

import (
    "net/http"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"

)

func handleCreate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    payload, err := decodeQaPayload(r)
    
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }
    log.Printf("Payload: %s", payload)
    

    qa, err := service.CreateQa(ctx, client, payload)
    if err == service.ErrDuplicateTag {
        httpx.WriteError(w, http.StatusConflict, err.Error())
        return
    }
    if err != nil {
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
        return
    }

    httpx.WriteJSON(w, http.StatusCreated, qa)
}
