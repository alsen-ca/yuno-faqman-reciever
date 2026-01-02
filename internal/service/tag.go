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

var ErrDuplicateEnOriginal = errors.New("english original already exists")
var ErrDuplicateDeTranslation = errors.New("german translation already exists")
var ErrDuplicateEsTranslation = errors.New("spanish translation already exists")
var ErrNotFound = errors.New("tag not found")

func CreateTag(ctx context.Context, client *mongo.Client, en_og string, de_trans string, es_trans) (domain.Tag, error) {
    if en_og == "" && de_trans == "" && es_trans == "" {
        return domain.Tag{}, errors.New("tag needs the field in at least one language filled required")
    }

    thema := domain.Tag{
        ID:    uuid.New(),
        EnOg: en_og,
        DeTrans: de_trans,
        EsTrans: es_trans,
    }

    // TODO
    err := db.InsertTag(ctx, client, tag)
    if err != nil {
        if mongo.IsDuplicateKeyError(err) {
            return domain.Tag{}, ErrDuplicateTitle
        }
        return domain.Thema{}, err
    }

    log.Printf("Tag created: id=%s english original=%q german translation=%q spanish translation=%q", tag.ID, tag.EnOg, tag.DeTrans, tag.EsTrans)

    return tag, nil
}

func GetTagByID(ctx context.Context, client *mongo.Client, id uuid.UUID) (domain.Tag, error) {
    return db.FindTagByID(ctx, client, id)
}

func GetTagByEnOriginal(ctx context.Context, client *mongo.Client, en_og string) (domain.Tag, error) {
    return db.FindTagByEnOriginal(ctx, client, en_og)
}

func GetTagByDeTranslation(ctx context.Context, client *mongo.Client, de_trans string) (domain.Tag, error) {
    return db.FindTagByDeTranslation(ctx, client, de_trans)
}

func GetTagByEsTranslation(ctx context.Context, client *mongo.Client, es_trans string) (domain.Tag, error) {
    return db.FindTagByEsTranslation(ctx, client, es_trans)
}

func ListTags(ctx context.Context, client *mongo.Client) ([]domain.Tags, error) {
    return db.ListTags(ctx, client)
}

func UpdateTag(ctx context.Context, client *mongo.Client, id uuid.UUID, en_og string, de_trans string, es_trans string) error {
    if en_og == "" && de_trans == "" && es_trans == "" {
        return errors.New("at least one language field must be filled")
    }
    log.Printf("Tag id:%s will be updated with en: %q, de: %q, es: %q", id, de_trans, es_trans)
    return db.UpdateThemaTitle(ctx, client, id, en_og, de_trans, es_trans)
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
