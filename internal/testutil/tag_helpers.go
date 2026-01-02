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

type TagSelector struct {
    ID    *uuid.UUID
    EnOg *string
    DeTrans *string
    EsTrans *string
}

func ByID(id uuid.UUID) TagSelector {
    return TagSelector{
        ID: &id,
    }
}

func ByEn(en_og string) TagSelector {
    return TagSelector{
        EnOg: &en_og,
    }
}

func ByDe(de_trans string) TagSelector {
    return TagSelector{
        DeTrans: &de_trans,
    }
}
func ByEs(es_trans string) TagSelector {
    return TagSelector{
        EsTrans: &es_trans,
    }
}

func CreateTagHTTP(t *testing.T, handler http.Handler, en_og string, de_trans string, es_trans string) (int, domain.Tag) {
    t.Helper()

    payload := fmt.Sprintf(`{"en_og":"%s", "de_trans":"%s", "es_trans":"%s"}`, en_og, de_trans, es_trans)
    req := httptest.NewRequest(http.MethodPost, "/tag", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    var tag domain.Tag
    if err := json.NewDecoder(rr.Body).Decode(&tag); err != nil {
        t.Fatalf("decode error: %v", err)
    }

    return rr.Code, tag
}

func UpdateTag(t *testing.T, handler http.Handler, sel TagSelector, newTitle string) int {
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
