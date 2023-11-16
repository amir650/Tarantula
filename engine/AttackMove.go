package engine

type AttackMove struct {
	attackedPiece Piece
	BaseMove
}

func NewAttackMove(b *Board, p Piece, to int, attackedPiece Piece) AttackMove {
	return AttackMove{BaseMove: NewBaseMove(b, p, to), attackedPiece: attackedPiece}
}

func (m AttackMove) Execute() *Board {
	builder := NewBoardBuilder()
	builder.basicSetup(m)
	finishedMovingPiece := m.GetMovedPiece().MovePiece(m)
	builder.SetPiece(finishedMovingPiece)
	builder.SetMoveMaker(m.GetBoard().GetCurrentPlayer().GetOpponent().GetAlliance())
	builder.setMoveTransition(m)
	return builder.Build()
}

func (m AttackMove) String() string {
	mp := m.GetMovedPiece()
	return mp.String() + "x" + AlgebraicNotation[m.toCoordinate]
}

func (m AttackMove) isAttack() bool {
	return true
}
