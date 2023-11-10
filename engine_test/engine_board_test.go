package engine_test

import (
	"Tarantula/engine"
	"testing"
)

func TestInitialBoardConfiguration(t *testing.T) {
	board := engine.CreateStandardChessBoard()

	//first white
	if board.GetCurrentPlayer() != board.GetWhitePlayer() {
		t.Errorf("Current Player should be white on the intial board!")
	}

	if len(board.GetCurrentPlayer().GetLegalMoves()) != 20 {
		t.Errorf("Initial Board should have 20 legal moves.  Instead got %d", len(board.GetCurrentPlayer().GetLegalMoves()))
	}

	if board.GetCurrentPlayer().IsInCheck() {
		t.Errorf("Initial Board no one should be in check! Board should have 20 legal moves")
	}

	if board.GetCurrentPlayer().IsInCheckMate() {
		t.Errorf("Initial Board no one should be in checkmate! Board should have 20 legal moves")
	}

	if board.GetCurrentPlayer().IsInStaleMate() {
		t.Errorf("Initial Board no one should be in stalemate!")
	}

	if board.GetCurrentPlayer().IsCastled() {
		t.Errorf("Initial Board no one should be castled!")
	}

	//now black
	if len(board.GetCurrentPlayer().GetOpponent().GetLegalMoves()) != 20 {
		t.Errorf("Initial Board should have 20 legal moves.  Instead got %d", len(board.GetCurrentPlayer().GetLegalMoves()))
	}

	if board.GetCurrentPlayer().GetOpponent().IsInCheck() {
		t.Errorf("Initial Board no one should be in check! Board should have 20 legal moves")
	}

	if board.GetCurrentPlayer().GetOpponent().IsInCheckMate() {
		t.Errorf("Initial Board no one should be in checkmate! Board should have 20 legal moves")
	}

	if board.GetCurrentPlayer().GetOpponent().IsInStaleMate() {
		t.Errorf("Initial Board no one should be in stalemate!")
	}

	if board.GetCurrentPlayer().GetOpponent().IsCastled() {
		t.Errorf("Initial Board no one should be castled!")
	}

	boardEvaluator := engine.StandardBoardEvaluator{}
	boardScore := boardEvaluator.Evaluate(board, 6)

	if boardScore != 0 {
		t.Errorf("Not scored symmetrically")
	}

}

func TestFoolsMate(t *testing.T) {
	board := engine.CreateStandardChessBoard()

	bt := board.GetCurrentPlayer().MakeMove(engine.GetMove(board, "f2", "f3"))
	if bt.GetMoveStatus() != engine.Done {
		t.Errorf("u maki da fuck up")
	}

	board = bt.GetToBoard()
	bt = board.GetCurrentPlayer().MakeMove(engine.GetMove(board, "e7", "e5"))
	if bt.GetMoveStatus() != engine.Done {
		t.Errorf("u maki da fuck up")
	}
	board = bt.GetToBoard()

	board = bt.GetToBoard()
	bt = board.GetCurrentPlayer().MakeMove(engine.GetMove(board, "g2", "g4"))
	if bt.GetMoveStatus() != engine.Done {
		t.Errorf("u maki da fuck up")
	}
	board = bt.GetToBoard()

	board = bt.GetToBoard()
	bt = board.GetCurrentPlayer().MakeMove(engine.GetMove(board, "d8", "h4"))
	if bt.GetMoveStatus() != engine.Done {
		t.Errorf("u maki da fuck up")
	}
	board = bt.GetToBoard()

	if !board.GetCurrentPlayer().IsInCheckMate() {
		t.Errorf("u maki da fuck up")
	}
	board = bt.GetToBoard()

}

func TestTwoBishopsMate(t *testing.T) {

	builder := engine.NewBoardBuilder()

	builder.SetPiece(engine.NewKing(engine.BLACK, 7, false, false, false, false))
	builder.SetPiece(engine.NewPawn(engine.BLACK, 8, false))
	builder.SetPiece(engine.NewPawn(engine.BLACK, 10, false))
	builder.SetPiece(engine.NewPawn(engine.BLACK, 15, false))
	builder.SetPiece(engine.NewPawn(engine.BLACK, 17, false))

	builder.SetPiece(engine.NewBishop(engine.WHITE, 40, false))
	builder.SetPiece(engine.NewBishop(engine.WHITE, 48, false))
	builder.SetPiece(engine.NewKing(engine.WHITE, 53, false, false, false, false))

	builder.SetMoveMaker(engine.WHITE)

	board := builder.Build()

	bt := board.GetCurrentPlayer().MakeMove(engine.GetMove(board, "a3", "b2"))
	if bt.GetMoveStatus() != engine.Done {
		t.Errorf("u maki da fuck up")
	}
	board = bt.GetToBoard()
	if !board.GetCurrentPlayer().IsInCheckMate() {
		t.Errorf("u maki da fuck up")
	}

}

func TestQueenKnightMate(t *testing.T) {

	builder := engine.NewBoardBuilder()

	builder.SetPiece(engine.NewKing(engine.BLACK, 4, false, false, false, false))

	builder.SetPiece(engine.NewQueen(engine.WHITE, 15, false))
	builder.SetPiece(engine.NewKnight(engine.WHITE, 29, false))
	builder.SetPiece(engine.NewPawn(engine.WHITE, 55, false))
	builder.SetPiece(engine.NewKing(engine.WHITE, 60, false, false, false, false))

	builder.SetMoveMaker(engine.WHITE)

	board := builder.Build()

	bt := board.GetCurrentPlayer().MakeMove(engine.GetMove(board, "h7", "e7"))

	board = bt.GetToBoard()

	if bt.GetMoveStatus() != engine.Done {
		t.Errorf("u maki da fuck up")
	}

	if !board.GetCurrentPlayer().IsInCheckMate() {
		t.Errorf("u maki da fuck up")
	}

}
