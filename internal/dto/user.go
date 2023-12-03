package dto

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Name         string
	Description  string
	AccessToken  string
	RefreshToken string
}

type UserDto struct {
	ID          int
	Name        string
	Description string
	CreatedDate time.Time
}

type ReqCreateUser struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	Description string `form:"description" json:"description"`
}

type ReqLoginUser struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AccessClaims struct {
	Username   string
	ExpireTime float64
}

type RefreshClaims struct {
	UuID     uuid.UUID
	Username string
}
