package qa

import (
    "encoding/json"
    "net/http"
    "fmt"
    "yuno-faqman-reciever/internal/domain"
)

func decodeQaPayload(r *http.Request) (domain.QaPayload, error) {
	var payload domain.QaPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return payload, fmt.Errorf("invalid json: %v", err)
	}

	if payload.Question == "" {
		return payload, fmt.Errorf("question required")
	}
	if len(payload.QuestionWeights) == 0 {
		return payload, fmt.Errorf("question weights required")
	}
	if payload.Answer == "" {
		return payload, fmt.Errorf("answer required")
	}
	if payload.Language == "" && (payload.Language != "en" || payload.Language == "de" || payload.Language == "es") {
		return payload, fmt.Errorf("language must be either 'en', 'de' or 'es'")
	}

	return payload, nil
}
/*
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

func resolveSelector(r *http.Request) (*domain.TagSelector, error) {
    var sel domain.TagSelector
    count := 0

    if idStr := r.URL.Query().Get("id"); idStr != "" {
        id, err := uuid.Parse(idStr)
        if err != nil {
            return nil, errors.New("invalid uuid")
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
        return nil, nil
    }
    if count > 1 {
        return nil, errors.New("only one selector allowed")
    }

    return &sel, nil
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
*/