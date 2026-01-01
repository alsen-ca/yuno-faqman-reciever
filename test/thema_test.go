package thema_test

import (
	"context"
	"time"
    "net/http"
    "net/http/httptest"
    "testing"
	"strings"
	"encoding/json"

    "go.mongodb.org/mongo-driver/mongo"

	"yuno-faqman-reciever/internal/http/thema"
	"yuno-faqman-reciever/internal/domain"
	"yuno-faqman-reciever/internal/middleware"
	"yuno-faqman-reciever/internal/db"
)

func setupTestServer(client *mongo.Client) http.Handler {
    mux := http.NewServeMux()
    thema.RegisterRoutes(mux, client)

    return middleware.Logging(mux)
}

func testMongoClient(t *testing.T) (*mongo.Client, *mongo.Database) {
    t.Helper()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := "mongodb://faqman-db:27017"
    client, err := db.ConnectMongo(uri)
    if err != nil {
        t.Fatalf("connect failed: %v", err)
    }

    // Switch to the test DB
    orig := db.DatabaseName
    db.DatabaseName = "faqman_test"
    t.Cleanup(func() {
        db.DatabaseName = orig 
        _ = client.Disconnect(context.Background())
    })

    // Ensure a clean slate
    testDB := client.Database(db.DatabaseName)
    if err = testDB.Drop(ctx); err != nil {
        t.Fatalf("drop failed: %v", err)
    }

    return client, testDB
}

func TestCreateThema(t *testing.T) {
    client, _ := testMongoClient(t)

    handler := setupTestServer(client)

    payload := `{"title":"my Thema"}`
    req := httptest.NewRequest(http.MethodPost, "/thema", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusCreated {
        t.Fatalf("expected 201, got %d", rr.Code)
    }

    var resp domain.Thema
    if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
        t.Fatalf("decode error: %v", err)
    }
    if resp.Title != "my Thema" {
        t.Fatalf("unexpected title: %s", resp.Title)
    }
}
