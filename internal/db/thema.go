package db

import (
    "context"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/google/uuid"

    "yuno-faqman-reciever/internal/domain"
)

func InsertThema(ctx context.Context, client *mongo.Client, thema domain.Thema) error {
    _, err := ThemaCollection(client).InsertOne(ctx, thema)
    return err
}

func FindThemaByID(ctx context.Context, client *mongo.Client, id uuid.UUID) (domain.Thema, error) {
    var thema domain.Thema
    err := ThemaCollection(client).
        FindOne(ctx, bson.M{"_id": id}).
        Decode(&thema)
    return thema, err
}

func FindThemaByTitle(ctx context.Context, client *mongo.Client, title string) (domain.Thema, error) {
    var thema domain.Thema
    err := ThemaCollection(client).
        FindOne(ctx, bson.M{"title": title}).
        Decode(&thema)
    return thema, err
}

func ListThemas(ctx context.Context, client *mongo.Client) ([]domain.Thema, error) {
    cursor, err := ThemaCollection(client).Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var themas []domain.Thema
    err = cursor.All(ctx, &themas)
    return themas, err
}

func UpdateThemaTitle(ctx context.Context, client *mongo.Client, id uuid.UUID, title string) error {
    res, err := ThemaCollection(client).UpdateOne(
        ctx,
        bson.M{"_id": id},
        bson.M{"$set": bson.M{"title": title}},
    )
    if err != nil {
        return err
    }
    if res.MatchedCount == 0 {
        return mongo.ErrNoDocuments
    }
    return nil
}

func UpdateThemaByTitle(ctx context.Context, client *mongo.Client, oldTitle, newTitle string) error {
    res, err := ThemaCollection(client).UpdateOne(
        ctx,
        bson.M{"title": oldTitle},
        bson.M{"$set": bson.M{"title": newTitle}},
    )
    if err != nil {
        return err
    }
    if res.MatchedCount == 0 {
        return mongo.ErrNoDocuments
    }
    return nil
}

func DeleteThemaByID(ctx context.Context, client *mongo.Client, id uuid.UUID) error {
    res, err := ThemaCollection(client).
        DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        return err
    }

    if res.DeletedCount == 0 {
        return mongo.ErrNoDocuments
    }

    return nil
}

func DeleteThemaByTitle(ctx context.Context, client *mongo.Client, title string) error {
    res, err := ThemaCollection(client).
        DeleteOne(ctx, bson.M{"title": title})
    if err != nil {
        return err
    }
    if res.DeletedCount == 0 {
        return mongo.ErrNoDocuments
    }
    return nil
}
