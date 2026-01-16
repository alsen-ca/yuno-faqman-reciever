package qa

import (
    "encoding/json"
    "net/http"
    "fmt"
    "errors"
    "log"
    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"
    "yuno-faqman-reciever/internal/domain"
    "yuno-faqman-reciever/internal/httpx"
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

func mapWriteResult(w http.ResponseWriter, err error) {
    switch err {
    case nil:
        w.WriteHeader(http.StatusNoContent)
    case mongo.ErrNoDocuments:
        httpx.WriteError(w, http.StatusNotFound, "not found")
    default:
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
    }
}

func resolveSelector(r *http.Request) (*domain.QaSelector, error) {
    var sel domain.QaSelector

    if idStr := r.URL.Query().Get("id"); idStr != "" {
        id, err := uuid.Parse(idStr)
        if err != nil {
            return nil, errors.New("invalid uuid")
        }
        sel.ID = &id
        return &sel, nil
    }
    if questionStr := r.URL.Query().Get("question"); questionStr != "" {
        sel.Question = &questionStr
        return &sel, nil
    }
    
    return nil, nil
}

func respondSingle(w http.ResponseWriter, qa domain.Qa, err error) {
    if err == mongo.ErrNoDocuments {
        httpx.WriteError(w, http.StatusNotFound, "not found")
        return
    }
    if err != nil {
        httpx.WriteError(w, http.StatusInternalServerError, "internal error")
        return
    }
    json.NewEncoder(w).Encode(qa)
}
