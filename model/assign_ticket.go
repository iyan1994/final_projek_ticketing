package model

import (
	"database/sql"
	"time"
)

type TablerAssignTicket interface {
	TableName() string
}

func (AssignTicket) TableName() string {
	return "assign_tickets"
}

type AssignTicket struct {
	IdAssignTicket int            `json:"id_assign_ticket" gorm:"primaryKey"`
	IdTicket       int            `json:"id_ticket"`
	IdPriority     int            `json:"id_priority"`
	IdTeknisi      int            `json:"id_teknisi"`
	IdAdmin        int            `json:"id_admin"`
	StartTicket    sql.NullTime   `json:"start_ticket"`
	CloseTicket    sql.NullTime   `json:"close_ticket"`
	FinishTime     sql.NullString `json:"finish_time" gorm:"column:finish_time"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

type AssignTicketDto struct {
	IdAssignTicket int        `json:"id_assign_ticket" gorm:"primaryKey"`
	IdTicket       int        `json:"id_ticket"`
	IdPriority     int        `json:"id_priority"`
	IdTeknisi      int        `json:"id_teknisi"`
	IdAdmin        int        `json:"id_admin"`
	StartTicket    *time.Time `json:"start_ticket"`
	CloseTicket    *time.Time `json:"close_ticket"`
	FinishTime     *string    `json:"finish_time" gorm:"column:finish_time"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *AssignTicketDto) FillFromModel(model AssignTicket) {
	u.IdAssignTicket = model.IdAssignTicket
	u.IdTicket = model.IdTicket
	u.IdPriority = model.IdPriority
	u.IdTeknisi = model.IdTeknisi
	u.IdAdmin = model.IdAdmin
	if model.StartTicket.Valid {
		u.StartTicket = &model.StartTicket.Time
	}
	if model.CloseTicket.Valid {
		u.CloseTicket = &model.CloseTicket.Time
	}
	if model.FinishTime.Valid {
		u.FinishTime = &model.FinishTime.String
	}

	u.CreatedAt = model.CreatedAt
	u.UpdatedAt = model.UpdatedAt

}

func (u AssignTicketDto) ToModel() AssignTicket {
	model := AssignTicket{
		IdAssignTicket: u.IdAssignTicket,
		IdTicket:       u.IdTicket,
		IdPriority:     u.IdPriority,
		IdTeknisi:      u.IdTeknisi,
		IdAdmin:        u.IdAdmin,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
	if u.StartTicket != nil {
		model.StartTicket.Time = *u.StartTicket
		model.StartTicket.Valid = true
	}
	if u.CloseTicket != nil {
		model.CloseTicket.Time = *u.CloseTicket
		model.CloseTicket.Valid = true
	}
	if u.FinishTime != nil {
		model.FinishTime.String = *u.FinishTime
		model.FinishTime.Valid = true
	}

	return model
}
