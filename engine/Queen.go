package engine

import (
	"fmt"
	"sync"
)

type Queen struct {
	BasePiece
}

func NewQueen(alliance Alliance, position int, isMoved bool) *Piece {
	queen := GetOrCreateQueen(alliance, position, isMoved)
	return &queen
}

var queenCache = struct {
	sync.Mutex
	pieces map[string]Piece
}{
	pieces: make(map[string]Piece),
}

func GetQueenFromCache(key string) (Piece, bool) {
	queenCache.Lock()
	defer queenCache.Unlock()

	piece, found := queenCache.pieces[key]
	return piece, found
}

func AddQueenToCache(key string, piece Piece) {
	queenCache.Lock()
	defer queenCache.Unlock()
	queenCache.pieces[key] = piece
}

func GetOrCreateQueen(alliance Alliance, position int, isMoved bool) Piece {
	key := fmt.Sprintf("%d-%d-%d", alliance, position, isMoved)
	// Check if the piece is already in the cache
	cachedPiece, found := GetQueenFromCache(key)
	if !found {
		// If not found, create a new piece
		var newQueen = Queen{BasePiece{alliance, position, isMoved}}
		var queenPtr Piece = &newQueen
		// Add it to the cache
		AddQueenToCache(key, queenPtr)
		// Use the newly created piece
		cachedPiece = queenPtr
	}
	return cachedPiece
}

func (queen *Queen) GetAlliance() Alliance {
	return queen.alliance
}

func (queen *Queen) GetPiecePosition() int {
	return queen.position
}

func (queen *Queen) String() string {
	return "Q"
}

func (queen *Queen) CalculateLegalMoves(board *Board) []Move {
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

func (queen *Queen) MovePiece(m Move) *Piece {
	return NewQueen(queen.GetAlliance(), m.GetTo(), true)
}

func (queen *Queen) Equals(other Piece) bool {
	if q, ok := other.(*Queen); ok {
		return queen.GetPiecePosition() == q.GetPiecePosition() && queen.GetAlliance() == q.GetAlliance()
	} else {
		return false
	}
}

func (queen *Queen) GetPieceValue() int {
	return 900
}

func (queen *Queen) GetLocationBonus() int {
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
