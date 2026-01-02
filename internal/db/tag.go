package db

import (
    "context"
    "errors"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/google/uuid"

    "yuno-faqman-reciever/internal/domain"
)

type TagUpdate struct {
    EnOg string
    DeTrans string
    EsTrans string
}

func InsertTag(ctx context.Context, client *mongo.Client, tag domain.Tag) error {
    _, err := TagCollection(client).InsertOne(ctx, tag)
    return err
}

func FindTag(ctx context.Context, client *mongo.Client, sel domain.TagSelector) (domain.Tag, error) {
    col := client.Database(DatabaseName).Collection("tags")

    filter := bson.M{}

    switch {
    case sel.ID != nil:
        filter["_id"] = *sel.ID
    case sel.EnOg != nil:
        filter["en_og"] = *sel.EnOg
    case sel.DeTrans != nil:
        filter["de_trans"] = *sel.DeTrans
    case sel.EsTrans != nil:
        filter["es_trans"] = *sel.EsTrans
    }

    var tag domain.Tag
    err := col.FindOne(ctx, filter).Decode(&tag)
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

func UpdateTag(ctx context.Context, client *mongo.Client, id uuid.UUID, upd TagUpdate) error {
    set := bson.M{}

    if upd.EnOg != "" {
        set["en_og"] = upd.EnOg
    }
    if upd.DeTrans != "" {
        set["de_trans"] = upd.DeTrans
    }
    if upd.EsTrans != "" {
        set["es_trans"] = upd.EsTrans
    }

    if len(set) == 0 {
        return errors.New("no fields to update")
    }

    res, err := TagCollection(client).UpdateOne(
        ctx,
        bson.M{"_id": id},
        bson.M{"$set": set},
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
