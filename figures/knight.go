package figures

var _ FigureBehaviour = (*Knight)(nil)

type Knight struct {
	Figure
}

func (knight *Knight) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool(getCountOfEmptyPlaces(board))
	dimension := getDimensionFromBoard(board)

	for i := 0; i < len(board) && len(board) == ((dimension+1)*dimension); i++ {
		if board[i] == emptyField {
			// Get from pool
			outPtr := getBoardFromPool(dimension)
			out := *outPtr
			copy(out, board)

			if !isAnotherFigurePresentBelow(out, i, dimension) && !isAnotherFigurePresentAbove(out, i, dimension) {
				placeAttackPlacesBelow(out, i, dimension)
				placeAttackPlacesAbove(out, i, dimension)
				out[i] = knight.GetName()

				// Make permanent copy for the map
				permanent := make([]byte, len(out))
				copy(permanent, out)
				boards[GenerateHash(permanent)] = permanent
			}

			// Always return to pool
			boardPool.Put(outPtr)
		}
	}
	return boards
}

func placeAttackPlacesBelow(out []byte, position int, dimension int) {
	if position >= len(out) || position == dimension || position%(dimension+1) == dimension {
		return
	}
	currentLine := position/(dimension+1) + 1
	lineBelow := currentLine + 1

	positionBelowLineLeft := position + dimension + 1 - 2
	belowLineLeft := positionBelowLineLeft/(dimension+1) + 1
	if lineBelow == belowLineLeft && positionBelowLineLeft < len(out) && out[positionBelowLineLeft] == emptyField {
		out[positionBelowLineLeft] = attackPlace
	}
	positionBelowLineRight := position + dimension + 1 + 2
	belowLineRight := positionBelowLineRight/(dimension+1) + 1
	if lineBelow == belowLineRight && positionBelowLineRight < len(out) && out[positionBelowLineRight] == emptyField {
		out[positionBelowLineRight] = attackPlace
	}

	line2Below := currentLine + 1 + 1

	positionBelow2LinesLeft := position + 2*(dimension+1) - 1
	below2LineLeft := positionBelow2LinesLeft/(dimension+1) + 1
	if line2Below == below2LineLeft && positionBelow2LinesLeft < len(out) && out[positionBelow2LinesLeft] == emptyField {
		out[positionBelow2LinesLeft] = attackPlace
	}

	positionBelow2LinesRight := position + 2*(dimension+1) + 1
	below2LineRight := positionBelow2LinesRight/(dimension+1) + 1
	if line2Below == below2LineRight && positionBelow2LinesRight < len(out) && out[positionBelow2LinesRight] == emptyField {
		out[positionBelow2LinesRight] = attackPlace
	}
}

func placeAttackPlacesAbove(out []byte, position int, dimension int) {
	if position >= len(out) || position == dimension || position%(dimension+1) == dimension {
		return
	}
	currentLine := position/(dimension+1) + 1

	positionAboveLineLeft := position - dimension - 1 - 2
	positionAboveLineRight := position - dimension - 1 + 2

	lineAbove := currentLine - 1
	aboveLineLeft := positionAboveLineLeft/(dimension+1) + 1
	aboveLineRight := positionAboveLineRight/(dimension+1) + 1

	if aboveLineLeft == lineAbove && positionAboveLineLeft >= 0 && out[positionAboveLineLeft] == emptyField {
		out[positionAboveLineLeft] = attackPlace
	}
	if aboveLineRight == lineAbove && positionAboveLineRight >= 0 && out[positionAboveLineRight] == emptyField {
		out[positionAboveLineRight] = attackPlace
	}

	positionAbove2LinesLeft := position - 2*(dimension+1) - 1
	positionAbove2LinesRight := position - 2*(dimension+1) + 1

	above2LineLeft := positionAbove2LinesLeft/(dimension+1) + 1
	above2LineRight := positionAbove2LinesRight/(dimension+1) + 1
	line2Above := currentLine - 1 - 1
	if above2LineLeft == line2Above && positionAbove2LinesLeft >= 0 && out[positionAbove2LinesLeft] != '\n' {
		out[positionAbove2LinesLeft] = attackPlace
	}
	if above2LineRight == line2Above && positionAbove2LinesRight >= 0 && out[positionAbove2LinesRight] != '\n' {
		out[positionAbove2LinesRight] = attackPlace
	}
}

func isAnotherFigurePresentBelow(out []byte, position int, dimension int) bool {
	currentLine := position/(dimension+1) + 1
	var numbersToCheck []int

	positionBelowLineLeft := position + dimension + 1 - 2
	positionBelowLineRight := position + dimension + 1 + 2

	lineBelow := currentLine + 1
	belowLineLeft := positionBelowLineLeft/(dimension+1) + 1
	belowLineRight := positionBelowLineRight/(dimension+1) + 1

	if lineBelow == belowLineLeft && positionBelowLineLeft < len(out) && out[positionBelowLineLeft] != '\n' {
		numbersToCheck = append(numbersToCheck, positionBelowLineLeft)
	}
	if lineBelow == belowLineRight && positionBelowLineRight < len(out) && out[positionBelowLineRight] != '\n' {
		numbersToCheck = append(numbersToCheck, positionBelowLineRight)
	}

	positionBelow2LinesLeft := position + 2*(dimension+1) - 1
	positionBelow2LinesRight := position + 2*(dimension+1) + 1

	below2LineLeft := positionBelow2LinesLeft/(dimension+1) + 1
	below2LineRight := positionBelow2LinesRight/(dimension+1) + 1
	line2Below := currentLine + 1 + 1
	if line2Below == below2LineRight && positionBelow2LinesRight < len(out) && out[positionBelow2LinesRight] != '\n' {
		numbersToCheck = append(numbersToCheck, positionBelow2LinesRight)
	}
	if line2Below == below2LineLeft && positionBelow2LinesLeft < len(out) && out[positionBelow2LinesLeft] != '\n' {
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

	currentLine := position/(dimension+1) + 1
	var numbersToCheck []int

	positionAboveLineLeft := position - dimension - 1 - 2
	positionAboveLineRight := position - dimension - 1 + 2

	lineAbove := currentLine - 1
	aboveLineLeft := positionAboveLineLeft/(dimension+1) + 1
	aboveLineRight := positionAboveLineRight/(dimension+1) + 1

	if aboveLineLeft == lineAbove && positionAboveLineLeft >= 0 && out[positionAboveLineLeft] != '\n' {
		numbersToCheck = append(numbersToCheck, positionAboveLineLeft)
	}
	if aboveLineRight == lineAbove && positionAboveLineRight >= 0 && out[positionAboveLineRight] != '\n' {
		numbersToCheck = append(numbersToCheck, positionAboveLineRight)
	}

	positionAbove2LinesLeft := position - 2*(dimension+1) - 1
	positionAbove2LinesRight := position - 2*(dimension+1) + 1

	above2LineLeft := positionAbove2LinesLeft/(dimension+1) + 1
	above2LineRight := positionAbove2LinesRight/(dimension+1) + 1
	line2Above := currentLine - 1 - 1
	if above2LineLeft == line2Above && positionAbove2LinesLeft >= 0 && out[positionAbove2LinesLeft] != '\n' {
		numbersToCheck = append(numbersToCheck, positionAbove2LinesLeft)
	}
	if above2LineRight == line2Above && positionAbove2LinesRight >= 0 && out[positionAbove2LinesRight] != '\n' {
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
