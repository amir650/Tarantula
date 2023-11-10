package engine

import (
	"fmt"
	"strings"
	"sync"
)

type BoardBuilder struct {
	piecesMap      map[int]Piece
	nextMoveMaker  *Alliance
	enPassantPawn  *Pawn
	moveTransition *MoveTransition
}

func NewBoardBuilder() *BoardBuilder {
	return &BoardBuilder{piecesMap: make(map[int]Piece)}
}

func (builder *BoardBuilder) SetMoveMaker(alliance Alliance) *BoardBuilder {
	builder.nextMoveMaker = &alliance
	return builder
}

func (builder *BoardBuilder) SetPiece(piecePtr *Piece) *BoardBuilder {
	piece := *piecePtr
	builder.piecesMap[piece.GetPiecePosition()] = piece
	return builder
}

func (builder *BoardBuilder) setEnPassantPawn(pawn Pawn) *BoardBuilder {
	builder.enPassantPawn = &pawn
	return builder
}

func (builder *BoardBuilder) setMoveTransition(moveTransition MoveTransition) *BoardBuilder {
	builder.moveTransition = &moveTransition
	return builder
}

func (builder *BoardBuilder) Build() *Board {
	return NewBoard(builder)
}

type Board struct {
	gameBoard     *[]Tile
	whitePieces   []*Piece
	blackPieces   []*Piece
	whiteLegals   *[]Move
	blackLegals   *[]Move
	whitePlayer   *Player
	blackPlayer   *Player
	currentPlayer *Player
	enPassantPawn *Pawn
}

func (b Board) GetCurrentPlayer() Player {
	return *b.currentPlayer
}

func (b Board) GetWhitePlayer() Player {
	return *b.whitePlayer
}

func (b Board) GetBlackPlayer() Player {
	return *b.blackPlayer
}

func (b Board) getPiecesByAlliance(alliance Alliance) []Piece {
	var result []Piece
	var t Tile
	fmt.Println(t)
	for _, tile := range b.GetAllTiles() {
		if tile.IsOccupied() {
			if tile.GetPiece().GetAlliance() == alliance {
				result = append(result, tile.GetPiece())
			}
		}
	}
	return result
}

func (b Board) GetPiece(coordinate int) Piece {
	return b.GetTile(coordinate).GetPiece()
}

func (b Board) GetTile(coordinate int) Tile {
	return (*b.gameBoard)[coordinate]
}

func (b Board) GetAllPieces() []*Piece {
	result := make([]*Piece, len(b.GetWhitePieces())+len(b.GetBlackPieces()))
	copy(result[:len(b.GetWhitePieces())], b.GetWhitePieces())
	copy(result[len(b.GetBlackPieces()):], b.GetBlackPieces())
	return result
}

func (b Board) GetEnPassantPawn() *Pawn {
	return b.enPassantPawn
}

func (b Board) GetAllTiles() []Tile {
	return *b.gameBoard
}

func (b Board) GetAllLegalMoves() []Move {
	allMoves := append(*b.whiteLegals, *b.blackLegals...)
	return allMoves
}

