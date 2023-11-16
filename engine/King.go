package engine

type King struct {
	BasePiece
	isCastled              bool
	kingSideCastleCapable  bool
	queenSideCastleCapable bool
}

func NewKing(alliance Alliance,
	position int,
	isMoved bool,
	isCastled bool,
	kingSideCastleCapable bool,
	queenSideCastleCapable bool) King {
	return King{
		BasePiece:              BasePiece{alliance: alliance, position: position, isMoved: isMoved},
		isCastled:              isCastled,
		kingSideCastleCapable:  kingSideCastleCapable,
		queenSideCastleCapable: queenSideCastleCapable,
	}
}

func (king King) GetAlliance() Alliance {
	return king.alliance
}

func (king King) GetPiecePosition() int {
	return king.position
}

func (king King) String() string {
	return KingIdentifier
}

func (king King) CalculateLegalMoves(board *Board) []Move {
	isFirstColumnExclusion := func(currentCandidate int, candidateDestinationCoordinate int) bool {
		return FirstColumn[currentCandidate] && ((candidateDestinationCoordinate == -9) || (candidateDestinationCoordinate == -1) ||
			(candidateDestinationCoordinate == 7))
	}

	isEighthColumnExclusion := func(currentCandidate int, candidateDestinationCoordinate int) bool {
		return EighthColumn[currentCandidate] && ((candidateDestinationCoordinate == -7) || (candidateDestinationCoordinate == 1) ||
			(candidateDestinationCoordinate == 9))
	}

	var legalMoves []Move

	var CandidateMoveCoordinates = []int{-9, -8, -7, -1, 1, 7, 8, 9}
	candidateDestinationCoordinate := 0
	for _, currentCandidate := range CandidateMoveCoordinates {
		if isFirstColumnExclusion(king.position, currentCandidate) || isEighthColumnExclusion(king.position, currentCandidate) {
			continue
		}
		candidateDestinationCoordinate = king.position + currentCandidate
		if IsValidTileCoordinate(candidateDestinationCoordinate) {
			var candidateDestinationTile = board.GetTile(candidateDestinationCoordinate)
			if !candidateDestinationTile.IsOccupied() {
				legalMoves = append(legalMoves, NewMajorMove(board, king, candidateDestinationCoordinate))
			} else {
				pieceAtDestination := candidateDestinationTile.GetPiece()
				pieceAtDestinationAllegiance := pieceAtDestination.GetAlliance()
				if king.alliance != pieceAtDestinationAllegiance {
					legalMoves = append(legalMoves, NewAttackMove(board, king, candidateDestinationCoordinate, pieceAtDestination))
				}
			}
		}
	}

	return legalMoves
}

func (king King) MovePiece(m Move) Piece {
	return King{BasePiece{king.GetAlliance(), m.GetTo(), true}, king.isCastled, king.kingSideCastleCapable, king.queenSideCastleCapable}
}

func (king King) Equals(other Piece) bool {
	op := other
	k, ok := op.(King)
	if ok {
		return king.GetPiecePosition() == k.GetPiecePosition() && king.GetAlliance() == k.GetAlliance()
	}
	return false
}

func (king King) GetPieceValue() int {
	return 10000
}

func (king King) GetLocationBonus() int {
	a := king.GetAlliance()
	switch a {
	case WHITE:
		return WhiteKingPreferredCoordinates[king.GetPiecePosition()]
	case BLACK:
		return BlackKingPreferredCoordinates[king.GetPiecePosition()]
	default:
		panic("wtf")
	}
}
