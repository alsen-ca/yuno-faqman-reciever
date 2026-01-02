package domain

import "github.com/google/uuid"

type Tag struct {
    ID uuid.UUID `json:"id" bson:"_id"`
    EnOg string `json:"en_og" bson:"en_og"`
    DeTrans string `json:"de_trans" bson:"de_trans"`
    EsTrans string `json:"es_trans" bson:"es_trans"`
}
