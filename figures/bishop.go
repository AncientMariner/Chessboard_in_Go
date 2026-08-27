package figures

var _ FigureBehaviour = (*Bishop)(nil)

type Bishop struct {
	Figure
}

func (bishop *Bishop) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool(0)
	dimension := getDimensionFromBoard(board)
	boardHash := ZobristHash(board)

	for i := 0; i < len(board) && len(board) == (dimension*dimension); i++ {
		if board[i] == emptyField {
			// Check validity first before doing any allocation
			if !isAnotherFigurePresentDiag(board, i, dimension) {
				// Get working buffer from pool
				outPtr := getBoardFromPool(dimension)
				out := *outPtr
				copy(out, board)

				h := boardHash
				h = placeAttackPlacesDiagonallyAbove(out, i, dimension, h)
				h = placeAttackPlacesDiagonallyBelow(out, i, dimension, h)
				h = ZobristUpdate(h, i, emptyField, bishop.GetName())
				out[i] = bishop.GetName()

				// Make permanent copy for storage
				permanent := make([]byte, len(out))
				copy(permanent, out)
				boards[h] = permanent

				// Return working buffer to pool
				boardPool.Put(outPtr)
			}
		}
	}
	return boards
}

func placeAttackPlacesDiagonallyBelow(out []byte, position int, dimension int, h uint64) uint64 {
	if position >= len(out) {
		return h
	}
	diagBelowRight := position + dimension + 1
	diagBelowLeft := position + dimension - 1

	currentLine := position/dimension + 1
	for lineBelow := currentLine + 1; lineBelow <= dimension; lineBelow++ {
		lineOfTheDiagBelowRight := diagBelowRight/dimension + 1
		lineOfTheDiagBelowLeft := diagBelowLeft/dimension + 1

		if lineBelow == lineOfTheDiagBelowRight && diagBelowRight < len(out) && position < len(out)-dimension && out[diagBelowRight] == emptyField {
			out[diagBelowRight] = attackPlace
			h = ZobristUpdate(h, diagBelowRight, emptyField, attackPlace)
		}
		diagBelowRight = diagBelowRight + dimension + 1

		if lineBelow == lineOfTheDiagBelowLeft && diagBelowLeft < len(out) && position < len(out)-dimension && out[diagBelowLeft] == emptyField {
			out[diagBelowLeft] = attackPlace
			h = ZobristUpdate(h, diagBelowLeft, emptyField, attackPlace)
		}
		diagBelowLeft = diagBelowLeft + dimension - 1
	}
	return h
}

func placeAttackPlacesDiagonallyAbove(out []byte, position int, dimension int, h uint64) uint64 {
	if position >= len(out) {
		return h
	}
	diagAboveLeft := position - dimension - 1
	diagAboveRight := position - dimension + 1

	currentLine := position/dimension + 1
	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		lineOfTheDiagAboveLeft := diagAboveLeft/dimension + 1
		lineOfTheDiagAboveRight := diagAboveRight/dimension + 1

		if lineOfTheDiagAboveLeft == lineAbove && position >= dimension && diagAboveLeft >= 0 && out[diagAboveLeft] == emptyField {
			out[diagAboveLeft] = attackPlace
			h = ZobristUpdate(h, diagAboveLeft, emptyField, attackPlace)
		}
		diagAboveLeft = diagAboveLeft - dimension - 1

		if lineOfTheDiagAboveRight == lineAbove && position >= dimension && diagAboveRight >= 0 && out[diagAboveRight] == emptyField {
			out[diagAboveRight] = attackPlace
			h = ZobristUpdate(h, diagAboveRight, emptyField, attackPlace)
		}
		diagAboveRight = diagAboveRight - dimension + 1
	}
	return h
}

func isAnotherFigurePresentDiag(out []byte, position int, dimension int) bool {
	isPiece := func(i int) bool { return out[i] != emptyField && out[i] != attackPlace }
	currentLine := position/dimension + 1

	diagAboveLeft := position - dimension - 1
	diagAboveRight := position - dimension + 1
	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		if diagAboveLeft/dimension+1 == lineAbove && diagAboveLeft >= 0 && isPiece(diagAboveLeft) {
			return true
		}
		diagAboveLeft -= dimension + 1
		if diagAboveRight/dimension+1 == lineAbove && diagAboveRight >= 0 && isPiece(diagAboveRight) {
			return true
		}
		diagAboveRight -= dimension - 1
	}

	diagBelowLeft := position + dimension - 1
	diagBelowRight := position + dimension + 1
	for lineBelow := currentLine + 1; lineBelow <= dimension; lineBelow++ {
		if diagBelowRight/dimension+1 == lineBelow && diagBelowRight < len(out) && isPiece(diagBelowRight) {
			return true
		}
		diagBelowRight += dimension + 1
		if diagBelowLeft/dimension+1 == lineBelow && diagBelowLeft < len(out) && isPiece(diagBelowLeft) {
			return true
		}
		diagBelowLeft += dimension - 1
	}
	return false
}

func (*Bishop) GetName() byte {
	return 'b'
}
