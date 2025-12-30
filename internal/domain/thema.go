package domain

import "github.com/google/uuid"

type Thema struct {
    ID uuid.UUID `json:"id" bson:"_id"`
    Title string `json:"title" bson:"title"`
}
