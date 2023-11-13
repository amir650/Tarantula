package engine

type AttackMove struct {
	BaseMove
	attackedPiece Piece
}

func NewAttackMove(b *Board, p Piece, to int, attackedPiece Piece) AttackMove {
	return AttackMove{NewBaseMove(b, p, to), attackedPiece}
}

func (m AttackMove) Execute() *Board {
	builder := NewBoardBuilder()
	for _, piece := range m.GetBoard().GetCurrentPlayer().GetActivePieces() {
		mp := m.GetMovedPiece()
		if !mp.Equals(piece) {
			builder.SetPiece(piece)
		}
	}
	for _, piecePtr := range m.GetBoard().GetCurrentPlayer().GetOpponent().GetActivePieces() {
		builder.SetPiece(piecePtr)
	}
	mp := m.GetMovedPiece()
	var finishedMovingPiece = mp.MovePiece(m)
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
