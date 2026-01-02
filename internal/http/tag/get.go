package tag

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleGet(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    sel, err := resolveSelector(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    tag, err := service.GetTag(ctx, client, sel)
    respondSingle(w, tag, err)
}
