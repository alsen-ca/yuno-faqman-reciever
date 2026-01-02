package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"strings"
	"encoding/json"
    "net/url"
    "github.com/google/uuid"
	"yuno-faqman-reciever/internal/domain"
)

type ThemaSelector struct {
    ID    *uuid.UUID
    Title *string
}

func ByID(id uuid.UUID) ThemaSelector {
    return ThemaSelector{
        ID: &id,
    }
}

func ByTitle(title string) ThemaSelector {
    return ThemaSelector{
        Title: &title,
    }
}

func CreateThemaHTTP(t *testing.T, handler http.Handler, title string) (int, domain.Thema) {
    t.Helper()

    payload := fmt.Sprintf(`{"title":"%s"}`, title)
    req := httptest.NewRequest(http.MethodPost, "/thema", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    var thema domain.Thema
    if err := json.NewDecoder(rr.Body).Decode(&thema); err != nil {
        t.Fatalf("decode error: %v", err)
    }

    return rr.Code, thema
}

func UpdateThema(t *testing.T, handler http.Handler, sel ThemaSelector, newTitle string) int {
    t.Helper()

    var path string
    switch {
    case sel.ID != nil:
        path = "/thema?id=" + sel.ID.String()
    case sel.Title != nil:
        path = "/thema?title=" + url.QueryEscape(*sel.Title)
    default:
        t.Fatal("selector required")
    }

    body := fmt.Sprintf(`{"title":%q}`, newTitle)

    req := httptest.NewRequest(http.MethodPut, path, strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    return rr.Code
}

func DeleteThema(t *testing.T, handler http.Handler, sel ThemaSelector) int {
    t.Helper()

    var path string
    switch {
    case sel.ID != nil:
        path = "/thema?id=" + sel.ID.String()
    case sel.Title != nil:
        path = "/thema?title=" + url.QueryEscape(*sel.Title)
    default:
        t.Fatal("selector required")
    }

    req := httptest.NewRequest(http.MethodDelete, path, nil)
    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    return rr.Code
}

func GetThema(t *testing.T, handler http.Handler, sel ThemaSelector) (int, domain.Thema) {
    t.Helper()

    var path string
    switch {
    case sel.ID != nil:
        path = "/thema?id=" + sel.ID.String()
    case sel.Title != nil:
        path = "/thema?title=" + url.QueryEscape(*sel.Title)
    default:
        t.Fatal("selector required")
    }

    req := httptest.NewRequest(http.MethodGet, path, nil)
    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    var thema domain.Thema
    if rr.Code == http.StatusOK {
        if err := json.NewDecoder(rr.Body).Decode(&thema); err != nil {
            t.Fatalf("decode failed: %v", err)
        }
    }

    return rr.Code, thema
}
