package engine

import (
	"fmt"
	"sync"
)

type King struct {
	alliance               Alliance
	position               int
	isMoved                bool
	isCastled              bool
	kingSideCastleCapable  bool
	queenSideCastleCapable bool
}

func NewKing(alliance Alliance,
	position int,
	isMoved bool,
	isCastled bool,
	kingSideCastleCapable bool,
	queenSideCastleCapable bool) *Piece {
	king := GetOrCreateKing(alliance, position, isMoved, isCastled, kingSideCastleCapable, queenSideCastleCapable)
	return &king
}

var kingCache = struct {
	sync.Mutex
	pieces map[string]Piece
}{
	pieces: make(map[string]Piece),
}

func GetKingFromCache(key string) (Piece, bool) {
	kingCache.Lock()
	defer kingCache.Unlock()

	piece, found := kingCache.pieces[key]
	return piece, found
}

func AddKingToCache(key string, piece Piece) {
	kingCache.Lock()
	defer kingCache.Unlock()
	kingCache.pieces[key] = piece
}

func GetOrCreateKing(alliance Alliance,
	position int,
	isMoved bool,
	isCastled bool,
	kingSideCastleCapable bool,
	queenSideCastleCapable bool) Piece {
	key := fmt.Sprintf("%d-%d", alliance, position)
	// Check if the piece is already in the cache
	cachedPiece, found := GetKingFromCache(key)
	if !found {
		// If not found, create a new piece
		var newKing = King{alliance, position, isMoved, isCastled, kingSideCastleCapable, queenSideCastleCapable}
		var kingPtr Piece = &newKing
		// Add it to the cache
		AddKingToCache(key, kingPtr)
		// Use the newly created piece
		cachedPiece = kingPtr
	}
	return cachedPiece
}

func (king *King) GetAlliance() Alliance {
	return king.alliance
}

func (king *King) GetPiecePosition() int {
	return king.position
}

func (king *King) String() string {
	return "K"
}

func (king *King) CalculateLegalMoves(board *Board) []Move {

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
	var candidateDestinationCoordinate = 0
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

func (king *King) MovePiece(m Move) *Piece {
	return NewKing(king.GetAlliance(), m.GetTo(), true, king.isCastled, king.kingSideCastleCapable, king.queenSideCastleCapable)
}

func (king *King) Equals(other Piece) bool {
	op := other
	if k, ok := op.(*King); ok {
		return king.GetPiecePosition() == k.GetPiecePosition() && king.GetAlliance() == k.GetAlliance()
	} else {
		return false
	}
}

func (king *King) GetPieceValue() int {
	return 10000
}

func (king *King) GetLocationBonus() int {
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