func NewBoard(builder *BoardBuilder) *Board {
	tiles := make([]Tile, 64)
	var whitePieces []*Piece
	var blackPieces []*Piece
	var whiteKing *King
	var blackKing *King
	for i := 0; i < 64; i++ {
		if builder.piecesMap[i] != nil {
			piece := builder.piecesMap[i]
			tiles[i] = Tile{
				tileID: i,
				piece:  &piece,
			}
			if piece.GetAlliance() == WHITE {
				whitePieces = append(whitePieces, &piece)
				switch (piece).(type) {
				case *King:
					if k, ok := piece.(*King); ok {
						whiteKing = k
					}
				}
			} else if piece.GetAlliance() == BLACK {
				blackPieces = append(blackPieces, &piece)
				switch piece.(type) {
				case *King:
					if k, ok := piece.(*King); ok {
						blackKing = k
					}
				}
			}
		} else {
			tiles[i] = Tile{
				tileID: i,
			}
		}
	}

	var whiteLegals []Move
	var blackLegals []Move
	var whitePlayer Player
	var blackPlayer Player
	var currentPlayer Player

	protoBoard := Board{
		gameBoard:     &tiles,
		whitePieces:   whitePieces,
		blackPieces:   blackPieces,
		whiteLegals:   &whiteLegals,
		blackLegals:   &blackLegals,
		whitePlayer:   &whitePlayer,
		blackPlayer:   &blackPlayer,
		currentPlayer: &currentPlayer,
	}

	whiteLegals = calculatePieceLegals(&protoBoard, whitePieces)
	blackLegals = calculatePieceLegals(&protoBoard, blackPieces)
	whitePlayer = NewWhitePlayer(&protoBoard, &whiteLegals, &blackLegals, whiteKing)
	blackPlayer = NewBlackPlayer(&protoBoard, &whiteLegals, &blackLegals, blackKing)
	currentPlayer = PickPlayer(*builder.nextMoveMaker, whitePlayer, blackPlayer)

	return &protoBoard
}

func calculatePieceLegals(b *Board, pieces []*Piece) []Move {
	var legalMoves []Move
	for _, piece := range pieces {
		p := *piece
		var pieceLegals = p.CalculateLegalMoves(b)
		legalMoves = append(legalMoves, pieceLegals...)
		//fmt.Printf("%s: %v\n", piecePtr, pieceLegals)
	}
	return legalMoves
}

func (b Board) String() string {
	prettyPrintTile := func(t Tile) string {
		var tileText = t.String()
		if tileText != "-" {
			if t.GetPiece().GetAlliance() == WHITE {
				tileText = strings.ToUpper(tileText)[:1]
			} else {
				tileText = strings.ToLower(tileText)[:1]
			}
		}
		return tileText
	}
	var boardStrings []string
	for i, tile := range b.GetAllTiles() {
		var tileText = prettyPrintTile(tile)
		boardStrings = append(boardStrings, tileText)
		if (i+1)%8 == 0 {
			boardStrings = append(boardStrings, "\n")
		}
	}
	return " " + strings.Join(boardStrings, " ")
}

func (b Board) CalculateLegalMoves() *[]Move {
	var legalMoves []Move
	for _, tile := range b.GetAllTiles() {
		if tile.IsOccupied() {
			var pieceLegals = tile.GetPiece().CalculateLegalMoves(&b)
			legalMoves = append(legalMoves, pieceLegals...)
			//fmt.Printf("%s: %v\n", tile.piece, pieceLegals)
		}
	}
	return &legalMoves
}

func (b Board) CalculateLegalMovesParallelGPT() []Move {
	var legalMoves []Move
	resultChan := make(chan []Move, len(b.GetAllTiles())) // Buffered channel to prevent blocking

	// Counter to track the number of goroutines
	var wg sync.WaitGroup

	for _, tile := range b.GetAllTiles() {
		if tile.IsOccupied() {
			wg.Add(1)         // Increment counter
			go func(t Tile) { // Start a goroutine
				defer wg.Done() // Decrement counter when done
				resultChan <- t.GetPiece().CalculateLegalMoves(&b)
			}(tile)
		}
	}

	// Start a goroutine to close the result channel after all calculations are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect the results from the channel
	for mvs := range resultChan {
		legalMoves = append(legalMoves, mvs...)
	}

	return legalMoves
}

