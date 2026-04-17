package figures

var _ FigureBehaviour = (*Knight)(nil)

type Knight struct {
	Figure
}

func (knight *Knight) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool(getCountOfEmptyPlaces(board))
	dimension := getDimensionFromBoard(board)

	for i := 0; i < len(board) && len(board) == (dimension*dimension); i++ {
		if board[i] == emptyField {
			// Check validity first before doing any allocation
			if !isAnotherFigurePresentBelow(board, i, dimension) && !isAnotherFigurePresentAbove(board, i, dimension) {
				// Get working buffer from pool
				outPtr := getBoardFromPool(dimension)
				out := *outPtr
				copy(out, board)

				placeAttackPlacesBelow(out, i, dimension)
				placeAttackPlacesAbove(out, i, dimension)
				out[i] = knight.GetName()

				// Make permanent copy for storage
				permanent := make([]byte, len(out))
				copy(permanent, out)
				boards[GenerateHash(permanent)] = permanent

				// Return working buffer to pool
				boardPool.Put(outPtr)
			}
		}
	}
	return boards
}

func placeAttackPlacesBelow(out []byte, position int, dimension int) {
	if position >= len(out) || position == dimension || position%dimension == dimension {
		return
	}
	currentLine := position/dimension + 1
	lineBelow := currentLine + 1

	positionBelowLineLeft := position + dimension - 2
	belowLineLeft := positionBelowLineLeft/dimension + 1
	if lineBelow == belowLineLeft && positionBelowLineLeft < len(out) && out[positionBelowLineLeft] == emptyField {
		out[positionBelowLineLeft] = attackPlace
	}
	positionBelowLineRight := position + dimension + 2
	belowLineRight := positionBelowLineRight/dimension + 1
	if lineBelow == belowLineRight && positionBelowLineRight < len(out) && out[positionBelowLineRight] == emptyField {
		out[positionBelowLineRight] = attackPlace
	}

	line2Below := currentLine + 1 + 1

	positionBelow2LinesLeft := position + 2*dimension - 1
	below2LineLeft := positionBelow2LinesLeft/dimension + 1
	if line2Below == below2LineLeft && positionBelow2LinesLeft < len(out) && out[positionBelow2LinesLeft] == emptyField {
		out[positionBelow2LinesLeft] = attackPlace
	}

	positionBelow2LinesRight := position + 2*dimension + 1
	below2LineRight := positionBelow2LinesRight/dimension + 1
	if line2Below == below2LineRight && positionBelow2LinesRight < len(out) && out[positionBelow2LinesRight] == emptyField {
		out[positionBelow2LinesRight] = attackPlace
	}
}

func placeAttackPlacesAbove(out []byte, position int, dimension int) {
	if position >= len(out) || position == dimension || position%dimension == dimension {
		return
	}
	currentLine := position/dimension + 1

	positionAboveLineLeft := position - dimension - 2
	positionAboveLineRight := position - dimension + 2

	lineAbove := currentLine - 1
	aboveLineLeft := positionAboveLineLeft/dimension + 1
	aboveLineRight := positionAboveLineRight/dimension + 1

	if aboveLineLeft == lineAbove && positionAboveLineLeft >= 0 && out[positionAboveLineLeft] == emptyField {
		out[positionAboveLineLeft] = attackPlace
	}
	if aboveLineRight == lineAbove && positionAboveLineRight >= 0 && out[positionAboveLineRight] == emptyField {
		out[positionAboveLineRight] = attackPlace
	}

	positionAbove2LinesLeft := position - 2*dimension - 1
	positionAbove2LinesRight := position - 2*dimension + 1

	above2LineLeft := positionAbove2LinesLeft/dimension + 1
	above2LineRight := positionAbove2LinesRight/dimension + 1
	line2Above := currentLine - 1 - 1
	if above2LineLeft == line2Above && positionAbove2LinesLeft >= 0 {
		out[positionAbove2LinesLeft] = attackPlace
	}
	if above2LineRight == line2Above && positionAbove2LinesRight >= 0 {
		out[positionAbove2LinesRight] = attackPlace
	}
}

func isAnotherFigurePresentBelow(out []byte, position int, dimension int) bool {
	currentLine := position/dimension + 1
	var numbersToCheck []int

	positionBelowLineLeft := position + dimension - 2
	positionBelowLineRight := position + dimension + 2

	lineBelow := currentLine + 1
	belowLineLeft := positionBelowLineLeft/dimension + 1
	belowLineRight := positionBelowLineRight/dimension + 1

	if lineBelow == belowLineLeft && positionBelowLineLeft < len(out) {
		numbersToCheck = append(numbersToCheck, positionBelowLineLeft)
	}
	if lineBelow == belowLineRight && positionBelowLineRight < len(out) {
		numbersToCheck = append(numbersToCheck, positionBelowLineRight)
	}

	positionBelow2LinesLeft := position + 2*dimension - 1
	positionBelow2LinesRight := position + 2*dimension + 1

	below2LineLeft := positionBelow2LinesLeft/dimension + 1
	below2LineRight := positionBelow2LinesRight/dimension + 1
	line2Below := currentLine + 1 + 1
	if line2Below == below2LineRight && positionBelow2LinesRight < len(out) {
		numbersToCheck = append(numbersToCheck, positionBelow2LinesRight)
	}
	if line2Below == below2LineLeft && positionBelow2LinesLeft < len(out) {
		numbersToCheck = append(numbersToCheck, positionBelow2LinesLeft)
	}

	for _, number := range numbersToCheck {
		if out[number] != emptyField && out[number] != attackPlace {
			return true
		}
	}
	return false
}

func isAnotherFigurePresentAbove(out []byte, position int, dimension int) bool {

	currentLine := position/dimension + 1
	var numbersToCheck []int

	positionAboveLineLeft := position - dimension - 2
	positionAboveLineRight := position - dimension + 2

	lineAbove := currentLine - 1
	aboveLineLeft := positionAboveLineLeft/dimension + 1
	aboveLineRight := positionAboveLineRight/dimension + 1

	if aboveLineLeft == lineAbove && positionAboveLineLeft >= 0 {
		numbersToCheck = append(numbersToCheck, positionAboveLineLeft)
	}
	if aboveLineRight == lineAbove && positionAboveLineRight >= 0 {
		numbersToCheck = append(numbersToCheck, positionAboveLineRight)
	}

	positionAbove2LinesLeft := position - 2*dimension - 1
	positionAbove2LinesRight := position - 2*dimension + 1

	above2LineLeft := positionAbove2LinesLeft/dimension + 1
	above2LineRight := positionAbove2LinesRight/dimension + 1
	line2Above := currentLine - 1 - 1
	if above2LineLeft == line2Above && positionAbove2LinesLeft >= 0 {
		numbersToCheck = append(numbersToCheck, positionAbove2LinesLeft)
	}
	if above2LineRight == line2Above && positionAbove2LinesRight >= 0 {
		numbersToCheck = append(numbersToCheck, positionAbove2LinesRight)
	}

	for _, number := range numbersToCheck {
		if out[number] != emptyField && out[number] != attackPlace {
			return true
		}
	}
	return false
}

func (*Knight) GetName() byte {
	return 'n'
}
