package model

import (
	"database/sql"
	"time"
)

type User struct {
	IdUser       int            `json:"id_user" gorm:"primaryKey"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Name         string         `json:"name"`
	Password     string         `json:"password"`
	Token        sql.NullString `json:"token"`
	ExpiredToken sql.NullTime   `json:"expired_token"`
	IdRole       int            `json:"id_role"`
	Image        sql.NullString `json:"image"`
	Title        string         `json:"title"`
	NoTelepon    int            `json:"no_telepon"`
	Address      string         `json:"address"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

type UserDto struct {
	IdUser       int        `json:"id_user"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	Name         string     `json:"name"`
	Password     string     `json:"password"`
	Token        *string    `json:"token"`
	ExpiredToken *time.Time `json:"expired_token"`
	IdRole       int        `json:"id_role"`
	Image        *string    `json:"Image"`
	Title        string     `json:"title"`
	NoTelepon    int        `json:"no_telepon"`
	Address      string     `json:"address"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *UserDto) FillFromModel(model User) {
	u.IdUser = model.IdUser
	u.Username = model.Username
	u.Email = model.Email
	u.Name = model.Name
	u.Password = model.Password
	if model.Token.Valid {
		u.Token = &model.Token.String
	}
	if model.ExpiredToken.Valid {
		u.ExpiredToken = &model.ExpiredToken.Time
	}
	u.IdRole = model.IdRole
	u.Title = model.Title
	u.NoTelepon = model.NoTelepon
	u.Address = model.Address
	u.CreatedAt = model.CreatedAt
	u.UpdatedAt = model.UpdatedAt
	if model.Image.Valid {
		u.Image = &model.Image.String
	}
}

func (u UserDto) ToModel() User {
	model := User{
		IdUser:    u.IdUser,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Name:      u.Name,
		IdRole:    u.IdRole,
		NoTelepon: u.NoTelepon,
		Title:     u.Title,
		Address:   u.Address,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	if u.Token != nil {
		model.Token.String = *u.Token
		model.Token.Valid = true
	}
	if u.ExpiredToken != nil {
		model.ExpiredToken.Time = *u.ExpiredToken
		model.ExpiredToken.Valid = true
	}
	if u.Image != nil {
		model.Image.String = *u.Image
		model.Image.Valid = true
	}
	return model
}
