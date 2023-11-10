package engine

type MajorMove struct {
	BaseMove
}

func NewMajorMove(b *Board, p Piece, to int) Move {
	return MajorMove{NewBaseMove(b, p, to)}
}

func (m MajorMove) Execute() *Board {
	builder := NewBoardBuilder()
	for _, piecePtr := range m.GetBoard().GetCurrentPlayer().GetActivePieces() {
		piece := *piecePtr
		mp := *m.GetMovedPiece()
		if !mp.Equals(piece) {
			builder.SetPiece(&piece)
		}
	}
	for _, piecePtr := range m.GetBoard().GetCurrentPlayer().GetOpponent().GetActivePieces() {
		builder.SetPiece(piecePtr)
	}
	mp := *m.GetMovedPiece()
	var finishedMovingPiece = mp.MovePiece(m)
	builder.SetPiece(finishedMovingPiece)
	builder.SetMoveMaker(m.GetBoard().GetCurrentPlayer().GetOpponent().GetAlliance())
	return builder.Build()
}

func (m MajorMove) String() string {
	mp := *m.GetMovedPiece()
	return mp.String() + AlgebraicNotation[m.toCoordinate]
}