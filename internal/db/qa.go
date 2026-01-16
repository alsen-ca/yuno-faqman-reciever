package db

import (
    "context"

    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/domain"
)


func InsertQa(ctx context.Context, client *mongo.Client, qa domain.Qa) error {
    _, err := QaCollection(client).InsertOne(ctx, qa)
    return err
}
/*
func FindTag(ctx context.Context, client *mongo.Client, sel domain.QaSelector) (domain.Qa, error) {
    col := client.Database(DatabaseName).Collection("qa")

    filter := bson.M{}

    switch {
    case sel.ID != nil:
        filter["_id"] = *sel.ID
    case sel.Question != nil:
        filter["question"] = *sel.Question
    }

    var qa domain.Qa
    err := col.FindOne(ctx, filter).Decode(&qa)
    return qa, err
}

func ListQas(ctx context.Context, client *mongo.Client) ([]string, error) {
    cursor, err := QaCollection(client).Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var qas []domain.Qa
    err = cursor.All(ctx, &qas)
    return qas, err
}

func UpdateQa(ctx context.Context, client *mongo.Client, id uuid.UUID, qaUpd QaUpdate) error {
    set := bson.M{}

    if qaUpd.Question == "" || qaUpd.QuestionWeights == "" || qaUpd.Answer == "" || qaUpd.Language == "" {
    	return errors.New("some fileds were sent empty. Update failed"
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
*/