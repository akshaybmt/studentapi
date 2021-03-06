package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type metaData struct {
	Pagination interface{} `json:"pagination"`
}

type GetAllResp struct {
	Data interface{} `json:"data"`
	Meta *metaData   `json:"meta"`
}

func NewDefaultResponse(data interface{}, p interface{}) *GetAllResp {
	var meta *metaData
	if p != nil {
		meta = &metaData{Pagination: p}
	}

	return &GetAllResp{
		Data: data,
		Meta: meta,
	}
}
