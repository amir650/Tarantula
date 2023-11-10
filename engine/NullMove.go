package engine

type NullMove struct {
	BaseMove
}

func NewNullMove() Move {
	return NullMove{}
}

func (n NullMove) String() string {
	return "Null Move"
}
