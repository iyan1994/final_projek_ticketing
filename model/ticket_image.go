package model

import (
	"time"
)

type TablerTicketImage interface {
	TableName() string
}

func (TicketImage) TableName() string {
	return "image_tickets"
}

type TicketImage struct {
	IdImageTicket int       `json:"id_image_ticket" gorm:"primaryKey"`
	IdTicket      int       `json:"id_ticket"`
	Image         string    `json:"image"`
	Deksripsi     string    `json:"deksripsi"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type TicketImageDto struct {
	IdImageTicket int       `json:"id_image_ticket" gorm:"primaryKey"`
	IdTicket      int       `json:"id_ticket"`
	Image         string    `json:"image"`
	Deksripsi     string    `json:"deksripsi"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *TicketImageDto) FillFromModel(model TicketImage) {
	u.IdImageTicket = model.IdImageTicket
	u.IdTicket = model.IdTicket
	u.Image = model.Image
	u.Deksripsi = model.Deksripsi
	u.CreatedAt = model.CreatedAt
	u.UpdatedAt = model.UpdatedAt

}

func (u TicketImageDto) ToModel() TicketImage {
	model := TicketImage{
		IdImageTicket: u.IdImageTicket,
		IdTicket:      u.IdTicket,
		Image:         u.Image,
		Deksripsi:     u.Deksripsi,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}

	return model
}
