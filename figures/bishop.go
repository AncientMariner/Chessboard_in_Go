package figures

var _ FigureBehaviour = (*Bishop)(nil)

type Bishop struct {
	Figure
}

func (bishop *Bishop) Handle(board string) map[string]string {
	boards := make(map[string]string, getCountOfEmptyPlaces(board))

	for i := 0; i < len(board) && len(board) == ((defaultDimension+1)*defaultDimension); i++ {
		if board[i] == emptyField {
			out := []rune(board)

			if !isAnotherFigurePresentDiag(out, i) {
				placeAttackPlacesDiagonallyAbove(out, i)
				placeAttackPlacesDiagonallyBelow(out, i)
				out[i] = bishop.GetName()
				outAsString := string(out)

				boards[GenerateHash(outAsString)] = outAsString
			}
		}
	}
	return boards
}

func placeAttackPlacesDiagonallyBelow(out []rune, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	diagBelowRight := position + defaultDimension + 1 + 1
	diagBelowLeft := position + defaultDimension + 1 - 1

	currentLine := position/(defaultDimension+1) + 1
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
	diagAboveLeft := position - defaultDimension - 1 - 1
	diagAboveRight := position - defaultDimension - 1 + 1

	currentLine := position/(defaultDimension+1) + 1
	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		if position >= defaultDimension+1 && diagAboveLeft >= 0 && out[diagAboveLeft] == emptyField {
			out[diagAboveLeft] = attackPlace
		}
		diagAboveLeft = diagAboveLeft - defaultDimension - 1 - 1

		if position >= defaultDimension+1 && diagAboveRight >= 0 && out[diagAboveRight] == emptyField {
			out[diagAboveRight] = attackPlace
		}
		diagAboveRight = diagAboveRight - defaultDimension - 1 + 1
	}
}

func isAnotherFigurePresentDiag(out []rune, position int) bool {
	currentLine := position/(defaultDimension+1) + 1
	var diagNumbers []int

	previousLinePositionLeft := position - defaultDimension - 1 - 1
	previousLinePositionRight := position - defaultDimension - 1 + 1

	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		lineOfTheDiagAboveLeft := previousLinePositionLeft/(defaultDimension+1) + 1
		lineOfTheDiagAboveRight := previousLinePositionRight/(defaultDimension+1) + 1

		if lineOfTheDiagAboveLeft == lineAbove && previousLinePositionLeft >= 0 && out[previousLinePositionLeft] != '\n' {
			diagNumbers = append(diagNumbers, previousLinePositionLeft)
		}
		previousLinePositionLeft = previousLinePositionLeft - defaultDimension - 1 - 1
		if lineOfTheDiagAboveRight == lineAbove && previousLinePositionRight >= 0 && out[previousLinePositionRight] != '\n' {
			diagNumbers = append(diagNumbers, previousLinePositionRight)
		}
		previousLinePositionRight = previousLinePositionRight - defaultDimension - 1 + 1
	}

	nextLinePositionLeft := position + defaultDimension + 1 - 1
	nextLinePositionRight := position + defaultDimension + 1 + 1

	for lineBelow := currentLine + 1; lineBelow <= defaultDimension; lineBelow++ {
		lineOfTheDiagBelowLeft := nextLinePositionLeft/(defaultDimension+1) + 1
		lineOfTheDiagBelowRight := nextLinePositionRight/(defaultDimension+1) + 1

		if lineBelow == lineOfTheDiagBelowRight && nextLinePositionRight < len(out) && out[nextLinePositionRight] != '\n' {
			diagNumbers = append(diagNumbers, nextLinePositionRight)
		}
		nextLinePositionRight = nextLinePositionRight + defaultDimension + 1 + 1
		if lineBelow == lineOfTheDiagBelowLeft && nextLinePositionLeft < len(out) && out[nextLinePositionLeft] != '\n' {
			diagNumbers = append(diagNumbers, nextLinePositionLeft)
		}
		nextLinePositionLeft = nextLinePositionLeft + defaultDimension + 1 - 1
	}

	for _, number := range diagNumbers {
		if out[number] != emptyField && out[number] != attackPlace {
			return true
		}
	}
	return false
}

func (*Bishop) GetName() rune {
	return 'b'
}
