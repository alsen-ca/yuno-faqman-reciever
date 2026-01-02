package testutil

import (
    "context"
    "testing"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "yuno-faqman-reciever/internal/db"
)

func TestMongoClient(t *testing.T) (*mongo.Client, *mongo.Database) {
    t.Helper()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := "mongodb://faqman-db:27017"
    client, err := db.ConnectMongo(uri)
    if err != nil {
        t.Fatalf("connect failed: %v", err)
    }

    orig := db.DatabaseName
    db.DatabaseName = "faqman_test"

    t.Cleanup(func() {
        db.DatabaseName = orig
        _ = client.Disconnect(context.Background())
    })

    testDB := client.Database(db.DatabaseName)
    if err := testDB.Drop(ctx); err != nil {
        t.Fatalf("drop failed: %v", err)
    }

    return client, testDB
}
