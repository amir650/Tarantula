package engine

type NullPiece struct {
}

func (n NullPiece) CalculateLegalMoves(b *Board) []Move {
	panic("NullPiece!")
}

func (n NullPiece) GetAlliance() Alliance {
	panic("NullPiece!")
}

func (n NullPiece) GetPiecePosition() int {
	panic("NullPiece!")
}

func (n NullPiece) String() string {
	panic("NullPiece!")
}

func (n NullPiece) MovePiece(m Move) *Piece {
	panic("NullPiece!")
}

func (n NullPiece) Equals(other Piece) bool {
	panic("NullPiece!")
}

func (n NullPiece) GetPieceValue() int {
	panic("NullPiece!")
}

func (n NullPiece) GetLocationBonus() int {
	panic("NullPiece!")
}
