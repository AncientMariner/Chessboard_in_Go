package figures

var _ FigureBehaviour = (*Knight)(nil)

type Knight struct {
	Figure
}

func (knight *Knight) Handle(board string) map[string]string {

	countOfEmptyPlaces := 0
	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			countOfEmptyPlaces++
		}
	}

	boards := make(map[string]string, countOfEmptyPlaces)

	for i := 0; i < len(board) && len(board) == ((defaultDimension+1)*defaultDimension); i++ {
		if board[i] == emptyField {
			out := []rune(board)

			if !isAnotherFigurePresentBelow(out, i) && !isAnotherFigurePresentBelow(out, i) {
				placeAttackPlacesBelow(out, i)
				placeAttackPlacesAbove(out, i)

				out[i] = knight.GetName()

				b := &BoardWithFigurePosition{}
				b.Board = string(out)
				boards[b.Hash()] = b.Board
			}
		}
	}
	return boards
}

func placeAttackPlacesBelow(out []rune, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	currentLine := position/(defaultDimension+1) + 1
	lineBelow := currentLine + 1

	positionBelowLineLeft := position + defaultDimension + 1 - 2
	belowLineLeft := positionBelowLineLeft/(defaultDimension+1) + 1
	if lineBelow == belowLineLeft && positionBelowLineLeft < len(out) && out[positionBelowLineLeft] == emptyField {
		out[positionBelowLineLeft] = attackPlace
	}
	positionBelowLineRight := position + defaultDimension + 1 + 2
	belowLineRight := positionBelowLineRight/(defaultDimension+1) + 1
	if lineBelow == belowLineRight && positionBelowLineRight < len(out) && out[positionBelowLineRight] == emptyField {
		out[positionBelowLineRight] = attackPlace
	}

	line2Below := currentLine + 1 + 1

	positionBelow2LinesLeft := position + 2*(defaultDimension+1) - 1
	below2LineLeft := positionBelow2LinesLeft/(defaultDimension+1) + 1
	if line2Below == below2LineLeft && positionBelow2LinesLeft < len(out) && out[positionBelow2LinesLeft] == emptyField {
		out[positionBelow2LinesLeft] = attackPlace
	}

	positionBelow2LinesRight := position + 2*(defaultDimension+1) + 1
	below2LineRight := positionBelow2LinesRight/(defaultDimension+1) + 1
	if line2Below == below2LineRight && positionBelow2LinesRight < len(out) && out[positionBelow2LinesRight] == emptyField {
		out[positionBelow2LinesRight] = attackPlace
	}
}

func placeAttackPlacesAbove(out []rune, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	currentLine := position/(defaultDimension+1) + 1

	positionAboveLineLeft := position - defaultDimension - 1 - 2
	positionAboveLineRight := position - defaultDimension - 1 + 2

	lineAbove := currentLine - 1
	aboveLineLeft := positionAboveLineLeft/(defaultDimension+1) + 1
	aboveLineRight := positionAboveLineRight/(defaultDimension+1) + 1

	if aboveLineLeft == lineAbove && positionAboveLineLeft >= 0 && out[positionAboveLineLeft] == emptyField {
		out[positionAboveLineLeft] = attackPlace
	}
	if aboveLineRight == lineAbove && positionAboveLineRight >= 0 && out[positionAboveLineRight] == emptyField {
		out[positionAboveLineRight] = attackPlace
	}

	positionAbove2LinesLeft := position - 2*(defaultDimension+1) - 1
	positionAbove2LinesRight := position - 2*(defaultDimension+1) + 1

	above2LineLeft := positionAbove2LinesLeft/(defaultDimension+1) + 1
	above2LineRight := positionAbove2LinesRight/(defaultDimension+1) + 1
	line2Above := currentLine - 1 - 1
	if above2LineLeft == line2Above && positionAbove2LinesLeft >= 0 && out[positionAbove2LinesLeft] != '\n' {
		out[positionAbove2LinesLeft] = attackPlace
	}
	if above2LineRight == line2Above && positionAbove2LinesRight >= 0 && out[positionAbove2LinesRight] != '\n' {
		out[positionAbove2LinesRight] = attackPlace
	}
}

func isAnotherFigurePresentBelow(out []rune, position int) bool {
	currentLine := position/(defaultDimension+1) + 1
	var numbersToCheck []int

	positionBelowLineLeft := position + defaultDimension + 1 - 2
	positionBelowLineRight := position + defaultDimension + 1 + 2

	lineBelow := currentLine + 1
	belowLineLeft := positionBelowLineLeft/(defaultDimension+1) + 1
	belowLineRight := positionBelowLineRight/(defaultDimension+1) + 1

	if lineBelow == belowLineLeft && positionBelowLineLeft < len(out) && out[positionBelowLineLeft] != '\n' {
		numbersToCheck = append(numbersToCheck, positionBelowLineLeft)
	}
	if lineBelow == belowLineRight && positionBelowLineRight < len(out) && out[positionBelowLineRight] != '\n' {
		numbersToCheck = append(numbersToCheck, positionBelowLineRight)
	}

	positionBelow2LinesLeft := position + 2*(defaultDimension+1) - 1
	positionBelow2LinesRight := position + 2*(defaultDimension+1) + 1

	below2LineLeft := positionBelow2LinesLeft/(defaultDimension+1) + 1
	below2LineRight := positionBelow2LinesRight/(defaultDimension+1) + 1
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

func isAnotherFigurePresentAbove(out []rune, position int) bool {

	currentLine := position/(defaultDimension+1) + 1
	var numbersToCheck []int

	positionAboveLineLeft := position - defaultDimension - 1 - 2
	positionAboveLineRight := position - defaultDimension - 1 + 2

	lineAbove := currentLine - 1
	aboveLineLeft := positionAboveLineLeft/(defaultDimension+1) + 1
	aboveLineRight := positionAboveLineRight/(defaultDimension+1) + 1

	if aboveLineLeft == lineAbove && positionAboveLineLeft >= 0 && out[positionAboveLineLeft] != '\n' {
		numbersToCheck = append(numbersToCheck, positionAboveLineLeft)
	}
	if aboveLineRight == lineAbove && positionAboveLineRight >= 0 && out[positionAboveLineRight] != '\n' {
		numbersToCheck = append(numbersToCheck, positionAboveLineRight)
	}

	positionAbove2LinesLeft := position - 2*(defaultDimension+1) - 1
	positionAbove2LinesRight := position - 2*(defaultDimension+1) + 1

	above2LineLeft := positionAbove2LinesLeft/(defaultDimension+1) + 1
	above2LineRight := positionAbove2LinesRight/(defaultDimension+1) + 1
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

func (*Knight) GetName() rune {
	return 'n'
}
