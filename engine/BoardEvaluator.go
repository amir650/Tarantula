package engine

import (
	"fmt"
	"strings"
)

const CheckMateBonus = 10_000
const CheckBonus = 45
const CastleBonus = 25
const MobilityMultiplier = 5
const AttackMultiplier = 1
const TwoBishopsBonus = 25

type BoardEvaluator interface {
	Evaluate(board *Board, depth int) int
}

type StandardBoardEvaluator struct {
}

func (evaluator StandardBoardEvaluator) Evaluate(board *Board, depth int) int {
	//fmt.Println(evaluationDetails(evaluator, board, depth))
	result := score(board.GetWhitePlayer(), depth) - score(board.GetBlackPlayer(), depth)
	return result
}

func score(player Player, depth int) int {
	return mobility(player) +
		kingThreats(player, depth) +
		attacks(player) +
		castle(player) +
		pieceEvaluations(player) +
		pawnStructure(player)
}

func pawnStructure(player Player) int {
	return 0
}

func castle(player Player) int {
	if player.IsCastled() {
		return CastleBonus
	} else {
		return 0
	}
}

func attacks(player Player) int {
	var attackScore = 0
	for _, move := range player.GetLegalMoves() {
		if move.IsAttack() {
			movedPiece := move.GetMovedPiece()
			attackedPiece := move.GetAttackedPiece()
			if movedPiece.GetPieceValue() <= attackedPiece.GetPieceValue() {
				attackScore++
			}
		}
	}
	return attackScore * AttackMultiplier
}

func kingThreats(player Player, depth int) int {

	if player.GetOpponent().IsInCheckMate() {
		return CheckMateBonus * depthBonus(depth)
	} else {
		return check(player)
	}
}

func check(player Player) int {
	return 0
}

func depthBonus(depth int) int {
	return 0
}

func mobility(player Player) int {
	return MobilityMultiplier * mobilityRatio(player)
}

func mobilityRatio(player Player) int {
	return len(player.GetLegalMoves()) * 10.0 / len(player.GetOpponent().GetLegalMoves())
}

func pieceEvaluations(player Player) int {
	var pieceValuationScore = 0
	var numBishops = 0
	for _, piece := range player.GetActivePieces() {
		pieceValuationScore += piece.GetPieceValue()
		_, ok := piece.(*Bishop)
		if ok {
			numBishops++
		}
	}
	if numBishops == 2 {
		return TwoBishopsBonus + pieceValuationScore
	} else {
		return pieceValuationScore
	}
}

func evaluationDetails(evaluator StandardBoardEvaluator, board *Board, depth int) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("White Mobility : %d\n", mobility(board.GetWhitePlayer())))
	sb.WriteString(fmt.Sprintf("White kingThreats : %d\n", kingThreats(board.GetWhitePlayer(), depth)))
	sb.WriteString(fmt.Sprintf("White attacks : %d\n", attacks(board.GetWhitePlayer())))
	sb.WriteString(fmt.Sprintf("White castle : %d\n", castle(board.GetWhitePlayer())))
	sb.WriteString(fmt.Sprintf("White pieceEval : %d\n", pieceEvaluations(board.GetWhitePlayer())))
	sb.WriteString(fmt.Sprintf("White pawnStructure : %d\n", pawnStructure(board.GetWhitePlayer())))
	sb.WriteString("---------------------\n")
	sb.WriteString(fmt.Sprintf("Black Mobility : %d\n", mobility(board.GetBlackPlayer())))
	sb.WriteString(fmt.Sprintf("Black kingThreats : %d\n", kingThreats(board.GetBlackPlayer(), depth)))
	sb.WriteString(fmt.Sprintf("Black attacks : %d\n", attacks(board.GetBlackPlayer())))
	sb.WriteString(fmt.Sprintf("Black castle : %d\n", castle(board.GetBlackPlayer())))
	sb.WriteString(fmt.Sprintf("Black pieceEval : %d\n", pieceEvaluations(board.GetBlackPlayer())))
	sb.WriteString(fmt.Sprintf("Black pawnStructure : %d\n", pawnStructure(board.GetBlackPlayer())))
	sb.WriteString("\n")
	//sb.WriteString(fmt.Sprintf("Final Score = %d", evaluator.Evaluate(board, depth)))
	return sb.String()
}
