package engine

type Tile struct {
	tileID int
	piece  *Piece
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
	} else {
		return "-"
	}
}
