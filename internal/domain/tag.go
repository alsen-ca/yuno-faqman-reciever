package domain

import "github.com/google/uuid"

type Tag struct {
    ID uuid.UUID `json:"id" bson:"_id"`
    EnOg string `json:"en_og" bson:"en_og"`
    DeTrans string `json:"de_trans" bson:"de_trans"`
    EsTrans string `json:"es_trans" bson:"es_trans"`
}

type TagPayload struct {
    EnOg string `json:"en_og"`
    DeTrans string `json:"de_trans,omitempty"`
    EsTrans string `json:"es_trans,omitempty"`
}

type TagCreate struct {
    EnOg string
    DeTrans string
    EsTrans string
}

type TagSelector struct {
    ID *uuid.UUID
    EnOg *string
    DeTrans *string
    EsTrans *string
}

func (p TagPayload) ToDomain() TagCreate {
    return TagCreate{
        EnOg: p.EnOg,
        DeTrans: p.DeTrans,
        EsTrans: p.EsTrans,
    }
}