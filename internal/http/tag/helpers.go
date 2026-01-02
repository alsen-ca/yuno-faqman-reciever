package tag

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

type TagPayload struct {
    EnOg string `json:"en_og"`
    DeTrans string `json:"de_trans"`
    EsTrans string `json:"es_trans"`
}

type TagSelector struct {
    ID       *uuid.UUID
    EnOg     *string
    DeTrans *string
    EsTrans *string
}


func decodeTagPayload(r *http.Request) (TitlePayload, error) {
    var payload TagPayload

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        return payload, errors.New("invalid json")
    }
    if payload.EnOg == "" {
        return payload, errors.New("en_og required")
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

func resolveSelector(r *http.Request) (TagSelector, error) {
    var payload TagPayload
    count := 0

    if idStr := r.URL.Query().Get("id"); idStr != "" {
        id, err := uuid.Parse(idStr)
        if err != nil {
            return sel, errors.New("invalid uuid")
        }
        sel.ID = &id
        count++
    }

    if v := r.URL.Query().Get("en_og"); v != "" {
        sel.EnOg = &v
        count++
    }

    if v := r.URL.Query().Get("de_trans"); v != "" {
        sel.DeTrans = &v
        count++
    }

    if v := r.URL.Query().Get("es_trans"); v != "" {
        sel.EsTrans = &v
        count++
    }

    if count == 0 {
        return sel, errors.New("missing selector")
    }
    if count > 1 {
        return sel, errors.New("only one selector allowed")
    }

    return sel, nil
}

func respondSingle(w http.ResponseWriter, tag domain.Tag, err error) {
    if err == mongo.ErrNoDocuments {
        httpx.WriteError(w, http.StatusNotFound, "not found")
        return
    }
    if err != nil {
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
        return
    }
    json.NewEncoder(w).Encode(tag)
}
