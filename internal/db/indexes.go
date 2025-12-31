package db

import (
    "context"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func EnsureThemaIndexes(ctx context.Context, client *mongo.Client) error {
    index := mongo.IndexModel{
        Keys: bson.M{
            "title": 1,
        },
        Options: options.Index().
            SetUnique(true).
            SetName("unique_title"),
    }

    _, err := ThemaCollection(client).Indexes().CreateOne(ctx, index)
    return err
}
