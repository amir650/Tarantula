package engine

type Piece interface {
	CalculateLegalMoves(b *Board) []Move
	GetAlliance() Alliance
	GetPiecePosition() int
	String() string
	MovePiece(m Move) Piece
	Equals(other Piece) bool
	GetPieceValue() int
	GetLocationBonus() int
}
