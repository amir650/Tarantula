package engine

type MoveStrategy interface {
	Execute(board *Board) Move
	GetNumBoardsEvaluated() int64
}
