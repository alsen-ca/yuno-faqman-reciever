package service

import (
    "context"
    "errors"
    "log"

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
    }

    err := db.InsertQa(ctx, client, qa)
    if err != nil {
        if mongo.IsDuplicateKeyError(err) {
            return domain.Qa{}, ErrDuplicateQa
        }
        return domain.Qa{}, err
    }

    log.Printf("Qa created: id=%s Question=%q Questionweights=%q Answer=%q Language: %q", qa.ID, qa.Question, qa.QuestionWeights, qa.Answer, qa.Language)

    return qa, nil
}/*

func GetTag(ctx context.Context, client *mongo.Client, sel domain.TagSelector) (domain.Tag, error) {
    return db.FindTag(ctx, client, sel)
}

func ListTags(ctx context.Context, client *mongo.Client) ([]domain.Tag, error) {
    return db.ListTags(ctx, client)
}

func UpdateTag(ctx context.Context, client *mongo.Client, id uuid.UUID, payload domain.TagPayload) error {
    if payload.EnOg == "" {
        return errors.New("en_og cannot be empty")
    }

    upd := db.TagUpdate{
        EnOg:     payload.EnOg,
        DeTrans: payload.DeTrans,
        EsTrans: payload.EsTrans,
    }

    return db.UpdateTag(ctx, client, id, upd)
}

func DeleteTag(ctx context.Context, client *mongo.Client, id uuid.UUID) error {
    log.Printf("Tag id:%s is being deleted", id)
    return db.DeleteTagByID(ctx, client, id)
}
*/