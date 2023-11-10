package engine

type BasePlayer struct {
	board      *Board
	legalMoves *[]Move
	playerKing *King
	inCheck    bool
}

func (p BasePlayer) hasEscapeMoves() bool {
	for _, m := range p.GetLegalMoves() {
		result := p.MakeMove(m)
		if result.GetMoveStatus() == Done {
			return true
		}
	}
	return false
}

func (p BasePlayer) GetActivePieces() []Piece {
	panic("no!")
}

func (p BasePlayer) GetLegalMoves() []Move {
	return *p.legalMoves
}

func (p BasePlayer) GetAlliance() Alliance {
	return p.playerKing.GetAlliance()
}

func (p BasePlayer) GetOpponent() Player {
	panic("no!")
}

func (p BasePlayer) IsInCheck() bool {
	return p.inCheck
}

func (p BasePlayer) IsInCheckMate() bool {
	return p.inCheck && !p.hasEscapeMoves()
}

func (p BasePlayer) IsInStaleMate() bool {
	return !p.inCheck && !p.hasEscapeMoves()
}

func (p BasePlayer) MakeMove(move Move) *MoveTransition {
	for _, m := range p.GetLegalMoves() {
		if m.GetFrom() == move.GetFrom() &&
			m.GetTo() == m.GetTo() {
			var toBoard = move.Execute()
			if toBoard.GetCurrentPlayer().GetOpponent().IsInCheck() {
				return &MoveTransition{FromBoard: p.board, ToBoard: p.board, TransitionMove: &move, MoveStatus: LeavesPlayerInCheck}
			} else {
				return &MoveTransition{FromBoard: p.board, ToBoard: toBoard, TransitionMove: &move, MoveStatus: Done}
			}
		}
	}
	return &MoveTransition{FromBoard: p.board, ToBoard: p.board, TransitionMove: &move, MoveStatus: IllegalMove}
}

func (p BasePlayer) IsCastled() bool {
	return p.playerKing.isCastled
}

func (p BasePlayer) GetBoard() *Board {
	return p.board
}

func (p BasePlayer) String() string {
	return p.GetAlliance().String()
}
