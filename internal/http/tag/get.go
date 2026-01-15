package tag

import (
    "net/http"
    "log"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleGet(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    selPoE, err := resolveSelector(r)
    if err != nil {
        log.Printf("No selector was given? %s", err)
        httpx.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    if selPoE == nil {
        tags, err := service.ListTags(ctx, client)
        if err != nil {
            httpx.WriteError(w, http.StatusInternalServerError, "internal error")
            return
        }
        httpx.WriteJSON(w, http.StatusOK, tags)
        return
    }


    // Dereferences pointer sent from resolveSelector (plain domain.Selector instead of *domain.Selector)
    sel := *selPoE
    tag, err := service.GetTag(ctx, client, sel)
    respondSingle(w, tag, err)
}
