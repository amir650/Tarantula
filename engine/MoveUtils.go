package engine

import "fmt"

func GetMove(board *Board, current string, destination string) Move {
	c := PositionToCoordinateMap[current]
	d := PositionToCoordinateMap[destination]

	if c > NumTiles || c < 0 {
		fmt.Println("returning null")
		return NullMove{}
	}

	return getMoveImpl(board, c, d)
}

func getMoveImpl(board *Board, currentPosition int, destination int) Move {
	for _, move := range board.GetAllLegalMoves() {
		if move.GetFrom() == currentPosition && move.GetTo() == destination {
			return move
		}
	}
	return NullMove{}
}
