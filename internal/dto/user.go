package dto

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type User struct {
	Name         string
	Description  string
	AccessToken  string
	RefreshToken string
}

type UserDto struct {
	ID int
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
	UuID     uuid.UUID
	Username string
	jwt.StandardClaims
}

type RefreshClaims struct {
	UuID     uuid.UUID
	Username string
}
