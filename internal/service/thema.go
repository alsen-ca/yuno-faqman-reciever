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

var ErrDuplicateTitle = errors.New("title already exists")
var ErrNotFound = errors.New("thema not found")

func CreateThema(ctx context.Context, client *mongo.Client, title string) (domain.Thema, error) {
    if title == "" {
        return domain.Thema{}, errors.New("title required")
    }

    thema := domain.Thema{
        ID:    uuid.New(),
        Title: title,
    }

    err := db.InsertThema(ctx, client, thema)
    if err != nil {
        if mongo.IsDuplicateKeyError(err) {
            return domain.Thema{}, ErrDuplicateTitle
        }
        return domain.Thema{}, err
    }

    log.Printf("Thema created: id=%s title=%q", thema.ID, thema.Title)

    return thema, nil
}

func GetThemaByID(ctx context.Context, client *mongo.Client, id uuid.UUID) (domain.Thema, error) {
    return db.FindThemaByID(ctx, client, id)
}

func GetThemaByTitle(ctx context.Context, client *mongo.Client, title string) (domain.Thema, error) {
    return db.FindThemaByTitle(ctx, client, title)
}

func ListThemas(ctx context.Context, client *mongo.Client) ([]domain.Thema, error) {
    return db.ListThemas(ctx, client)
}

func UpdateThemaTitle(ctx context.Context, client *mongo.Client, id uuid.UUID, title string) error {
    if title == "" {
        return errors.New("title required")
    }
    log.Printf("Thema id:%s updated to title:%q", id, title)
    return db.UpdateThemaTitle(ctx, client, id, title)
}

func UpdateThemaTitleByTitle(ctx context.Context, client *mongo.Client, oldTitle string, newTitle string) error {
    if newTitle == "" {
        return errors.New("title required")
    }
    log.Printf("Updating thema: %s to: %q", oldTitle, newTitle)
    return db.UpdateThemaByTitle(ctx, client, oldTitle, newTitle)
}

func DeleteThema(ctx context.Context, client *mongo.Client, id uuid.UUID, title string) error {
    switch {
    case id != uuid.Nil:
        log.Printf("Thema id:%s is being deleted", id)
        return db.DeleteThemaByID(ctx, client, id)
    case title != "":
        log.Printf("Thema with title: %s has been deleted", title)
        return db.DeleteThemaByTitle(ctx, client, title)
    default:
        return errors.New("no selector")
    }
}
