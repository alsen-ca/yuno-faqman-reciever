package qa

import (
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleGet(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    selPoE, err := resolveSelector(r)
    if err != nil {
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    if selPoE == nil {
        qas, err := service.ListQas(ctx, client)
        if err != nil {
            httpx.WriteError(w, http.StatusInternalServerError, "internal error")
            return
        }
        httpx.WriteJSON(w, http.StatusOK, qas)
        return
    }


    sel := *selPoE
    qa, err := service.GetQa(ctx, client, sel)
    respondSingle(w, qa, err)
}
