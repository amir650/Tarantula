package engine

type Rook struct {
	BasePiece
}

func NewRook(alliance Alliance, position int, isMoved bool) Rook {
	return Rook{BasePiece{alliance: alliance, position: position, isMoved: isMoved}}
}

func (rook Rook) GetAlliance() Alliance {
	return rook.alliance
}

func (rook Rook) GetPiecePosition() int {
	return rook.position
}

func (rook Rook) String() string {
	return RookIdentifier
}

func (rook Rook) CalculateLegalMoves(board *Board) []Move {
	isColumnExclusion := func(currentCandidate int, candidateDestinationCoordinate int) bool {
		return (FirstColumn[candidateDestinationCoordinate] && (currentCandidate == -1)) ||
			(EighthColumn[candidateDestinationCoordinate] && (currentCandidate == 1))
	}

	var legalMoves []Move
	var CandidateMoveCoordinates = []int{-8, -1, 1, 8}
	candidateDestinationCoordinate := 0

	for _, currentCandidate := range CandidateMoveCoordinates {
		candidateDestinationCoordinate = rook.position
		for IsValidTileCoordinate(candidateDestinationCoordinate) {
			if isColumnExclusion(currentCandidate, candidateDestinationCoordinate) {
				break
			}
			candidateDestinationCoordinate += currentCandidate
			if IsValidTileCoordinate(candidateDestinationCoordinate) {
				var candidateDestinationTile = board.GetTile(candidateDestinationCoordinate)
				if !candidateDestinationTile.IsOccupied() {
					legalMoves = append(legalMoves, NewMajorMove(board, rook, candidateDestinationCoordinate))
				} else {
					var pieceAtDestination = candidateDestinationTile.GetPiece()
					var pieceAtDestinationAllegiance = pieceAtDestination.GetAlliance()
					if rook.alliance != pieceAtDestinationAllegiance {
						legalMoves = append(legalMoves, NewAttackMove(board, rook, candidateDestinationCoordinate, pieceAtDestination))
					}
					break
				}
			}
		}
	}

	return legalMoves
}

func (rook Rook) MovePiece(m Move) Piece {
	return Rook{BasePiece{rook.GetAlliance(), m.GetTo(), true}}
}

func (rook Rook) Equals(other Piece) bool {
	r, ok := other.(Rook)
	if ok {
		return rook.GetPiecePosition() == r.GetPiecePosition() && rook.GetAlliance() == r.GetAlliance()
	}
	return false
}

func (rook Rook) GetPieceValue() int {
	return 500
}

func (rook Rook) GetLocationBonus() int {
	a := rook.GetAlliance()
	switch a {
	case WHITE:
		return WhiteRookPreferredCoordinates[rook.GetPiecePosition()]
	case BLACK:
		return BlackRookPreferredCoordinates[rook.GetPiecePosition()]
	default:
		panic("wtf")
	}
}
