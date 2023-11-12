package engine

type Bishop struct {
	BasePiece
}

func NewBishop(alliance Alliance, position int, isMoved bool) Bishop {
	return Bishop{BasePiece{alliance: alliance, position: position, isMoved: isMoved}}
}

func (bishop Bishop) GetAlliance() Alliance {
	return bishop.alliance
}

func (bishop Bishop) GetPiecePosition() int {
	return bishop.position
}

func (bishop Bishop) String() string {
	return BishopIdentifier
}

func (bishop Bishop) CalculateLegalMoves(board *Board) []Move {

	isDiagonalExclusion := func(currentCandidate int, candidateDestinationCoordinate int) bool {
		return (FirstColumn[candidateDestinationCoordinate] &&
			((currentCandidate == -9) || (currentCandidate == 7))) ||
			(EighthColumn[candidateDestinationCoordinate] &&
				((currentCandidate == -7) || (currentCandidate == 9)))
	}

	var legalMoves []Move
	var CandidateMoveCoordinates = []int{-9, -7, 7, 9}
	var candidateDestinationCoordinate = 0

	for _, currentCandidate := range CandidateMoveCoordinates {
		candidateDestinationCoordinate = bishop.position
		for IsValidTileCoordinate(candidateDestinationCoordinate) {
			if isDiagonalExclusion(currentCandidate, candidateDestinationCoordinate) {
				break
			}
			candidateDestinationCoordinate += currentCandidate
			if IsValidTileCoordinate(candidateDestinationCoordinate) {
				var candidateDestinationTile = board.GetTile(candidateDestinationCoordinate)
				if !candidateDestinationTile.IsOccupied() {
					legalMoves = append(legalMoves, NewMajorMove(board, bishop, candidateDestinationCoordinate))
				} else {
					var pieceAtDestination = candidateDestinationTile.GetPiece()
					var pieceAtDestinationAllegiance = pieceAtDestination.GetAlliance()
					if bishop.alliance != pieceAtDestinationAllegiance {
						legalMoves = append(legalMoves, NewAttackMove(board, bishop, candidateDestinationCoordinate, pieceAtDestination))
					}
					break
				}
			}
		}
	}

	return legalMoves
}

func (bishop Bishop) MovePiece(m Move) Piece {
	return Bishop{BasePiece{bishop.GetAlliance(), m.GetTo(), true}}
}

func (bishop Bishop) Equals(other Piece) bool {
	if kn, ok := other.(*Bishop); ok {
		return bishop.GetPiecePosition() == kn.GetPiecePosition() && bishop.GetAlliance() == kn.GetAlliance()
	} else {
		return false
	}
}

func (bishop Bishop) GetPieceValue() int {
	return 330
}

func (bishop Bishop) GetLocationBonus() int {
	a := bishop.GetAlliance()
	switch a {
	case WHITE:
		return WhiteBishopPreferredCoordinates[bishop.GetPiecePosition()]
	case BLACK:
		return BlackBishopPreferredCoordinates[bishop.GetPiecePosition()]
	default:
		panic("wtf")
	}
}
