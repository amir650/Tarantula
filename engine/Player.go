package engine

type Player interface {
	GetBoard() *Board
	GetActivePieces() []Piece
	GetLegalMoves() []Move
	GetAlliance() Alliance
	GetOpponent() Player
	IsInCheck() bool
	IsInCheckMate() bool
	IsInStaleMate() bool
	IsCastled() bool
	MakeMove(m Move) *MoveTransition
	String() string
}

func PickPlayer(a Alliance,
	w Player,
	b Player) Player {
	switch a {
	case WHITE:
		return w
	case BLACK:
		return b
	default:
		panic("no!")
	}
}
