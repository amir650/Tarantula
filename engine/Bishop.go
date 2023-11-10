package engine

import (
	"fmt"
	"sync"
)

type Bishop struct {
	BasePiece
}

func NewBishop(alliance Alliance, position int, isMoved bool) *Piece {
	bishop := GetOrCreateBishop(alliance, position, isMoved)
	return &bishop
}

var bishopCache = struct {
	sync.Mutex
	pieces map[string]Piece
}{
	pieces: make(map[string]Piece),
}

func GetBishopFromCache(key string) (Piece, bool) {
	bishopCache.Lock()
	defer bishopCache.Unlock()

	piece, found := bishopCache.pieces[key]
	return piece, found
}

func AddPBishopToCache(key string, piece Piece) {
	bishopCache.Lock()
	defer bishopCache.Unlock()
	bishopCache.pieces[key] = piece
}

func GetOrCreateBishop(alliance Alliance, position int, isMoved bool) Piece {
	key := fmt.Sprintf("%d-%d-%d", alliance, position, isMoved)
	// Check if the piece is already in the cache
	cachedPiece, found := GetBishopFromCache(key)
	if !found {
		// If not found, create a new piece
		var newBishop = Bishop{BasePiece{alliance, position, isMoved}}
		var bishopPtr Piece = &newBishop
		// Add it to the cache
		AddPBishopToCache(key, bishopPtr)
		// Use the newly created piece
		cachedPiece = bishopPtr
	}
	return cachedPiece
}

func (bishop *Bishop) GetAlliance() Alliance {
	return bishop.alliance
}

func (bishop *Bishop) GetPiecePosition() int {
	return bishop.position
}

func (bishop *Bishop) String() string {
	return "B"
}

func (bishop *Bishop) CalculateLegalMoves(board *Board) []Move {

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

func (bishop *Bishop) MovePiece(m Move) *Piece {
	m.GetMovedPiece()
	return NewBishop(bishop.GetAlliance(), m.GetTo(), true)
}

func (bishop *Bishop) Equals(other Piece) bool {
	if kn, ok := other.(*Bishop); ok {
		return bishop.GetPiecePosition() == kn.GetPiecePosition() && bishop.GetAlliance() == kn.GetAlliance()
	} else {
		return false
	}
}

func (bishop *Bishop) GetPieceValue() int {
	return 330
}

func (bishop *Bishop) GetLocationBonus() int {
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
