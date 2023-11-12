package engine

type PawnJump struct {
	PawnMove
}

func NewPawnJump(board *Board, p Piece, to int) PawnJump {
	return PawnJump{NewPawnMove(board, p, to)}
}

func (m PawnJump) Execute() *Board {
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
	return builder.Build()
}

func (m PawnJump) String() string {
	return AlgebraicNotation[m.toCoordinate]
}
