package model

import (
	"time"
)

type Solution struct {
	IdSolution     int       `json:"id_solution" gorm:"primaryKey"`
	IdAssignTicket int       `json:"id_assign_ticket"`
	Image          string    `json:"image"`
	Deksripsi      string    `json:"deksripsi"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type SolutionDto struct {
	IdSolution     int       `json:"id_solution" gorm:"primaryKey"`
	IdAssignTicket int       `json:"id_assign_ticket"`
	Image          string    `json:"image"`
	Deksripsi      string    `json:"deksripsi"`
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *SolutionDto) FillFromModel(model Solution) {
	u.IdSolution = model.IdSolution
	u.IdAssignTicket = model.IdAssignTicket
	u.Image = model.Image
	u.Deksripsi = model.Deksripsi
	u.CreatedAt = model.CreatedAt
	u.UpdatedAt = model.UpdatedAt

}

func (u SolutionDto) ToModel() Solution {
	model := Solution{
		IdSolution:     u.IdSolution,
		IdAssignTicket: u.IdAssignTicket,
		Image:          u.Image,
		Deksripsi:      u.Deksripsi,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}

	return model
}
