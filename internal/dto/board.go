package dto

import "time"

type Board struct {
	ID          uint
	Title       string
	Content     string
	Writer      string
	View        uint
	Recommend   uint
	Hot         bool
	CreatedDate time.Time
}

type BoardDto struct {
	ID int
}

type ReqCreateBoard struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content"`
	//Writer  string `form:"writer" json:"writer" binding:"required"`
}

type ReqDeleteBoard struct {
	ID      string `form:"id" json:"id" binding:"required"`
	Content string `form:"content" json:"content"`
	Writer  string `form:"writer" json:"writer" binding:"required"`
}

type ReqUserBoardList struct {
	Writer string `form:"writer" json:"writer" binding:"required"`
}

type ReqRecommendBoard struct {
	ID int `form:"id" json:"id" binding:"required"`
}
