package dto

type Admin struct {
	Name        string
	Description string
}

type AdminDto struct {
	ID int
}

type ReqCreateAdmin struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description"`
}
