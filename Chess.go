package main

import (
	"Tarantula/engine"
	"bufio"
	"fmt"
	_ "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
)

func ConvertNanoToMilli(nanoSeconds uint64) float64 {
	return float64(nanoSeconds) / 1000000.0
}

func printGCStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Number of GC Cycles: %d\n", m.NumGC)
	fmt.Printf("Total GC Pause Time: %v millis\n", ConvertNanoToMilli(m.PauseTotalNs))
	// Other GC-related stats
}

func main() {

	go func() {
		fmt.Println("Starting pprof server at http://localhost:6060/debug/pprof/")
		// Log errors that http.ListenAndServe returns
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()
	selfPlay(true)
}

func selfPlay(forever bool) {
	board := engine.CreateStandardChessBoard()
	strategy := engine.NewStockAlphaBeta(6, true)
	fmt.Println(board.String())
	m := strategy.Execute(board)
	printGCStats()
	mt := board.GetCurrentPlayer().MakeMove(m)
	board = mt.GetToBoard()
	if forever {
		for {
			fmt.Println(board.String())
			m := strategy.Execute(board)
			mt := board.GetCurrentPlayer().MakeMove(m)
			board = mt.GetToBoard()
			printGCStats()
		}
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

		// Print the input
		move := engine.GetMove(board, source, dest)
		fmt.Printf("You entered: %s %s : [%s]\n", source, dest, move)

		bt := board.GetCurrentPlayer().MakeMove(move)
		board = bt.GetToBoard()
	}
}
