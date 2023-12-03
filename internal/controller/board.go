package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/dto"
	"github.com/go-board/internal/service"
	"github.com/the-medium-tech/platform-externals/log"
)

type Board struct {
	board service.Board
}

func NewBoard(board service.Board) *Board {
	return &Board{board: board}
}

// Controller
func (s *Board) BoardList(c *gin.Context) {
	boards, err := s.board.BoardList(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, boards)
}

func (s *Board) UserBoardList(c *gin.Context) {

	token := c.MustGet("token").(*dto.AccessClaims)
	log.Infof("token: %+v", token)
	boards, err := s.board.UserBoardList(c, token.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, boards)
}

func (s *Board) CreateBoard(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqCreateBoard)
	token := c.MustGet("token").(*dto.AccessClaims)
	board, err := s.board.CreateBoard(c, req.Title, req.Content, token.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, board)
}

func (s *Board) ReadBoard(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqRecommendBoard)
	board, err := s.board.ReadBoard(c, req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, board)
}

func (s *Board) DeleteBoard(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqCreateBoard)
	token := c.MustGet("token").(*dto.AccessClaims)
	board, err := s.board.CreateBoard(c, req.Title, req.Content, token.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, board)
}

func (s *Board) RecommendBoard(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqRecommendBoard)
	a8m, err := s.board.RecommendBoard(c, req.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, a8m)
}

func (s *Board) HotBoardList(c *gin.Context) {
	hotBoards, err := s.board.HotBoardList(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, hotBoards)
}
