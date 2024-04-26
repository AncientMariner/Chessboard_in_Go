package figures

type Bishop struct {
	Figure
}

func (bishop *Bishop) Handle(board string) map[uint32]string {
	countOfEmptyPlaces := 0
	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			countOfEmptyPlaces++
		}
	}

	hashSetOfBoards := make(map[uint32]string, countOfEmptyPlaces)

	for i := 0; i < len(board) && len(board) == ((defaultDimension+1)*defaultDimension); i++ {
		if board[i] == emptyField {
			out := []rune(board)

			if !isAnotherFigurePresentDiag(out, i) {
				placeAttackPlacesDiagonallyAbove(out, i)
				placeAttackPlacesDiagonallyBelow(out, i)
				out[i] = bishop.GetName()

				item := &FigurePosition{}
				item.Board = string(out)
				item.number = i
				item.Hash()

				hashSetOfBoards[item.hash] = item.Board
			}
		}
	}
	return hashSetOfBoards
}

func placeAttackPlacesDiagonallyBelow(out []rune, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	diagBelowRight := position + defaultDimension + 1 + 1
	diagBelowLeft := position + defaultDimension + 1 - 1

	currentLine := position/defaultDimension + 1
	for lineBelow := currentLine + 1; lineBelow <= defaultDimension; lineBelow++ {
		lineOfTheDiagBelowRight := diagBelowRight/(defaultDimension+1) + 1
		lineOfTheDiagBelowLeft := diagBelowLeft/(defaultDimension+1) + 1

		if lineBelow == lineOfTheDiagBelowRight && diagBelowRight < len(out) && position < len(out)-defaultDimension-1 && out[diagBelowRight] == emptyField {
			out[diagBelowRight] = attackPlace
		}
		diagBelowRight = diagBelowRight + defaultDimension + 1 + 1

		if lineBelow == lineOfTheDiagBelowLeft && diagBelowLeft < len(out) && position < len(out)-defaultDimension-1 && out[diagBelowLeft] == emptyField {
			out[diagBelowLeft] = attackPlace
		}
		diagBelowLeft = diagBelowLeft + defaultDimension + 1 - 1
	}
}

func placeAttackPlacesDiagonallyAbove(out []rune, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	diagAboveRight := position - defaultDimension - 1 + 1
	diagAboveLeft := position - defaultDimension - 1 - 1

	for linesAbove := position / (defaultDimension + 1); linesAbove > 0; linesAbove-- {
		if position >= defaultDimension+1 && diagAboveRight >= 0 && out[diagAboveRight] == emptyField {
			out[diagAboveRight] = attackPlace
		}
		diagAboveRight = diagAboveRight - defaultDimension - 1 + 1
		if position >= defaultDimension+1 && diagAboveLeft >= 0 && out[diagAboveLeft] == emptyField {
			out[diagAboveLeft] = attackPlace
		}
		diagAboveLeft = diagAboveLeft - defaultDimension - 1 - 1
	}
}

func isAnotherFigurePresentDiag(out []rune, position int) bool {
	numberOfLines := len(out) / (defaultDimension + 1)
	currentLine := position / (defaultDimension + 1)

	var diagNumbers []int

	previousLinePositionLeft := position - defaultDimension - 1 - 1
	previousLinePositionRight := position - defaultDimension - 1 + 1
	for counterAboveLines := currentLine; previousLinePositionLeft >= 0 && previousLinePositionRight >= 0 && counterAboveLines > 0; counterAboveLines-- {
		if out[previousLinePositionRight] != '\n' {
			diagNumbers = append(diagNumbers, previousLinePositionRight)
		}
		previousLinePositionRight = previousLinePositionRight - defaultDimension - 1 + 1
		if out[previousLinePositionLeft] != '\n' {
			diagNumbers = append(diagNumbers, previousLinePositionLeft)
		}
		previousLinePositionLeft = previousLinePositionLeft - defaultDimension - 1 - 1
	}

	nextLinePositionLeft := position + defaultDimension + 1 - 1
	nextLinePositionRight := position + defaultDimension + 1 + 1
	for counterBelowLines := numberOfLines - currentLine - 1; nextLinePositionRight < len(out) && nextLinePositionLeft < len(out) && counterBelowLines > 0; counterBelowLines-- {
		if out[nextLinePositionRight] != '\n' {
			diagNumbers = append(diagNumbers, nextLinePositionRight)
		}
		nextLinePositionRight = nextLinePositionRight + defaultDimension + 1 + 1
		if out[nextLinePositionLeft] != '\n' {
			diagNumbers = append(diagNumbers, nextLinePositionLeft)
		}
		nextLinePositionLeft = nextLinePositionLeft + defaultDimension + 1 - 1
	}

	for _, number := range diagNumbers {
		if out[number] != emptyField && out[number] != attackPlace && out[number] != '\n' {
			return true
		}
	}
	return false
}

func (*Bishop) GetName() rune {
	return 'b'
}
