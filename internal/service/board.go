package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-board/ent"
	"github.com/go-board/internal/dto"
	"github.com/go-board/internal/repository"
)

type Board struct {
	board repository.Board
}

func NewBoard(board repository.Board) *Board {
	return &Board{board: board}
}

func (s *Board) BoardList(c *gin.Context) ([]dto.Board, error) {
	boards, err := s.board.BoardList(c)
	if err != nil {
		return nil, err
	}
	return convertBoardArrayEntToDto(boards), nil
}

func (s *Board) UserBoardList(c *gin.Context, writer string) ([]dto.Board, error) {
	boards, err := s.board.UserBoardList(c, writer)
	if err != nil {
		return nil, err
	}
	return convertBoardArrayEntToDto(boards), nil
}

func (s *Board) CreateBoard(c *gin.Context, title, content, writer string) (*dto.Board, error) {
	board, err := s.board.CreateBoard(c, title, content, writer)
	if err != nil {
		return nil, err
	}
	return convertBoardEntToDto(board), nil
}

func (s *Board) ReadBoard(c *gin.Context, boardId int) (*dto.Board, error) {
	board, err := s.board.ReadBoard(c, boardId)
	if err != nil {
		return nil, err
	}
	return convertBoardEntToDto(board), nil
}

func (s *Board) DeleteBoard(c *gin.Context, boardId int, name string) error {

	board, err := s.board.ReadBoard(c, boardId)
	if err != nil {
		return err
	}
	if board.Writer != name {
		return errors.New("User is not writer")
	}

	err = s.board.DeleteBoard(c, boardId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Board) PatchBoard(c *gin.Context, req dto.ReqPatchBoard, name string) (*dto.Board, error) {

	board, err := s.board.ReadBoard(c, req.ID)
	if err != nil {
		return nil, err
	}
	if board.Writer != name {
		return nil, errors.New("User is not writer")
	}

	patchBoard, err := s.board.PatchBoard(c, req.ID, req.Title, req.Content)
	if err != nil {
		return nil, err
	}
	return convertBoardEntToDto(patchBoard), nil
}

func (s *Board) RecommendBoard(c *gin.Context, boardId int) (int, error) {
	recommendCnt, err := s.board.RecommendBoard(c, boardId)
	if err != nil {
		return -1, err
	}

	if recommendCnt == 10 {
		err = s.board.TurnToHotBoard(c, boardId)
		if err != nil {
			return -1, err
		}
	}
	return recommendCnt, nil
}

func (s *Board) HotBoardList(c *gin.Context) ([]dto.Board, error) {
	hotBoards, err := s.board.HotBoardList(c)
	if err != nil {
		return nil, err
	}
	return convertBoardArrayEntToDto(hotBoards), nil
}

func convertBoardEntToDto(entBoards *ent.Board) *dto.Board {
	dtoBoard := dto.Board{
		ID:          uint(entBoards.ID),
		Title:       entBoards.Title,
		Content:     entBoards.Content,
		Writer:      entBoards.Writer,
		View:        entBoards.View,
		Recommend:   entBoards.Recommend,
		Hot:         entBoards.Hot,
		CreatedDate: entBoards.CreatedDate,
	}

	return &dtoBoard
}

func convertBoardArrayEntToDto(entBoards []*ent.Board) []dto.Board {
	dto_boards := make([]dto.Board, 0)
	for _, v := range entBoards {
		dto_board := dto.Board{
			ID:          v.ID,
			Title:       v.Title,
			Content:     v.Content,
			Writer:      v.Writer,
			View:        v.View,
			Recommend:   v.Recommend,
			Hot:         v.Hot,
			CreatedDate: v.CreatedDate,
		}
		dto_boards = append(dto_boards, dto_board)
	}
	return dto_boards
}
