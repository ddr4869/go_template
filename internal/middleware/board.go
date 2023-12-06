package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/controller"
	"github.com/go-board/internal/dto"
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
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Board) CreateBoard(c *gin.Context) {
	var req dto.ReqCreateBoard
	if err := c.ShouldBind(&req); err != nil {
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Board) DeleteBoard(c *gin.Context) {
	var req dto.ReqDeleteBoard
	if err := c.ShouldBindUri(&req); err != nil {
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Board) PatchBoard(c *gin.Context) {
	var req dto.ReqPatchBoard
	if err := c.ShouldBind(&req); err != nil {
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	if req.Title == "" && req.Content == "" {
		MwAbortJson(c, http.StatusBadRequest, errors.New("please put patch data"))
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Board) RecommendBoard(c *gin.Context) {
	var req dto.ReqRecommendBoard
	if err := c.ShouldBind(&req); err != nil {
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}
