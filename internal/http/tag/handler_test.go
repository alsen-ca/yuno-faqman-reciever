package tag

import (
    "testing"
    "net/http"

    "yuno-faqman-reciever/internal/testutil"
)

func TestCreateTag(t *testing.T) {
    client, _ := testutil.TestMongoClient(t)
    handler := testutil.SetupTestServer(func(mux *http.ServeMux) {
        RegisterRoutes(mux, client)
    })

    tagCode, tag := testutil.CreateTagHTTP(t, handler, "my Tag", "mein Tag", "mi Tag")
    _, tag2 := testutil.CreateTagHTTP(t, handler, "second Tag", "zweites Tag", "segundo Tag")

    if tag.EnOg != "my Tag" {
        t.Fatalf("unexpected English Original: %s", tag.EnOg)
    }
    if tag.DeTrans != "mein Tag" {
        t.Fatalf("unexpected German Translation: %s", tag.DeTrans)
    }
    if tag.EsTrans != "mi Tag" {
        t.Fatalf("unexpected Spanish Translation: %s", tag.EsTrans)
    }
    if tagCode != http.StatusCreated {
        t.Fatalf("unexpected code for creation: %d", tagCode)
    }
    if tag2.EnOg != "second Tag" {
        t.Fatalf("unexpected English Original: %s", tag.EnOg)
    }
    if tag.EnOg == tag2.EnOg {
        t.Fatalf("Tags have the same title")
    }
}

// TODO
func TestGetTag(t *testing.T) {
    client, _ := testutil.TestMongoClient(t)
    handler := testutil.SetupTestServer(func(mux *http.ServeMux) {
        RegisterRoutes(mux, client)
    })

    themaCode, themaDomain := testutil.CreateThemaHTTP(t, handler, "Test Thema")

    if themaDomain.Title != "Test Thema" {
        t.Fatalf("unexpected title: %s", themaDomain.Title)
    } else if themaCode != http.StatusCreated {
        t.Fatalf("unexpected code for creation: %d", themaCode)
    }

    code, thema := testutil.GetThema(t, handler, testutil.ByID(themaDomain.ID))

    if code != http.StatusOK {
        t.Fatalf("unexpected code for GET: %d", code)
    }
    if thema.Title != themaDomain.Title {
        t.Fatalf("Title from GET was not the same as creation title: %s", thema.Title)
    }

    code2, thema2 := testutil.GetThema(t, handler, testutil.ByTitle(themaDomain.Title))

    if code2 != http.StatusOK {
        t.Fatalf("unexpected code for GETL %d", code2)
    }
    if thema2.Title != thema.Title {
        t.Fatalf("Titles of the same thema do not match: %s", thema2.Title)
    }
}

func TestModifyThema(t *testing.T) {
    client, _ := testutil.TestMongoClient(t)
    handler := testutil.SetupTestServer(func(mux *http.ServeMux) {
        RegisterRoutes(mux, client)
    })
    _, themaDomain := testutil.CreateThemaHTTP(t, handler, "Test Thema")
    if themaDomain.Title != "Test Thema" {
        t.Fatalf("unexpected title: %s", themaDomain.Title)
    }

    codeGet, themaGet := testutil.GetThema(t, handler, testutil.ByTitle(themaDomain.Title))
    if codeGet != http.StatusOK {
        t.Fatalf("unexpected code for GET: %d", codeGet)
    }
    if themaGet.Title != "Test Thema" {
        t.Fatalf("Titles of the same thema do not match: %s", themaGet.Title)
    }

    codeUpdate := testutil.UpdateThema(t, handler, testutil.ByTitle(themaDomain.Title), "Changed Title")
    if codeUpdate != http.StatusNoContent {
        t.Fatalf("unexpected code for GET, should be 204    : %d", codeUpdate)
    }

    codeGetNew, _ := testutil.GetThema(t, handler, testutil.ByTitle("Test Thema"))
    if codeGetNew == http.StatusOK {
        t.Fatalf("GET request should not have been valid: %d", codeGetNew)
    }

    codeGetNew2, themaGetNew2 := testutil.GetThema(t, handler, testutil.ByID(themaDomain.ID))
    if codeGetNew2 != http.StatusOK {
        t.Fatalf("Unexpected code for GET: %d", codeGetNew2)
    }
    if themaGetNew2.Title != "Changed Title" {
        t.Fatalf("Unexpected title for GET: %s", themaGetNew2.Title)
    }

    codeGetNew3, themaGetNew3 := testutil.GetThema(t, handler, testutil.ByTitle("Changed Title"))
    if codeGetNew3 != http.StatusOK {
        t.Fatalf("Unexpected code for GET: %d", codeGetNew3)
    }
    if themaGetNew3.Title != "Changed Title" {
        t.Fatalf("Unexpected title for GET: %s", themaGetNew3.Title)
    }
}

func TestDeleteThema(t *testing.T) {
    client, _ := testutil.TestMongoClient(t)
    handler := testutil.SetupTestServer(func(mux *http.ServeMux) {
        RegisterRoutes(mux, client)
    })

    _, ogThema := testutil.CreateThemaHTTP(t, handler, "Test Thema")
    code, thema := testutil.GetThema(t, handler, testutil.ByID(ogThema.ID))

    if code != http.StatusOK {
        t.Fatalf("unexpected code for GET: %d", code)
    }
    if thema.Title != "Test Thema" {
        t.Fatalf("Title from GET was not the same as creation title: %s", thema.Title)
    }

    codeDelete := testutil.DeleteThema(t, handler, testutil.ByTitle(ogThema.Title))
    if codeDelete != http.StatusNoContent {
        t.Fatalf("unexpected code for GET, should have 204: %d", codeDelete)
    }

    codeGet2, _ := testutil.GetThema(t, handler, testutil.ByID(ogThema.ID))
    if codeGet2 != http.StatusNotFound {
        t.Fatalf("unexpected code for GET, should have been 405: %d", codeGet2)
    }
}