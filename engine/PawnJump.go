package engine

type PawnJump struct {
	PawnMove
}

func NewPawnJump(board *Board, p Piece, to int) PawnJump {
	return PawnJump{NewPawnMove(board, p, to)}
}

func (m PawnJump) Execute() *Board {
	builder := NewBoardBuilder()
	builder.basicSetup(m)
	finishedMovingPiece := m.GetMovedPiece().MovePiece(m)
	builder.SetPiece(finishedMovingPiece)
	builder.SetMoveMaker(m.GetBoard().GetCurrentPlayer().GetOpponent().GetAlliance())
	builder.setMoveTransition(m)
	return builder.Build()
}

func (m PawnJump) String() string {
	return AlgebraicNotation[m.toCoordinate]
}
