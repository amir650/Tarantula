package main

import (
	"Tarantula/engine"
	"testing"
)

// Assuming you have a setup function to create a Board instance
func setupBoard() *engine.Board {
	board := engine.CreateStandardChessBoard()
	return board
}

func BenchmarkSerialLegalMoveCalculationFunc(b *testing.B) {
	board := setupBoard()
	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		board.CalculateLegalMoves()
	}
}

func BenchmarkParallelLegalMoveCalculationFunc(b *testing.B) {
	board := setupBoard()
	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		board.CalculateLegalMovesParallel()
	}
}

func BenchmarkParallelLegalMoveCalculationFuncGPT(b *testing.B) {
	board := setupBoard()
	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		board.CalculateLegalMovesParallelGPT()
	}
}
