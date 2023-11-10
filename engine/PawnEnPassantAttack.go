package engine

type PawnEnPassantAttack struct {
	PawnAttackMove
}

func NewPawnEnPassantAttack(board *Board, move PawnAttackMove) PawnEnPassantAttack {
	return PawnEnPassantAttack{move}
}
