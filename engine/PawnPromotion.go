package engine

type PawnPromotion struct {
	PawnMove
	decoratedMove   *PawnMove
	promotedPawn    Piece
	promotionTarget Piece
}

func NewPawnPromotion(move PawnMove,
	target Piece) PawnPromotion {
	return PawnPromotion{decoratedMove: &move, promotedPawn: move.GetMovedPiece(), promotionTarget: target}
}
