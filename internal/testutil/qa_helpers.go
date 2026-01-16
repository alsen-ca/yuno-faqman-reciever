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
/*
type TagSelector struct {
    ID    *uuid.UUID
    EnOg *string
    DeTrans *string
    EsTrans *string
}

func TagByID(id uuid.UUID) TagSelector {
    return TagSelector{
        ID: &id,
    }
}

func TagByEn(en_og string) TagSelector {
    return TagSelector{
        EnOg: &en_og,
    }
}

func TagByDe(de_trans string) TagSelector {
    return TagSelector{
        DeTrans: &de_trans,
    }
}
func TagByEs(es_trans string) TagSelector {
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

func UpdateTag(t *testing.T, handler http.Handler, sel TagSelector, payload domain.Tag) int {
    t.Helper()

    if sel.ID == nil {
        t.Fatal("UpdateTag requires a selector with ID")
    }
    path := "/tag?id=" + sel.ID.String()

    body, err := json.Marshal(payload)
    if err != nil {
        t.Fatalf("failed to marshal payload: %v", err)
    }

    req := httptest.NewRequest(http.MethodPut, path, strings.NewReader(string(body)))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    return rr.Code
}

func DeleteTag(t *testing.T, handler http.Handler, id uuid.UUID) int {
	t.Helper()
	path := "/tag?id=" + url.QueryEscape(id.String())

	req := httptest.NewRequest(http.MethodDelete, path, nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	return rr.Code
}


func GetTag(t *testing.T, handler http.Handler, sel TagSelector) (int, domain.Tag) {
	t.Helper()

	var path string
	switch {
	case sel.ID != nil:
		path = "/tag?id=" + sel.ID.String()
	case sel.EnOg != nil:
		path = "/tag?en_og=" + url.QueryEscape(*sel.EnOg)
	case sel.DeTrans != nil:
		path = "/tag?de_trans=" + url.QueryEscape(*sel.DeTrans)
	case sel.EsTrans != nil:
		path = "/tag?es_trans=" + url.QueryEscape(*sel.EsTrans)
	default:
		t.Fatal("selector required")
	}

	req := httptest.NewRequest(http.MethodGet, path, nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	var tg domain.Tag
	if rr.Code == http.StatusOK {
		if err := json.NewDecoder(rr.Body).Decode(&tg); err != nil {
			t.Fatalf("decode failed: %v", err)
		}
	}
	return rr.Code, tg
}
*/