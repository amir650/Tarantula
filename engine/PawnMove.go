package engine

type PawnMove struct {
	BaseMove
}

func NewPawnMove(b *Board, p Piece, to int) PawnMove {
	return PawnMove{NewBaseMove(b, p, to)}
}

func (m PawnMove) String() string {
	return AlgebraicNotation[m.toCoordinate]
}

func (m PawnMove) Execute() *Board {
	builder := NewBoardBuilder()
	builder.basicSetup(m)
	finishedMovingPiece := m.GetMovedPiece().MovePiece(m)
	builder.SetPiece(finishedMovingPiece)
	builder.SetMoveMaker(m.GetBoard().GetCurrentPlayer().GetOpponent().GetAlliance())
	builder.setMoveTransition(m)
	return builder.Build()
}
