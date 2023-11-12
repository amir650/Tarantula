package engine

const NumTiles = 64

const KingIdentifier = "K"
const QueenIdentifier = "Q"
const RookIdentifier = "R"
const BishopIdentifier = "B"
const KnightIdentifier = "N"
const PawnIdentifier = "P"

func IsValidTileCoordinate(coordinate int) bool {
	return coordinate >= 0 && coordinate < NumTiles
}

func createRowValues(rowNumber int) []bool {
	rowValues := make([]bool, 64)
	rowNumber -= 1
	limit := 8*rowNumber + 8
	for i := 8 * rowNumber; i < limit; i++ {
		rowValues[i] = true
	}
	return rowValues
}

func createColumnValues(columnNumber int) []bool {
	columnValues := make([]bool, 64)
	for i := columnNumber - 1; i < 64; i += 8 {
		columnValues[i] = true
	}
	return columnValues
}

func isEndGameScenario(board *Board) bool {
	return board.GetCurrentPlayer().IsInCheckMate() || board.GetCurrentPlayer().IsInStaleMate()
}

var FirstColumn = createColumnValues(1)
var SecondColumn = createColumnValues(2)
var SeventhColumn = createColumnValues(7)
var EighthColumn = createColumnValues(8)
var SecondRow = createRowValues(2)
var SeventhRow = createRowValues(7)

var WhitePawnPreferredCoordinates = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	90, 90, 90, 90, 90, 90, 90, 90,
	30, 30, 40, 60, 60, 40, 30, 30,
	10, 10, 20, 40, 40, 20, 10, 10,
	5, 5, 10, 20, 20, 10, 5, 5,
	0, 0, 0, -10, -10, 0, 0, 0,
	5, -5, -10, 0, 0, -10, -5, 5,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var BlackPawnPreferredCoordinates = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	5, -5, -10, 0, 0, -10, -5, 5,
	0, 0, 0, -10, -10, 0, 0, 0,
	5, 5, 10, 20, 20, 10, 5, 5,
	10, 10, 20, 40, 40, 20, 10, 10,
	30, 30, 40, 60, 60, 40, 30, 30,
	90, 90, 90, 90, 90, 90, 90, 90,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var WhiteKnightPreferredCoordinates = [64]int{
	-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50,
}

var BlackKnightPreferredCoordinates = [64]int{
	-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50,
}

var WhiteBishopPreferredCoordinates = [64]int{
	-20, -10, -10, -10, -10, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 10, 10, 5, 0, -10,
	-10, 5, 5, 10, 10, 5, 5, -10,
	-10, 0, 10, 15, 15, 10, 0, -10,
	-10, 10, 10, 10, 10, 10, 10, -10,
	-10, 5, 0, 0, 0, 0, 5, -10,
	-20, -10, -10, -10, -10, -10, -10, -20,
}

var BlackBishopPreferredCoordinates = [64]int{
	-20, -10, -10, -10, -10, -10, -10, -20,
	-10, 5, 0, 0, 0, 0, 5, -10,
	-10, 10, 10, 10, 10, 10, 10, -10,
	-10, 0, 10, 15, 15, 10, 0, -10,
	-10, 5, 10, 15, 15, 10, 5, -10,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-20, -10, -10, -10, -10, -10, -10, -20,
}

var WhiteRookPreferredCoordinates = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 20, 20, 20, 20, 20, 20, 5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	0, 0, 0, 5, 5, 0, 0, 0,
}

var BlackRookPreferredCoordinates = [64]int{
	0, 0, 0, 5, 5, 0, 0, 0,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	5, 20, 20, 20, 20, 20, 20, 5,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var WhiteQueenPreferredCoordinates = [64]int{
	-20, -10, -10, -5, -5, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 5, 5, 5, 0, -10,
	-5, 0, 5, 10, 10, 5, 0, -5,
	-5, 0, 5, 10, 10, 5, 0, -5,
	-10, 0, 5, 5, 5, 5, 0, -10,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-20, -10, -10, -5, -5, -10, -10, -20,
}

var BlackQueenPreferredCoordinates = [64]int{
	-20, -10, -10, -5, -5, -10, -10, -20,
	-10, 0, 5, 0, 0, 0, 0, -10,
	-10, 5, 5, 5, 5, 5, 0, -10,
	-5, 0, 5, 10, 10, 5, 0, -5,
	-5, 0, 5, 10, 10, 5, 0, -5,
	-10, 0, 5, 5, 5, 5, 0, -10,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-20, -10, -10, -5, -5, -10, -10, -20,
}

var WhiteKingPreferredCoordinates = [64]int{
	-50, -30, -30, -30, -30, -30, -30, -50,
	-30, -30, 0, 0, 0, 0, -30, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -20, -10, 0, 0, -10, -20, -30,
	-50, -40, -30, -20, -20, -30, -40, -50,
}

var BlackKingPreferredCoordinates = [64]int{
	-50, -40, -30, -20, -20, -30, -40, -50,
	-30, -20, -10, 0, 0, -10, -20, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -30, 0, 0, 0, 0, -30, -30,
	-50, -30, -30, -30, -30, -30, -30, -50,
}
