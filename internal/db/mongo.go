package db

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    DatabaseName   = "faqman"
    themaCollection = "themas"
    tagCollection = "tags"
)

func ConnectMongo(uri string) (*mongo.Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }

    // Verify connection
    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    return client, nil
}

func ThemaCollection(client *mongo.Client) *mongo.Collection {
    return client.Database(DatabaseName).Collection(themaCollection)
}

func TagCollection(client *mongo.Client) *mongo.Collection {
    return client.Database(DatabaseName).Collection(tagCollection)
}
