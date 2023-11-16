package engine

type Tile struct {
	piece  *Piece
	tileID int
}

func (t Tile) IsOccupied() bool {
	return t.piece != nil
}

func (t Tile) GetPiece() Piece {
	return *t.piece
}

func (t Tile) String() string {
	if t.IsOccupied() {
		return t.GetPiece().String()
	}
	return "-"
}
