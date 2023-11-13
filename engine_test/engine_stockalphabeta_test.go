package engine

import (
	"Tarantula/engine"
	"fmt"
	"testing"
)

func TestStockAlphaBeta(t *testing.T) {
	board := engine.CreateStandardChessBoard()
	strategy := engine.NewStockAlphaBeta(6, true)

	fmt.Println(board.String())
	m := strategy.Execute(board)
	mt := board.GetCurrentPlayer().MakeMove(m)
	board = mt.GetToBoard()

}
