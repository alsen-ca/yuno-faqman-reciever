package domain

import (
	"github.com/google/uuid"
)

type Qa struct {
    ID uuid.UUID `bson:"_id"`
    Question string `bson:"question"`
    QuestionWeights []QuestionWeight `bson:"question_weights"`
    Answer string `bson:"answer"`
    Language string `bson:"lang"`
    ThemaID uuid.UUID `bson:"thema_id"`
    TagIDs []uuid.UUID `bson:"tag_ids"`
}

type QaPayload struct {
    Question string `json:"question"`
    QuestionWeights []QuestionWeight `json:"question_weights"`
    Answer string `json:"answer"`
    Language string `json:"lang"`
    ThemaID uuid.UUID `json:"thema_id"`
    TagIDs []uuid.UUID `json:"tag_ids"`
}

type QuestionWeight struct {
    Word   string  `json:"word"`
    Weight float64 `json:"weight"`
}

type QaSelector struct {
    ID *uuid.UUID
    Question *string
}