func (b Board) CalculateLegalMovesParallel() []Move {
	var legalMoves []Move
	resultChan := make(chan []Move) // Channel to collect results

	// Counter to track the number of goroutines
	var wg sync.WaitGroup

	for _, tile := range b.GetAllTiles() {
		if tile.IsOccupied() {
			wg.Add(1)         // Increment counter
			go func(t Tile) { // Start a goroutine
				defer wg.Done() // Decrement counter when done
				resultChan <- t.GetPiece().CalculateLegalMoves(&b)
			}(tile)
		}
	}

	// Start a goroutine to close the result channel after all calculations are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect the results from the channel
	for mvs := range resultChan {
		legalMoves = append(legalMoves, mvs...)
	}

	return legalMoves
}

func (b Board) GetBlackPieces() []*Piece {
	return b.blackPieces
}

func (b Board) GetWhitePieces() []*Piece {
	return b.whitePieces
}

func concatenate(slice1, slice2 []Piece) []Piece {
	result := make([]Piece, len(slice1)+len(slice2))
	copy(result[:len(slice1)], slice1)
	copy(result[len(slice1):], slice2)
	return result
}

var AlgebraicNotation = []string{
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
}

var PositionToCoordinateMap = initializePositionToCoordinateMap()

func initializePositionToCoordinateMap() map[string]int {
	positionToCoordinateMap := make(map[string]int)
	for index, value := range AlgebraicNotation {
		positionToCoordinateMap[value] = index
	}
	return positionToCoordinateMap
}

func CreateStandardChessBoard() *Board {

	boardBuilder := NewBoardBuilder()

	boardBuilder.SetPiece(NewRook(BLACK, 0, false))
	boardBuilder.SetPiece(NewKnight(BLACK, 1, false))
	boardBuilder.SetPiece(NewBishop(BLACK, 2, false))
	boardBuilder.SetPiece(NewQueen(BLACK, 3, false))
	boardBuilder.SetPiece(NewKing(BLACK, 4, false, false, true, true))
	boardBuilder.SetPiece(NewBishop(BLACK, 5, false))
	boardBuilder.SetPiece(NewKnight(BLACK, 6, false))
	boardBuilder.SetPiece(NewRook(BLACK, 7, false))
	boardBuilder.SetPiece(NewPawn(BLACK, 8, false))
	boardBuilder.SetPiece(NewPawn(BLACK, 9, false))
	boardBuilder.SetPiece(NewPawn(BLACK, 10, false))
	boardBuilder.SetPiece(NewPawn(BLACK, 11, false))
	boardBuilder.SetPiece(NewPawn(BLACK, 12, false))
	boardBuilder.SetPiece(NewPawn(BLACK, 13, false))
	boardBuilder.SetPiece(NewPawn(BLACK, 14, false))
	boardBuilder.SetPiece(NewPawn(BLACK, 15, false))

	boardBuilder.SetPiece(NewPawn(WHITE, 48, false))
	boardBuilder.SetPiece(NewPawn(WHITE, 49, false))
	boardBuilder.SetPiece(NewPawn(WHITE, 50, false))
	boardBuilder.SetPiece(NewPawn(WHITE, 51, false))
	boardBuilder.SetPiece(NewPawn(WHITE, 52, false))
	boardBuilder.SetPiece(NewPawn(WHITE, 53, false))
	boardBuilder.SetPiece(NewPawn(WHITE, 54, false))
	boardBuilder.SetPiece(NewPawn(WHITE, 55, false))
	boardBuilder.SetPiece(NewRook(WHITE, 56, false))
	boardBuilder.SetPiece(NewKnight(WHITE, 57, false))
	boardBuilder.SetPiece(NewBishop(WHITE, 58, false))
	boardBuilder.SetPiece(NewQueen(WHITE, 59, false))
	boardBuilder.SetPiece(NewKing(WHITE, 60, false, false, true, true))
	boardBuilder.SetPiece(NewBishop(WHITE, 61, false))
	boardBuilder.SetPiece(NewKnight(WHITE, 62, false))
	boardBuilder.SetPiece(NewRook(WHITE, 63, false))
	boardBuilder.SetMoveMaker(WHITE)

	standardBoard := boardBuilder.Build()

	return standardBoard
}
