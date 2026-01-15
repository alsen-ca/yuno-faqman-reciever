package tag

import (
    "testing"
    "net/http"

    "yuno-faqman-reciever/internal/testutil"
    "yuno-faqman-reciever/internal/domain"
)

func TestCreateTag(t *testing.T) {
    client, _ := testutil.TestMongoClient(t)
    handler := testutil.SetupTestServer(func(mux *http.ServeMux) {
        RegisterRoutes(mux, client)
    })

    enOg := "Test Tag"
    deTrans := "Test-Tag"
    esTrans := "Tag de prueba"
    code, tagDomain := testutil.CreateTagHTTP(t, handler, enOg, deTrans, esTrans)
    if code != http.StatusCreated {
        t.Fatalf("unexpected code for creation: %d", code)
    }

    // Verify the tag was created correctly
    codeGet, tagGet := testutil.GetTag(t, handler, testutil.TagByID(tagDomain.ID))
    if codeGet != http.StatusOK {
        t.Fatalf("unexpected code for GET: %d", codeGet)
    }
    if tagGet.EnOg != enOg || tagGet.DeTrans != deTrans || tagGet.EsTrans != esTrans {
        t.Fatalf("tag fields from GET do not match: %+v", tagGet)
    }

    updatedTag := domain.Tag{EnOg: "Changed Tag", DeTrans: "Geänderter Tag", EsTrans: "Tag cambiado"}
    codeUpdate := testutil.UpdateTag(t, handler, testutil.TagByID(tagDomain.ID), updatedTag)
    if codeUpdate != http.StatusNoContent {
        t.Fatalf("unexpected code for PUT: %d", codeUpdate)
    }

    // Verify the old tag is no longer accessible by its old fields
    codeGetOld, _ := testutil.GetTag(t, handler, testutil.TagByEn(enOg))
    if codeGetOld == http.StatusOK {
        t.Fatalf("GET request for old EnOg should not have been valid: %d", codeGetOld)
    }

    // By ID
    codeGetNew, tagGetNew := testutil.GetTag(t, handler, testutil.TagByID(tagDomain.ID))
    if codeGetNew != http.StatusOK {
        t.Fatalf("unexpected code for GET: %d", codeGetNew)
    }
    if tagGetNew.EnOg != "Changed Tag" || tagGetNew.DeTrans != "Geänderter Tag" || tagGetNew.EsTrans != "Tag cambiado" {
        t.Fatalf("tag fields from GET do not match updated values: %+v", tagGetNew)
    }

    codeGetNew2, tagGetNew2 := testutil.GetTag(t, handler, testutil.TagByEn("Changed Tag"))
    if codeGetNew2 != http.StatusOK {
        t.Fatalf("unexpected code for GET: %d", codeGetNew2)
    }
    
    if tagGetNew2 != tagGetNew {
        t.Fatalf("tag fields dont match: %+v", tagGetNew2)
    }
}


func TestDeleteTag(t *testing.T) {
	client, _ := testutil.TestMongoClient(t)
    handler := testutil.SetupTestServer(func(mux *http.ServeMux) {
        RegisterRoutes(mux, client)
    })

	_, tag := testutil.CreateTagHTTP(t, handler, "to-be-deleted", "zu löschen", "para borrar")

	getCode, fetched := testutil.GetTag(t, handler, testutil.TagByID(tag.ID))
	if getCode != http.StatusOK {
		t.Fatalf("couldn't fetch newly created tag, got %d", getCode)
	}
	if fetched.EnOg != "to-be-deleted" {
		t.Fatalf("fetched tag does not match created one")
	}

	delCode := testutil.DeleteTag(t, handler, tag.ID)
	if delCode != http.StatusNoContent {
		t.Fatalf("expected 204 on delete, got %d", delCode)
	}

	afterDelCode, _ := testutil.GetTag(t, handler, testutil.TagByID(tag.ID))
	if afterDelCode != http.StatusNotFound {
		t.Fatalf("expected 404 after delete, got %d", afterDelCode)
	}
}