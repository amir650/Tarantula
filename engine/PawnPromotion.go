package engine

type PawnPromotion struct {
	decoratedMove   *PawnMove
	promotedPawn    Piece
	promotionTarget Piece
	PawnMove
}

func NewPawnPromotion(move PawnMove,
	target Piece) PawnPromotion {
	return PawnPromotion{decoratedMove: &move, promotedPawn: move.GetMovedPiece(), promotionTarget: target}
}
