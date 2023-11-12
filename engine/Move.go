package engine

type Move interface {
	GetBoard() *Board
	GetFrom() int
	GetTo() int
	IsAttack() bool
	GetMovedPiece() Piece
	GetAttackedPiece() Piece
	String() string
	Execute() *Board
}

type MoveStatus int

const (
	Done MoveStatus = iota
	IllegalMove
	LeavesPlayerInCheck
)

func (m MoveStatus) isDone() bool {
	switch m {
	case Done:
		return true
	default:
		return false
	}
}

type MoveTransition struct {
	FromBoard      *Board
	ToBoard        *Board
	TransitionMove *Move
	MoveStatus     MoveStatus
}

func (t MoveTransition) GetMoveStatus() MoveStatus {
	return t.MoveStatus
}

func (t MoveTransition) String() string {
	return t.ToBoard.String()
}

func (t MoveTransition) GetToBoard() *Board {
	return t.ToBoard
}

func calculateAttacksOnPosition(position int,
	moves *[]Move) []Move {
	var attacks []Move
	for _, m := range *moves {
		if m.GetTo() == position {
			attacks = append(attacks, m)
		}
	}
	return attacks
}
