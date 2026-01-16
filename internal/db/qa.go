package db

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/google/uuid"

    "yuno-faqman-reciever/internal/domain"
)

type QaUpdate struct {
    Question string
    QuestionWeights []domain.QuestionWeight
    Answer string
    Language string
}

func InsertQa(ctx context.Context, client *mongo.Client, qa domain.Qa) error {
    _, err := QaCollection(client).InsertOne(ctx, qa)
    return err
}

func FindQa(ctx context.Context, client *mongo.Client, sel domain.QaSelector) (domain.Qa, error) {
    col := client.Database(DatabaseName).Collection("qas")

    filter := bson.M{}

    switch {
    case sel.ID != nil:
        filter["_id"] = *sel.ID
    case sel.Question != nil:
        filter["question"] = *sel.Question
    }
    
    var qa domain.Qa
    err := col.FindOne(ctx, filter).Decode(&qa)
    
    if err != nil {
        return domain.Qa{}, err
    }
    return qa, err
}

func ListQas(ctx context.Context, client *mongo.Client) ([]domain.Qa, error) {
    cursor, err := QaCollection(client).Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var qas []domain.Qa
    err = cursor.All(ctx, &qas)
    return qas, err
}

func UpdateQa(ctx context.Context, client *mongo.Client, id uuid.UUID, qau QaUpdate) error {
    set := bson.M{
        "question":         qau.Question,
        "question_weights": qau.QuestionWeights,
        "answer":           qau.Answer,
        "language":         qau.Language,
    }

    res, err := QaCollection(client).UpdateOne(
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

func DeleteQaByID(ctx context.Context, client *mongo.Client, id uuid.UUID) error {
   res, err := QaCollection(client).
        DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        return err
    }

    if res.DeletedCount == 0 {
        return mongo.ErrNoDocuments
    }

    return nil
}
