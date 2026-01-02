package tag

import (
    "net/http"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

func handleGet(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()
    idStr := r.URL.Query().Get("id")
    en_og := r.URL.Query().Get("en_og")
    de_trans := r.URL.Query().Get("de_trans")
    es_trans := r.URL.Query().Get("es_trans")

    switch {
    case idStr != "":
        id, err := uuid.Parse(idStr)
        if err != nil {
            httpx.WriteError(w, http.StatusBadRequest, "invalid uuid")
            return
        }

        tag, err := service.GetTagByID(ctx, client, id)
        // TODO respondSingle
        respondSingle(w, tag, err)

    case en_og != "":
        tag, err := service.GetTagByEnOriginal(ctx, client, en_og)
        respondSingle(w, tag, err)
    case de_trans != "":
        tag, err := service.GetTagByDeTranslation(ctx, client, de_trans)
        respondSingle(w, tag, err)
    case es_trans != "":
        tag, err := service.GetTagByEsTranslation(ctx, client, es_trans)
        respondSingle(w, tag, err)

    default:
        tags, err := service.ListTags(ctx, client)
        if err != nil {
            httpx.WriteError(w, http.StatusInternalServerError, "internal error")
            return
        }
        httpx.WriteJSON(w, http.StatusOK, tags)
    }
}
