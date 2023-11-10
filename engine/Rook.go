package engine

import (
	"fmt"
	"sync"
)

type Rook struct {
	BasePiece
}

func NewRook(alliance Alliance, position int, isMoved bool) *Piece {
	rook := GetOrCreateRook(alliance, position, isMoved)
	return &rook
}

var rookCache = struct {
	sync.Mutex
	pieces map[string]Piece
}{
	pieces: make(map[string]Piece),
}

func GetRookFromCache(key string) (Piece, bool) {
	rookCache.Lock()
	defer rookCache.Unlock()

	piece, found := rookCache.pieces[key]
	return piece, found
}

func AddRookToCache(key string, piece Piece) {
	rookCache.Lock()
	defer rookCache.Unlock()
	rookCache.pieces[key] = piece
}

func GetOrCreateRook(alliance Alliance, position int, isMoved bool) Piece {
	key := fmt.Sprintf("%d-%d-%d", alliance, position, isMoved)
	cachedPiece, found := GetRookFromCache(key)
	if !found {
		var newRook = Rook{BasePiece{alliance, position, isMoved}}
		var rookPtr Piece = &newRook
		AddRookToCache(key, rookPtr)
		cachedPiece = rookPtr
	}
	return cachedPiece
}

func (rook Rook) GetAlliance() Alliance {
	return rook.alliance
}

func (rook Rook) GetPiecePosition() int {
	return rook.position
}

func (rook Rook) String() string {
	return "R"
}

func (rook Rook) CalculateLegalMoves(board *Board) []Move {
	isColumnExclusion := func(currentCandidate int, candidateDestinationCoordinate int) bool {
		return (FirstColumn[candidateDestinationCoordinate] && (currentCandidate == -1)) ||
			(EighthColumn[candidateDestinationCoordinate] && (currentCandidate == 1))
	}

	var legalMoves []Move
	var CandidateMoveCoordinates = []int{-8, -1, 1, 8}
	var candidateDestinationCoordinate = 0

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

func (rook Rook) MovePiece(m Move) *Piece {
	return NewRook(rook.GetAlliance(), m.GetTo(), true)
}

func (rook Rook) Equals(other Piece) bool {
	if r, ok := other.(Rook); ok {
		return rook.GetPiecePosition() == r.GetPiecePosition() && rook.GetAlliance() == r.GetAlliance()
	} else {
		return false
	}
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
