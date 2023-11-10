package engine

type Alliance int

const (
	WHITE Alliance = iota
	BLACK
)

func (a Alliance) GetDirection() int {
	switch a {
	case WHITE:
		return -1
	case BLACK:
		return 1
	default:
		return 0
	}
}

func (a Alliance) String() string {
	switch a {
	case WHITE:
		return "White"
	case BLACK:
		return "Black"
	default:
		return "Unknown"
	}
}

func (a Alliance) IsPawnPromotionSquare(coordinate int) bool {
	return false
}

func (a Alliance) GetOppositeDirection() Alliance {
	switch a {
	case WHITE:
		return BLACK
	case BLACK:
		return WHITE
	default:
		return -1
	}
}
