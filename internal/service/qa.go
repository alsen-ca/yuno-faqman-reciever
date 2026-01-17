package service

import (
    "context"
    "errors"
    "log"
    "fmt"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/db"
    "yuno-faqman-reciever/internal/domain"
)

// Whole question must be unique
var ErrDuplicateQa = errors.New("question already exists")

func CreateQa(ctx context.Context, client *mongo.Client, in domain.QaPayload) (domain.Qa, error) {
    if in.Question == "" || in.Answer == "" {
        return domain.Qa{}, errors.New("Question and Answer are necessary for creation")
    }

    qa := domain.Qa{
        ID: uuid.New(),
        Question: in.Question,
        QuestionWeights: in.QuestionWeights,
        Answer: in.Answer,
        Language: in.Language,
        ThemaID: in.ThemaID,
        TagIDs: in.TagIDs,
    }

    err := db.InsertQa(ctx, client, qa)
    if err != nil {
        if mongo.IsDuplicateKeyError(err) {
            return domain.Qa{}, ErrDuplicateQa
        }
        return domain.Qa{}, err
    }

    return qa, nil
}

func GetQa(ctx context.Context, client *mongo.Client, sel domain.QaSelector) (domain.Qa, error) {
    return db.FindQa(ctx, client, sel)
}

func ListQas(ctx context.Context, client *mongo.Client) ([]domain.Qa, error) {
    return db.ListQas(ctx, client)
}

func UpdateQa(ctx context.Context, client *mongo.Client, id uuid.UUID, payload domain.QaPayload) error {
    if id == uuid.Nil {
        return errors.New("Id must be sent:")
    }
    if payload.Question == "" {
        return errors.New("question cannot be empty")
    }
    if len(payload.QuestionWeights) == 0 {
        return errors.New("question weights cannot be empty")
    }
    if payload.Answer == "" {
        return errors.New("answer cannot be empty")
    }
    if payload.Language == "" {
        return errors.New("language cannot be empty")
    }

    upd := db.QaUpdate{
        Question:       payload.Question,
        QuestionWeights: payload.QuestionWeights,
        Answer:         payload.Answer,
        Language:       payload.Language,
        ThemaID: payload.ThemaID,
        TagIDs: payload.TagIDs,
    }

    err := db.UpdateQa(ctx, client, id, upd)
    if err != nil {
        if mongo.IsDuplicateKeyError(err) {
            return ErrDuplicateQa
        }
        return fmt.Errorf("failed to update QA: %v", err)
    }

    return nil
}

func DeleteQa(ctx context.Context, client *mongo.Client, id uuid.UUID) error {
    log.Printf("Qa id:%s is being deleted", id)
    return db.DeleteQaByID(ctx, client, id)
}
