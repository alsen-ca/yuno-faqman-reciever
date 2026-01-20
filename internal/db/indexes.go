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

func EnsureTagIndexes(ctx context.Context, client *mongo.Client) error {
    index := []mongo.IndexModel {
        {
            Keys:    bson.M{"en_og": 1},
            Options: options.Index().SetUnique(true).SetName("unique_en_tag").SetPartialFilterExpression(bson.M{"en_og": bson.M{"$exists": true, "$gt": ""}}),
        },
        {
            Keys:    bson.M{"de_trans": 1},
            Options: options.Index().SetUnique(true).SetName("unique_de_tag").SetPartialFilterExpression(bson.M{"de_trans": bson.M{"$exists": true, "$gt": ""}}),
        },
        {
            Keys:    bson.M{"es_trans": 1},
            Options: options.Index().SetUnique(true).SetName("unique_es_tag").SetPartialFilterExpression(bson.M{"es_trans": bson.M{"$exists": true, "$gt": ""}}),
        },
    }
    _, err := TagCollection(client).Indexes().CreateMany(ctx, index)
    return err
}

func EnsureQaIndexes(ctx context.Context, client *mongo.Client) error {
    index := mongo.IndexModel{
        Keys: bson.M{
            "question": 1,
        },
        Options: options.Index().
            SetUnique(true).
            SetName("unique_question"),
    }

    _, err := QaCollection(client).Indexes().CreateOne(ctx, index)
    return err
}