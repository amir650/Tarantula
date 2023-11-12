package engine

type WhitePlayer struct {
	BasePlayer
}

func NewWhitePlayer(board *Board,
	whiteLegals *[]Move,
	blackLegals *[]Move,
	whiteKing *King) Player {
	var attacksOnKing = calculateAttacksOnPosition(whiteKing.GetPiecePosition(), blackLegals)
	var whitePlayer = WhitePlayer{BasePlayer{
		board:      board,
		legalMoves: whiteLegals,
		playerKing: whiteKing,
		inCheck:    len(attacksOnKing) > 0,
	}}
	return whitePlayer
}

func (w WhitePlayer) GetActivePieces() []Piece {
	return w.board.GetWhitePieces()
}

func (w WhitePlayer) GetOpponent() Player {
	return w.GetBoard().GetBlackPlayer()
}
