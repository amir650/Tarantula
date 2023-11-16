package engine

type Pawn struct {
	BasePiece
}

func NewPawn(alliance Alliance, position int, isMoved bool) Pawn {
	return Pawn{BasePiece{alliance: alliance, position: position, isMoved: isMoved}}
}

func (pawn Pawn) GetAlliance() Alliance {
	return pawn.alliance
}

func (pawn Pawn) GetPiecePosition() int {
	return pawn.position
}

func (pawn Pawn) String() string {
	return PawnIdentifier
}

func (pawn Pawn) CalculateLegalMoves(board *Board) []Move {
	var legalMoves []Move
	var CandidateMoveCoordinates = []int{8, 16, 7, 9}
	for _, currentCandidateOffset := range CandidateMoveCoordinates {
		var candidateDestinationCoordinate = pawn.position + (pawn.alliance.GetDirection() * currentCandidateOffset)
		if !IsValidTileCoordinate(candidateDestinationCoordinate) {
			continue
		}
		t := board.GetTile(candidateDestinationCoordinate)

		if currentCandidateOffset == 8 && !t.IsOccupied() {
			if pawn.alliance.IsPawnPromotionSquare(candidateDestinationCoordinate) {
				var promotedToKnight = Knight{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
				var promotedToBishop = Bishop{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
				var promotedToRook = Rook{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
				var promotedToQueen = Queen{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
				legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToKnight))
				legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToBishop))
				legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToRook))
				legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToQueen))
			} else {
				legalMoves = append(legalMoves, NewPawnMove(board, pawn, candidateDestinationCoordinate))
			}
		} else if currentCandidateOffset == 16 && !pawn.isMoved &&
			((SecondRow[pawn.position] && pawn.alliance == BLACK) ||
				(SeventhRow[pawn.position] && pawn.alliance == WHITE)) {
			var behindCandidateDestinationCoordinate = pawn.position + (pawn.alliance.GetDirection() * 8)
			if !board.GetTile(candidateDestinationCoordinate).IsOccupied() &&
				!board.GetTile(behindCandidateDestinationCoordinate).IsOccupied() {
				legalMoves = append(legalMoves, NewPawnJump(board, pawn, candidateDestinationCoordinate))
			}
		} else if currentCandidateOffset == 7 &&
			!((EighthColumn[pawn.position] && pawn.alliance == WHITE) ||
				(FirstColumn[pawn.position] && pawn.alliance == BLACK)) {
			if board.GetTile(candidateDestinationCoordinate).IsOccupied() {
				var pieceOnCandidate = board.GetTile(candidateDestinationCoordinate).GetPiece()
				if pawn.alliance != pieceOnCandidate.GetAlliance() {
					if pawn.alliance.IsPawnPromotionSquare(candidateDestinationCoordinate) {
						var promotedToKnight = Knight{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
						var promotedToBishop = Bishop{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
						var promotedToRook = Rook{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
						var promotedToQueen = Queen{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
						legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToKnight))
						legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToBishop))
						legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToRook))
						legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToQueen))
					} else {
						legalMoves = append(legalMoves, NewPawnAttackMove(board, pawn, candidateDestinationCoordinate, board.GetPiece(candidateDestinationCoordinate)))
					}
				}
			} else if board.GetEnPassantPawn() != nil &&
				board.GetEnPassantPawn().GetPiecePosition() ==
					(pawn.position+int(pawn.alliance.GetOppositeDirection())) {
				var pieceOnCandidate = board.GetEnPassantPawn()
				if pawn.alliance != pieceOnCandidate.GetAlliance() {
					legalMoves = append(legalMoves, NewPawnEnPassantAttack(board, NewPawnAttackMove(board, pawn, candidateDestinationCoordinate, pieceOnCandidate)))
				}
			}
		} else if currentCandidateOffset == 9 &&
			!((FirstColumn[pawn.position] && pawn.alliance == WHITE) ||
				(EighthColumn[pawn.position] && pawn.alliance == BLACK)) {
			if board.GetTile(candidateDestinationCoordinate).IsOccupied() {
				if pawn.alliance != board.GetTile(candidateDestinationCoordinate).GetPiece().GetAlliance() {
					if pawn.alliance.IsPawnPromotionSquare(candidateDestinationCoordinate) {
						var promotedToKnight = Knight{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
						var promotedToBishop = Bishop{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
						var promotedToRook = Rook{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
						var promotedToQueen = Queen{BasePiece{pawn.alliance, candidateDestinationCoordinate, true}}
						legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToKnight))
						legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToBishop))
						legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToRook))
						legalMoves = append(legalMoves, NewPawnPromotion(NewPawnMove(board, pawn, candidateDestinationCoordinate), promotedToQueen))
					} else {
						legalMoves = append(legalMoves, NewPawnAttackMove(board, pawn, candidateDestinationCoordinate, board.GetPiece(candidateDestinationCoordinate)))
					}
				}
			} else if board.GetEnPassantPawn() != nil && board.GetEnPassantPawn().GetPiecePosition() ==
				(pawn.position-int(pawn.alliance.GetOppositeDirection())) {
				var pieceOnCandidate = board.GetEnPassantPawn()
				if pawn.alliance != pieceOnCandidate.GetAlliance() {
					//legalMoves = append(legalMoves, NewPawnEnPassantAttack(NewPawnAttackMove(boards, boards.gameBoard[pawn.position], boards.gameBoard[candidateDestinationCoordinate]), pieceOnCandidate))
				}
			}
		}
	}
	return legalMoves
}

func (pawn Pawn) MovePiece(m Move) Piece {
	return Pawn{BasePiece{pawn.GetAlliance(), m.GetTo(), true}}
}

func (pawn Pawn) Equals(other Piece) bool {
	p, ok := other.(Pawn)
	if ok {
		return pawn.GetPiecePosition() == p.GetPiecePosition() && pawn.GetAlliance() == p.GetAlliance()
	}
	return false
}

func (pawn Pawn) GetPieceValue() int {
	return 100
}

func (pawn Pawn) GetLocationBonus() int {
	a := pawn.GetAlliance()
	switch a {
	case WHITE:
		return WhitePawnPreferredCoordinates[pawn.GetPiecePosition()]
	case BLACK:
		return BlackPawnPreferredCoordinates[pawn.GetPiecePosition()]
	default:
		panic("wtf")
	}
}
