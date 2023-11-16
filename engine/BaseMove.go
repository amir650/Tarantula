package engine

type BaseMove struct {
	board         *Board
	movedPiece    Piece
	attackedPiece Piece
	toCoordinate  int
	firstMove     bool
}

func NewBaseMove(b *Board, p Piece, to int) BaseMove {
	return BaseMove{board: b, toCoordinate: to, movedPiece: p}
}

func (m BaseMove) GetBoard() *Board {
	return m.board
}

func (m BaseMove) GetFrom() int {
	return m.GetMovedPiece().GetPiecePosition()
}

func (m BaseMove) GetTo() int {
	return m.toCoordinate
}

func (m BaseMove) IsAttack() bool {
	return false
}

func (m BaseMove) GetMovedPiece() Piece {
	return m.movedPiece
}

func (m BaseMove) GetAttackedPiece() Piece {
	if m.IsAttack() {
		return m.attackedPiece
	}
	return nil
}

func (m BaseMove) Execute() *Board {
	panic("implement me")
}

func (m BaseMove) String() string {
	panic("implement me!")
}
