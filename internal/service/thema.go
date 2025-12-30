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
