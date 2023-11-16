package engine

type MajorMove struct {
	BaseMove
}

func NewMajorMove(b *Board, p Piece, to int) Move {
	return MajorMove{NewBaseMove(b, p, to)}
}

func (m MajorMove) Execute() *Board {
	builder := NewBoardBuilder()
	builder.basicSetup(m)
	finishedMovingPiece := m.GetMovedPiece().MovePiece(m)
	builder.SetPiece(finishedMovingPiece)
	builder.SetMoveMaker(m.GetBoard().GetCurrentPlayer().GetOpponent().GetAlliance())
	builder.setMoveTransition(m)
	return builder.Build()
}

func (m MajorMove) String() string {
	mp := m.GetMovedPiece()
	return mp.String() + AlgebraicNotation[m.toCoordinate]
}
