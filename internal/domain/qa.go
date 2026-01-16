package domain

import (
	"github.com/google/uuid"
)

type Qa struct {
    ID uuid.UUID `bson:"_id"`
    Question string `bson:"question"`
    QuestionWeights []QuestionWeight `bson:"question_weights`
    Answer string `bson:"answer"`
    Language string `bson:"lang"`
}

type QaPayload struct {
    Question string `json:"question"`
    QuestionWeights []QuestionWeight `json:"question_weights"`
    Answer string `json:"answer"`
    Language string `json:"lang"`
}

type QuestionWeight struct {
    Word   string  `json:"word"`
    Weight float64 `json:"weight"`
}
