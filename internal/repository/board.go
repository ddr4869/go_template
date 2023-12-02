package repository

import (
	"context"

	"github.com/go-board/ent"
	"github.com/go-board/ent/board"
	"github.com/go-board/ent/user"
	"github.com/the-medium-tech/platform-externals/log"
)

type Board struct {
	client *ent.Client
}

func NewBoard(client *ent.Client) *Board {
	return &Board{client: client}
}

func (r *Board) BoardList(ctx context.Context) ([]*ent.Board, error) {

	query := r.client.Board.Query()
	list, err := query.All(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return list, nil
}

func (r *Board) UserBoardList(ctx context.Context, writer string) ([]*ent.Board, error) {

	query := r.client.User.Query().Where(user.Name(writer)).QueryBoard()
	list, err := query.All(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return list, nil
}

func (r *Board) CreateBoard(ctx context.Context, title, content, writer string) (*ent.Board, error) {

	user, err := r.client.User.Query().Where(user.Name(writer)).Only(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	Board, err := r.client.Board.Create().
		SetTitle(title).
		SetContent(content).
		SetWriter(writer).
		SetUser(user).
		Save(ctx)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return Board, nil
}

func (r *Board) ReadBoard(ctx context.Context, boardId int) (*ent.Board, error) {
	board, err := r.client.Board.UpdateOneID(uint(boardId)).AddView(1).Save(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return board, nil
}

func (r *Board) RecommendBoard(ctx context.Context, boardId int) (int, error) {

	board, err := r.client.Board.UpdateOneID(uint(boardId)).AddRecommend(1).
		Save(ctx)

	if err != nil {
		log.Error(err)
		return -1, err
	}

	return int(board.Recommend), nil
}

func (r *Board) TurnToHotBoard(ctx context.Context, boardId int) error {

	_, err := r.client.Board.UpdateOneID(uint(boardId)).SetHot(true).Save(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *Board) HotBoardList(ctx context.Context) ([]*ent.Board, error) {

	boards, err := r.client.Board.Query().Where(board.Hot(true)).All(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return boards, nil
}
