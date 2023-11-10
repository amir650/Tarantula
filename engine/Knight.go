package engine

import (
	"fmt"
	"sync"
)

type Knight struct {
	BasePiece
}

var knightCache = struct {
	sync.Mutex
	pieces map[string]Piece
}{
	pieces: make(map[string]Piece),
}

func NewKnight(alliance Alliance,
	position int,
	isMoved bool) *Piece {
	knight := GetOrCreateKnight(alliance, position, isMoved)
	return &knight
}

func GetKnightFromCache(key string) (Piece, bool) {
	knightCache.Lock()
	defer knightCache.Unlock()

	piece, found := knightCache.pieces[key]
	return piece, found
}

func AddKnightToCache(key string, piece Piece) {
	knightCache.Lock()
	defer knightCache.Unlock()
	knightCache.pieces[key] = piece
}

func GetOrCreateKnight(alliance Alliance,
	position int,
	isMoved bool) Piece {
	key := fmt.Sprintf("%d-%d", alliance, position)
	// Check if the piece is already in the cache
	cachedPiece, found := GetKnightFromCache(key)
	if !found {
		// If not found, create a new piece
		var newKnight = Knight{BasePiece{alliance, position, isMoved}}
		var knightPtr Piece = &newKnight
		// Add it to the cache
		AddKnightToCache(key, knightPtr)
		// Use the newly created piece
		cachedPiece = knightPtr
	}
	return cachedPiece
}

func (knight *Knight) GetAlliance() Alliance {
	return knight.alliance
}

func (knight *Knight) GetPiecePosition() int {
	return knight.position
}

func (knight *Knight) String() string {
	return "N"
}

func (knight *Knight) CalculateLegalMoves(board *Board) []Move {

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

func (knight *Knight) MovePiece(m Move) *Piece {
	return NewKnight(knight.GetAlliance(), m.GetTo(), true)
}

func (knight *Knight) Equals(other Piece) bool {
	if kn, ok := other.(*Knight); ok {
		return knight.GetPiecePosition() == kn.GetPiecePosition() && knight.GetAlliance() == kn.GetAlliance()
	} else {
		return false
	}
}

func (knight *Knight) GetPieceValue() int {
	return 300
}

func (knight *Knight) GetLocationBonus() int {
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
