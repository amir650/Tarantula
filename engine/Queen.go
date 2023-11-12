package engine

type Queen struct {
	BasePiece
}

func NewQueen(alliance Alliance, position int, isMoved bool) Queen {
	return Queen{BasePiece{alliance: alliance, position: position, isMoved: isMoved}}
}

func (queen Queen) GetAlliance() Alliance {
	return queen.alliance
}

func (queen Queen) GetPiecePosition() int {
	return queen.position
}

func (queen Queen) String() string {
	return QueenIdentifier
}

func (queen Queen) CalculateLegalMoves(board *Board) []Move {
	isColumnExclusion := func(currentCandidate int, candidateDestinationCoordinate int) bool {
		return (FirstColumn[candidateDestinationCoordinate] && (currentCandidate == -9 || currentCandidate == -1 || currentCandidate == 7)) ||
			(EighthColumn[candidateDestinationCoordinate] && (currentCandidate == -7 || currentCandidate == 1 || currentCandidate == 9))
	}

	var legalMoves []Move
	var CandidateMoveCoordinates = []int{-9, -8, -7, -1, 1, 7, 8, 9}
	var candidateDestinationCoordinate = 0

	for _, currentCandidate := range CandidateMoveCoordinates {
		candidateDestinationCoordinate = queen.position
		for {
			if isColumnExclusion(currentCandidate, candidateDestinationCoordinate) {
				break
			}
			candidateDestinationCoordinate += currentCandidate
			if !IsValidTileCoordinate(candidateDestinationCoordinate) {
				break
			} else {
				var candidateDestinationTile = board.GetTile(candidateDestinationCoordinate)
				if !candidateDestinationTile.IsOccupied() {
					legalMoves = append(legalMoves, NewMajorMove(board, queen, candidateDestinationCoordinate))
				} else {
					var pieceAtDestination = candidateDestinationTile.GetPiece()
					var pieceAtDestinationAllegiance = pieceAtDestination.GetAlliance()
					if queen.alliance != pieceAtDestinationAllegiance {
						legalMoves = append(legalMoves, NewAttackMove(board, queen, candidateDestinationCoordinate, pieceAtDestination))
					}
					break
				}
			}
		}
	}
	return legalMoves
}

func (queen Queen) MovePiece(m Move) Piece {
	return Queen{BasePiece{queen.GetAlliance(), m.GetTo(), true}}
}

func (queen Queen) Equals(other Piece) bool {
	if q, ok := other.(Queen); ok {
		return queen.GetPiecePosition() == q.GetPiecePosition() && queen.GetAlliance() == q.GetAlliance()
	} else {
		return false
	}
}

func (queen Queen) GetPieceValue() int {
	return 900
}

func (queen Queen) GetLocationBonus() int {
	a := queen.GetAlliance()
	switch a {
	case WHITE:
		return WhiteQueenPreferredCoordinates[queen.GetPiecePosition()]
	case BLACK:
		return BlackQueenPreferredCoordinates[queen.GetPiecePosition()]
	default:
		panic("wtf")
	}
}
