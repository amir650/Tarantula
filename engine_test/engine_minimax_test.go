package engine

import (
	"Tarantula/engine"
	"fmt"
	"testing"
)

func TestDepth1InitialConfiguration(t *testing.T) {
	board := engine.CreateStandardChessBoard()
	strategy := engine.NewMinimax(1, false)

	fmt.Println(board.String())
	m := strategy.Execute(board)
	mt := board.GetCurrentPlayer().MakeMove(m)
	board = mt.GetToBoard()

	if strategy.GetNumBoardsEvaluated() != 20 {
		t.Errorf("should have gotten different number of legals")
	}
}

func TestDepth2InitialConfiguration(t *testing.T) {
	board := engine.CreateStandardChessBoard()
	strategy := engine.NewMinimax(2, false)

	fmt.Println(board.String())
	m := strategy.Execute(board)
	mt := board.GetCurrentPlayer().MakeMove(m)
	board = mt.GetToBoard()

	if strategy.GetNumBoardsEvaluated() != 400 {
		t.Errorf("should have gotten different number of legals")
	}
}

func TestDepth3InitialConfiguration(t *testing.T) {
	board := engine.CreateStandardChessBoard()
	strategy := engine.NewMinimax(3, false)

	fmt.Println(board.String())
	m := strategy.Execute(board)
	mt := board.GetCurrentPlayer().MakeMove(m)
	board = mt.GetToBoard()

	if strategy.GetNumBoardsEvaluated() != 8902 {
		t.Errorf("should have gotten different number of legals")
	}
}

func TestDepth4InitialConfiguration(t *testing.T) {
	board := engine.CreateStandardChessBoard()
	strategy := engine.NewMinimax(4, true)

	fmt.Println(board.String())
	m := strategy.Execute(board)
	mt := board.GetCurrentPlayer().MakeMove(m)
	board = mt.GetToBoard()

	if strategy.GetNumBoardsEvaluated() != 197281 {
		t.Errorf("should have gotten different number of legals")
	}
}

func TestDepth5InitialConfiguration(t *testing.T) {
	board := engine.CreateStandardChessBoard()
	strategy := engine.NewMinimax(5, true)

	fmt.Println(board.String())
	m := strategy.Execute(board)
	mt := board.GetCurrentPlayer().MakeMove(m)
	board = mt.GetToBoard()

	if strategy.GetNumBoardsEvaluated() != 4865609 {
		t.Errorf("should have gotten different number of legals")
	}
}
