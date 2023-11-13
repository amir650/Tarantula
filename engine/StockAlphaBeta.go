package engine

import (
	"fmt"
	"math"
	"time"
)

const MaxQuiescence = 5000
const quiescenceThreshold = 500

type StockAlphaBeta struct {
	boardEvaluator  BoardEvaluator
	searchDepth     int
	boardsEvaluated int64
	quiescenceCount int
	executionTime   int64
	frequencyTable  []BoardMoveCount
	fqIndex         int
	showDebug       bool
}

func NewStockAlphaBeta(searchDepth int, showDebug bool) MoveStrategy {
	return &StockAlphaBeta{searchDepth: searchDepth, boardEvaluator: StandardBoardEvaluator{}, showDebug: showDebug}
}

func (strategy *StockAlphaBeta) GetNumBoardsEvaluated() int64 {
	return strategy.boardsEvaluated
}

func (strategy *StockAlphaBeta) Execute(board *Board) Move {
	startTime := time.Now()
	bestMove := NewNullMove()
	highestSeenValue := math.MinInt32
	lowestSeenValue := math.MaxInt32
	currentValue := 0
	fmt.Println(board.GetCurrentPlayer().String(), "THINKING with depth =", strategy.searchDepth)
	moveCounter := 1
	numMoves := len(board.GetCurrentPlayer().GetLegalMoves())

	for _, move := range board.GetCurrentPlayer().GetLegalMoves() {
		strategy.quiescenceCount = 0
		moveTransition := board.GetCurrentPlayer().MakeMove(move)
		if moveTransition.GetMoveStatus() == Done {
			if board.GetCurrentPlayer().GetAlliance() == WHITE {
				currentValue = strategy.minimize(moveTransition.GetToBoard(), strategy.searchDepth-1, highestSeenValue, lowestSeenValue)
				if currentValue >= highestSeenValue {
					highestSeenValue = currentValue
					bestMove = move
					if moveTransition.GetToBoard().GetBlackPlayer().IsInCheckMate() {
						break
					}
				}
			} else {
				currentValue = strategy.maximize(moveTransition.GetToBoard(), strategy.searchDepth-1, highestSeenValue, lowestSeenValue)
				if currentValue <= lowestSeenValue {
					lowestSeenValue = currentValue
					bestMove = move
					if moveTransition.GetToBoard().GetWhitePlayer().IsInCheckMate() {
						break
					}
				}
			}
			if strategy.showDebug {
				fmt.Printf("\tMiniMax analyzing move (%d/%d) %v scores %d\n", moveCounter, numMoves, move, currentValue)
			}
		} else {
			if strategy.showDebug {
				fmt.Printf("\tMiniMax can't execute move (%d/%d) %v\n", moveCounter, numMoves, move)
			}
		}
		moveCounter++
	}

	strategy.executionTime = time.Since(startTime).Milliseconds()
	boardsPerSecond := float64(strategy.boardsEvaluated) / float64(strategy.executionTime) * 1000

	fmt.Printf("%s SELECTS %v [#boards = %d time taken = %d ms, rate = %.1f]\n",
		board.GetCurrentPlayer(), bestMove, strategy.boardsEvaluated, strategy.executionTime, boardsPerSecond)

	return bestMove
}

func (strategy *StockAlphaBeta) maximize(board *Board,
	depth int,
	highest int,
	lowest int) int {
	if depth == 0 || isEndGameScenario(board) {
		strategy.boardsEvaluated++
		return strategy.boardEvaluator.Evaluate(board, depth)
	}
	currentHighest := highest
	for _, move := range board.GetCurrentPlayer().GetLegalMoves() {
		moveTransition := board.GetCurrentPlayer().MakeMove(move)
		if moveTransition.GetMoveStatus() == Done {
			toBoard := moveTransition.GetToBoard()
			currentHighest = max(currentHighest,
				strategy.minimize(toBoard, depth-1, currentHighest, lowest))
			if currentHighest >= lowest {
				return lowest
			}
		}
	}
	return currentHighest
}

func calculateQuiescenceDepth(strategy *StockAlphaBeta,
	board *Board,
	depth int) int {
	if depth == 1 && strategy.quiescenceCount < MaxQuiescence {
		lastFourMoves := LastNBoards(board, 4)
		if lastFourMoves != nil {
			current := strategy.boardEvaluator.Evaluate(lastFourMoves[0], 0)
			previous := strategy.boardEvaluator.Evaluate(lastFourMoves[1], 0)
			pCurrent := strategy.boardEvaluator.Evaluate(lastFourMoves[2], 0)
			pPrevious := strategy.boardEvaluator.Evaluate(lastFourMoves[3], 0)
			delta1 := math.Abs(float64(current) - float64(previous))
			delta2 := math.Abs(float64(pCurrent) - float64(pPrevious))
			if delta1 > delta2 && (delta1-delta2) > quiescenceThreshold {
				strategy.quiescenceCount++
				return depth
			}
		}
	}
	return depth - 1
}

func LastNBoards(board *Board, n int) []*Board {
	boardHistory := make([]*Board, 0)
	boardHistory = append(boardHistory, board)
	i := 0
	currentBoard := board.GetTransitionMove().GetBoard()
	for currentBoard != nil && i < n {
		boardHistory = append(boardHistory, currentBoard)
		currentBoard = currentBoard.GetTransitionMove().GetBoard()
		i++
	}
	return boardHistory
}

func (strategy *StockAlphaBeta) minimize(board *Board,
	depth int,
	highest int,
	lowest int) int {
	if depth == 0 || isEndGameScenario(board) {
		strategy.boardsEvaluated++
		return strategy.boardEvaluator.Evaluate(board, depth)
	}
	currentLowest := lowest
	for _, move := range board.GetCurrentPlayer().GetLegalMoves() {
		moveTransition := board.GetCurrentPlayer().MakeMove(move)
		if moveTransition.GetMoveStatus() == Done {
			currentLowest = min(currentLowest, strategy.maximize(moveTransition.GetToBoard(), depth-1, highest, currentLowest))
			if currentLowest <= highest {
				return highest
			}
		}
	}
	return currentLowest

}
