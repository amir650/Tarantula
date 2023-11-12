package engine

type PawnMove struct {
	BaseMove
}

func NewPawnMove(b *Board, p Piece, to int) PawnMove {
	return PawnMove{NewBaseMove(b, p, to)}
}

func (m PawnMove) String() string {
	return AlgebraicNotation[m.toCoordinate]
}

func (m PawnMove) Execute() *Board {
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
