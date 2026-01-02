package db

import (
    "context"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/google/uuid"

    "yuno-faqman-reciever/internal/domain"
)

func InsertTag(ctx context.Context, client *mongo.Client, tag domain.Tag) error {
    _, err := TagCollection(client).InsertOne(ctx, tag)
    return err
}

func FindTagByID(ctx context.Context, client *mongo.Client, id uuid.UUID) (domain.Tag, error) {
    var tag domain.Tag
    err := TagCollection(client).
        FindOne(ctx, bson.M{"_id": id}).
        Decode(&tag)
    return tag, err
}

func FindTagByEnOriginal(ctx context.Context, client *mongo.Client, en_og string) (domain.Tag, error) {
    var tag domain.Tag
    err := TagCollection(client).
        FindOne(ctx, bson.M{"en_og": en_og}).
        Decode(&tag)
    return tag, err
}

func FindTagByDeTranslation(ctx context.Context, client *mongo.Client, de_trans string) (domain.Tag, error) {
    var tag domain.Tag
    err := TagCollection(client).
        FindOne(ctx, bson.M{"de_trans": de_trans}).
        Decode(&tag)
    return tag, err
}

func FindTagByEsTranslation(ctx context.Context, client *mongo.Client, es_trans string) (domain.Tag, error) {
    var tag domain.Tag
    err := TagCollection(client).
        FindOne(ctx, bson.M{"es_trans": es_trans}).
        Decode(&tag)
    return tag, err
}

func ListTags(ctx context.Context, client *mongo.Client) ([]domain.Tag, error) {
    cursor, err := TagCollection(client).Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var tags []domain.Tag
    err = cursor.All(ctx, &tags)
    return tags, err
}

func UpdateTag(ctx context.Context, client *mongo.Client, id uuid.UUID, en_og string, de_trans string, es_trans string) error {
    res, err := TagCollection(client).UpdateOne(
        ctx,
        bson.M{"_id": id},
        bson.M{"$set": bson.M{"en_og": en_og}},
        bson.M{"$set": bson.M{"de_trans": de_trans}},
        bson.M{"$set": bson.M{"es_trans": es_trans}},
    )
    if err != nil {
        return err
    }
    if res.MatchedCount == 0 {
        return mongo.ErrNoDocuments
    }
    return nil
}

func DeleteTagByID(ctx context.Context, client *mongo.Client, id uuid.UUID) error {
    res, err := TagCollection(client).
        DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        return err
    }

    if res.DeletedCount == 0 {
        return mongo.ErrNoDocuments
    }

    return nil
}
