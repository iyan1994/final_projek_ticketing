package model

import (
	"time"
)

type Ticket struct {
	IdTicket   int       `json:"id_ticket" gorm:"primaryKey"`
	IdUser     int       `json:"id_user"`
	IdCategory int       `json:"id_category"`
	Status     string    `json:"status"`
	NoTiket    string    `json:"no_tiket" gorm:"column:no_tiket"`
	Subjek     string    `json:"subjek"`
	Deksripsi  string    `json:"deksripsi"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type TicketDto struct {
	IdTicket   int       `json:"id_ticket" gorm:"primaryKey"`
	IdUser     int       `json:"id_user"`
	IdCategory int       `json:"id_category"`
	Status     string    `json:"status"`
	NoTiket    string    `json:"no_tiket" gorm:"column:no_tiket"`
	Subjek     string    `json:"subjek"`
	Deksripsi  string    `json:"deksripsi"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *TicketDto) FillFromModel(model Ticket) {
	u.IdTicket = model.IdTicket
	u.IdUser = model.IdUser
	u.IdCategory = model.IdCategory
	u.Status = model.Status
	u.NoTiket = model.NoTiket
	u.Subjek = model.Subjek
	u.Deksripsi = model.Deksripsi
	u.CreatedAt = model.CreatedAt
	u.UpdatedAt = model.UpdatedAt

}

func (u TicketDto) ToModel() Ticket {
	model := Ticket{
		IdTicket:   u.IdTicket,
		IdUser:     u.IdUser,
		IdCategory: u.IdCategory,
		Status:     u.Status,
		NoTiket:    u.NoTiket,
		Subjek:     u.Subjek,
		Deksripsi:  u.Deksripsi,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}

	return model
}
