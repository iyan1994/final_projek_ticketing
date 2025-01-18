package model

import (
	"time"
)

type Feedback struct {
	IdFeedback   int       `json:"id_feedback" gorm:"primaryKey"`
	IdTicket     int       `json:"id_ticket"`
	Satisfaction string    `json:"satisfaction"`
	Deksripsi    string    `json:"deksripsi"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type FeedbackDto struct {
	IdFeedback   int       `json:"id_feedback" gorm:"primaryKey"`
	IdTicket     int       `json:"id_ticket"`
	Satisfaction string    `json:"satisfaction"`
	Deksripsi    string    `json:"deksripsi"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *FeedbackDto) FillFromModel(model Feedback) {
	u.IdFeedback = model.IdFeedback
	u.IdTicket = model.IdTicket
	u.Satisfaction = model.Satisfaction
	u.Deksripsi = model.Deksripsi
	u.CreatedAt = model.CreatedAt
	u.UpdatedAt = model.UpdatedAt

}

func (u FeedbackDto) ToModel() Feedback {
	model := Feedback{
		IdFeedback:   u.IdFeedback,
		IdTicket:     u.IdTicket,
		Satisfaction: u.Satisfaction,
		Deksripsi:    u.Deksripsi,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}

	return model
}
