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

var ErrDuplicateTag = errors.New("tag already exists")

func CreateTag(ctx context.Context, client *mongo.Client, in domain.TagCreate) (domain.Tag, error) {
    if in.EnOg == "" {
        return domain.Tag{}, errors.New("English Original is required for creation")
    }

    tag := domain.Tag{
        ID: uuid.New(),
        EnOg: in.EnOg,
        DeTrans: in.DeTrans,
        EsTrans: in.EsTrans,
    }

    err := db.InsertTag(ctx, client, tag)
    if err != nil {
        if mongo.IsDuplicateKeyError(err) {
            return domain.Tag{}, ErrDuplicateTag
        }
        return domain.Tag{}, err
    }

    log.Printf("Tag created: id=%s english original=%q german translation=%q spanish translation=%q", tag.ID, tag.EnOg, tag.DeTrans, tag.EsTrans)

    return tag, nil
}

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
    switch {
    case id != uuid.Nil:
        log.Printf("Tag id:%s is being deleted", id)
        return db.DeleteTagByID(ctx, client, id)
    default:
        return errors.New("no selector")
    }
}
