package main

import (
	"Tarantula/engine"
	"bufio"
	"fmt"
	"os"
)

func main() {
	selfPlay()
}

func selfPlay() {
	board := engine.CreateStandardChessBoard()
	strategy := engine.NewMinimax(4, true)
	for {
		fmt.Println(board.String())
		m := strategy.Execute(board)
		mt := board.GetCurrentPlayer().MakeMove(m)
		board = mt.GetToBoard()
	}
}

func playInteractiveGame() {
	board := engine.CreateStandardChessBoard()
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println(board.String())

		fmt.Print("> Source:")
		scanner.Scan() // Read a line of input from the user
		source := scanner.Text()

		fmt.Print("> Dest:")
		scanner.Scan() // Read a line of input from the user
		dest := scanner.Text()

		for _, move := range board.GetAllLegalMoves() {
			fmt.Printf("debug print %s: %s\n", move.GetMovedPiece(), move)
		}

		// Print the input
		move := engine.GetMove(board, source, dest)
		fmt.Printf("You entered: %s %s : [%s]\n", source, dest, move)

		bt := board.GetCurrentPlayer().MakeMove(move)
		board = bt.GetToBoard()
	}
}
