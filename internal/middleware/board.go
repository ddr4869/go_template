package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/controller"
	"github.com/go-board/internal/dto"
	"github.com/the-medium-tech/platform-externals/log"
)

type Board struct {
	board controller.Board
}

func NewBoard(board controller.Board) *Board {
	return &Board{board: board}
}

func (s *Board) UserBoardList(c *gin.Context) {
	var req dto.ReqUserBoardList
	if err := c.ShouldBind(&req); err != nil {
		log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Board) CreateBoard(c *gin.Context) {
	var req dto.ReqCreateBoard
	if err := c.ShouldBind(&req); err != nil {
		log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Board) RecommendBoard(c *gin.Context) {
	var req dto.ReqRecommendBoard
	if err := c.ShouldBind(&req); err != nil {
		log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}
