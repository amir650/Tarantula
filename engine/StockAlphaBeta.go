package engine

//
//import (
//	"fmt"
//	"math"
//	"time"
//)
//
//type StockAlphaBeta struct {
//	boardsEvaluated int64
//	executionTime   int64
//	boardEvaluator  BoardEvaluator
//	searchDepth     int
//	frequencyTable  []BoardMoveCount
//	fqIndex         int
//	showDebug       bool
//}
//
//func NewStockAlphaBeta(searchDepth int, showDebug bool) MoveStrategy {
//	return &StockAlphaBeta{searchDepth: searchDepth, boardEvaluator: StandardBoardEvaluator{}, showDebug: showDebug}
//}
//
//func (strategy *StockAlphaBeta) GetNumBoardsEvaluated() int64 {
//	return strategy.boardsEvaluated
//}
//
//func (strategy *StockAlphaBeta) Execute(board *Board) Move {
//	startTime := time.Now()
//	bestMove := NewNullMove()
//	highestSeenValue := math.MinInt32
//	lowestSeenValue := math.MaxInt32
//
//	fmt.Println(board.GetCurrentPlayer().String(), "THINKING with depth =", strategy.searchDepth)
//	strategy.frequencyTable = make([]BoardMoveCount, len(board.GetCurrentPlayer().GetLegalMoves()))
//	moveCounter := 1
//	numMoves := len(board.GetCurrentPlayer().GetLegalMoves())
//
//	for _, move := range board.GetCurrentPlayer().GetLegalMoves() {
//		moveTransition := board.GetCurrentPlayer().MakeMove(move)
//
//		if moveTransition.GetMoveStatus() == Done {
//			strategy.frequencyTable[strategy.fqIndex] = NewBoardCountRow(move)
//			currentValue := 0
//			if board.GetCurrentPlayer().GetAlliance() == WHITE {
//				currentValue = strategy.minimize(moveTransition.GetToBoard(), strategy.searchDepth-1)
//			} else {
//				currentValue = strategy.maximize(moveTransition.GetToBoard(), strategy.searchDepth-1)
//			}
//
//			if board.GetCurrentPlayer().GetAlliance() == WHITE && currentValue >= highestSeenValue {
//				highestSeenValue = currentValue
//				bestMove = move
//			} else if !(board.GetCurrentPlayer().GetAlliance() == WHITE) && (currentValue <= lowestSeenValue) {
//				lowestSeenValue = currentValue
//				bestMove = move
//			}
//			if strategy.showDebug {
//				fmt.Printf("\tMiniMax analyzing move (%d/%d) %v scores %d\n", moveCounter, numMoves, move, currentValue)
//			}
//
//		} else {
//			if strategy.showDebug {
//				fmt.Printf("\tMiniMax can't execute move (%d/%d) %v\n", moveCounter, numMoves, move)
//			}
//		}
//		moveCounter++
//	}
//
//	strategy.executionTime = time.Since(startTime).Milliseconds()
//	boardsPerSecond := float64(strategy.boardsEvaluated) / float64(strategy.executionTime) * 1000
//
//	fmt.Printf("%s SELECTS %v [#boards = %d time taken = %d ms, rate = %.1f]\n",
//		board.GetCurrentPlayer(), bestMove, strategy.boardsEvaluated, strategy.executionTime, boardsPerSecond)
//
//	return bestMove
//}
//
//func (strategy *StockAlphaBeta) maximize(board *Board,
//										 depth int,
//										 highest int,,
//										 lowest int) int {
//	if depth == 0 || isEndGameScenario(board){
//		strategy.boardsEvaluated++
//		return strategy.boardEvaluator.Evaluate(board, depth)
//	}
//	lowestSeenValue := lowest
//	for _, move := range board.GetCurrentPlayer().GetLegalMoves() {
//		moveTransition := board.GetCurrentPlayer().MakeMove(move)
//		if moveTransition.GetMoveStatus() == Done {
//			currentValue := min(lowestSeenValue, strategy.minimize(moveTransition.GetToBoard(), depth-1)
//			if currentValue <= lowestSeenValue {
//				lowestSeenValue = currentValue
//			}
//		}
//	}
//	return lowestSeenValue
//}
//
//func (strategy *StockAlphaBeta) minimize(board *Board, depth int) int {
//	if depth == 0 {
//		strategy.boardsEvaluated++
//		strategy.frequencyTable[strategy.fqIndex].increment()
//		return strategy.boardEvaluator.Evaluate(board, depth)
//	}
//	if isEndGameScenario(board) {
//		return strategy.boardEvaluator.Evaluate(board, depth)
//	}
//	highestSeenValue := math.MinInt32
//	for _, move := range board.GetCurrentPlayer().GetLegalMoves() {
//		moveTransition := board.GetCurrentPlayer().MakeMove(move)
//		if moveTransition.GetMoveStatus() == Done {
//			currentValue := strategy.maximize(moveTransition.GetToBoard(), depth-1)
//			if currentValue >= highestSeenValue {
//				highestSeenValue = currentValue
//			}
//		}
//	}
//	return highestSeenValue
//
//}
