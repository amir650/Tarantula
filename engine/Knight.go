package engine

type Knight struct {
	BasePiece
}

func NewKnight(alliance Alliance, position int, isMoved bool) Knight {
	return Knight{BasePiece{alliance: alliance, position: position, isMoved: isMoved}}
}

func (knight Knight) GetAlliance() Alliance {
	return knight.alliance
}

func (knight Knight) GetPiecePosition() int {
	return knight.position
}

func (knight Knight) String() string {
	return KnightIdentifier
}

func (knight Knight) CalculateLegalMoves(board *Board) []Move {

	isFirstColumnExclusion := func(piecePosition int, currentCandidate int) bool {
		return FirstColumn[piecePosition] && ((currentCandidate == -17) ||
			(currentCandidate == -10) || (currentCandidate == 6) || (currentCandidate == 15))
	}

	isSecondColumnExclusion := func(piecePosition int, currentCandidate int) bool {
		return SecondColumn[piecePosition] && ((currentCandidate == -10) || (currentCandidate == 6))
	}

	isSeventhColumnExclusion := func(piecePosition int, currentCandidate int) bool {
		return SeventhColumn[piecePosition] && ((currentCandidate == -6) || (currentCandidate == 10))
	}

	isEighthColumnExclusion := func(piecePosition int, currentCandidate int) bool {
		return EighthColumn[piecePosition] && ((currentCandidate == -15) || (currentCandidate == -6) ||
			(currentCandidate == 10) || (currentCandidate == 17))
	}

	var legalMoves []Move
	var CandidateMoveCoordinates = []int{-17, -15, -10, -6, 6, 10, 15, 17}

	for _, currentCandidate := range CandidateMoveCoordinates {
		if isFirstColumnExclusion(knight.position, currentCandidate) ||
			isSecondColumnExclusion(knight.position, currentCandidate) ||
			isSeventhColumnExclusion(knight.position, currentCandidate) ||
			isEighthColumnExclusion(knight.position, currentCandidate) {
			continue
		}
		candidateDestinationCoordinate := knight.position + currentCandidate
		if IsValidTileCoordinate(candidateDestinationCoordinate) {
			candidateDestinationTile := board.GetTile(candidateDestinationCoordinate)
			if !candidateDestinationTile.IsOccupied() {
				mm := NewMajorMove(board, knight, candidateDestinationCoordinate)
				legalMoves = append(legalMoves, mm)
			} else {
				pieceAtDestination := candidateDestinationTile.GetPiece()
				pieceAtDestinationAllegiance := pieceAtDestination.GetAlliance()
				if knight.alliance != pieceAtDestinationAllegiance {
					legalMoves = append(legalMoves, NewAttackMove(board, knight, candidateDestinationCoordinate, pieceAtDestination))
				}
			}
		}
	}
	return legalMoves
}

func (knight Knight) MovePiece(m Move) Piece {
	return Knight{BasePiece{knight.GetAlliance(), m.GetTo(), true}}
}

func (knight Knight) Equals(other Piece) bool {
	if kn, ok := other.(Knight); ok {
		return knight.GetPiecePosition() == kn.GetPiecePosition() && knight.GetAlliance() == kn.GetAlliance()
	} else {
		return false
	}
}

func (knight Knight) GetPieceValue() int {
	return 300
}

func (knight Knight) GetLocationBonus() int {
	a := knight.GetAlliance()
	switch a {
	case WHITE:
		return WhiteKnightPreferredCoordinates[knight.GetPiecePosition()]
	case BLACK:
		return BlackKnightPreferredCoordinates[knight.GetPiecePosition()]
	default:
		panic("wtf")
	}
}
