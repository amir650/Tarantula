package engine

type PawnAttackMove struct {
	AttackMove
}

func NewPawnAttackMove(b *Board, p Piece, to int, attackedPiece Piece) PawnAttackMove {
	return PawnAttackMove{NewAttackMove(b, p, to, attackedPiece)}
}

func (m PawnAttackMove) String() string {
	mp := *m.GetMovedPiece()
	return AlgebraicNotation[mp.GetPiecePosition()][0:1] + "x" + AlgebraicNotation[m.toCoordinate]
}
