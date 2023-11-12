package engine

type BlackPlayer struct {
	BasePlayer
}

func NewBlackPlayer(board *Board,
	whiteLegals *[]Move,
	blackLegals *[]Move,
	blackKing *King) Player {
	var attacksOnKing = calculateAttacksOnPosition(blackKing.GetPiecePosition(), whiteLegals)
	var blackPlayer = BlackPlayer{BasePlayer{
		board:      board,
		legalMoves: blackLegals,
		playerKing: blackKing,
		inCheck:    len(attacksOnKing) > 0,
	}}
	return blackPlayer
}

func (p BlackPlayer) GetActivePieces() []Piece {
	return p.board.GetBlackPieces()
}

func (p BlackPlayer) GetOpponent() Player {
	return p.board.GetWhitePlayer()
}
