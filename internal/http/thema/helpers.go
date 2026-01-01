package thema

import (
    "encoding/json"
    "net/http"
    "errors"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/domain"
    "yuno-faqman-reciever/internal/service"
    "yuno-faqman-reciever/internal/httpx"
)

type TitlePayload struct {
    Title string `json:"title"`
}

func decodeTitlePayload(r *http.Request) (TitlePayload, error) {
    var payload TitlePayload

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        return payload, errors.New("invalid json")
    }
    if payload.Title == "" {
        return payload, errors.New("title required")
    }
    return payload, nil
}

func mapWriteResult(w http.ResponseWriter, err error) {
    switch err {
    case nil:
        w.WriteHeader(http.StatusNoContent)
    case mongo.ErrNoDocuments:
        httpx.WriteError(w, http.StatusNotFound, "not found")
    case service.ErrDuplicateTitle:
        httpx.WriteError(w, http.StatusConflict, err.Error())
    default:
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
    }
}

func resolveSelector(r *http.Request) (uuid.UUID, string, error) {
    idStr := r.URL.Query().Get("id")
    title := r.URL.Query().Get("title")

    switch {
    case idStr != "" && title != "":
        return uuid.Nil, "", errors.New("only one of id or title allowed")

    case idStr != "":
        id, err := uuid.Parse(idStr)
        if err != nil {
            return uuid.Nil, "", errors.New("invalid uuid")
        }
        return id, "", nil

    case title != "":
        return uuid.Nil, title, nil

    default:
        return uuid.Nil, "", errors.New("missing id or title")
    }
}

func respondSingle(w http.ResponseWriter, thema domain.Thema, err error) {
    if err == mongo.ErrNoDocuments {
        httpx.WriteError(w, http.StatusNotFound, "not found")
        return
    }
    if err != nil {
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
        return
    }
    json.NewEncoder(w).Encode(thema)
}
