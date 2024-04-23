package figures

import "github.com/hashicorp/go-set/v2"

type Bishop struct {
	Figure
}

func (bishop *Bishop) Handle(board string) *set.HashSet[*FigurePosition, string] {
	countOfEmptyPlaces := 0
	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			countOfEmptyPlaces++
		}
	}

	hashSetOfBoards := set.NewHashSet[*FigurePosition, string](countOfEmptyPlaces)

	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			out := []rune(board)

			out[i] = bishop.GetName()

			bishop.placeAttackPlacesDiagonallyAbove(out, i)
			bishop.placeAttackPlacesDiagonallyBelow(out, i)

			hashSetOfBoards.Insert(&FigurePosition{string(out), i})
		}
	}

	return hashSetOfBoards

}

func (bishop *Bishop) placeAttackPlacesDiagonallyBelow(out []rune, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	diagBelowRight := position + defaultDimension + 1 + 1
	diagBelowLeft := position + defaultDimension + 1 - 1

	for linesBelow := defaultDimension - position/(defaultDimension+1); linesBelow > 0; linesBelow-- {
		if diagBelowRight < len(out) && position < len(out)-defaultDimension-1 && out[diagBelowRight] == emptyField {
			out[diagBelowRight] = attackPlace
			diagBelowRight = diagBelowRight + defaultDimension + 1 + 1
		}
		if diagBelowLeft < len(out) && position < len(out)-defaultDimension-1 && out[diagBelowLeft] == emptyField {
			out[diagBelowLeft] = attackPlace
			diagBelowLeft = diagBelowLeft + defaultDimension + 1 - 1
		}
	}
}

func (bishop *Bishop) placeAttackPlacesDiagonallyAbove(out []rune, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	diagAboveRight := position - defaultDimension - 1 + 1
	diagAboveLeft := position - defaultDimension - 1 - 1

	for linesAbove := position / (defaultDimension + 1); linesAbove > 0; linesAbove-- {
		if position >= defaultDimension+1 && diagAboveRight >= 0 && out[diagAboveRight] == emptyField {
			out[diagAboveRight] = attackPlace
			diagAboveRight = diagAboveRight - defaultDimension - 1 + 1
		}
		if position >= defaultDimension+1 && diagAboveLeft >= 0 && out[diagAboveLeft] == emptyField {
			out[diagAboveLeft] = attackPlace
			diagAboveLeft = diagAboveLeft - defaultDimension - 1 - 1
		}
	}
}

func isAnotherFigurePresentDiag(out []rune, position int) bool {
	numberOfLines := len(out) / (defaultDimension + 1)
	currentLine := position / (defaultDimension + 1)

	var diagNumbers []int

	// check diag left and right is not in the same line

	previousLinePositionLeft := position - defaultDimension - 1 - 1
	previousLinePositionRight := position - defaultDimension - 1 + 1
	for counterAboveLines := currentLine; previousLinePositionLeft >= 0 && previousLinePositionRight >= 0 && counterAboveLines > 0; counterAboveLines-- {
		// if out[previousLinePositionRight] == emptyField || out[previousLinePositionRight] == attackPlace {
		diagNumbers = append(diagNumbers, previousLinePositionRight)
		previousLinePositionRight = previousLinePositionRight - defaultDimension - 1 + 1
		// }
		// if out[previousLinePositionLeft] == emptyField || out[previousLinePositionLeft] == attackPlace {
		diagNumbers = append(diagNumbers, previousLinePositionLeft)
		previousLinePositionLeft = previousLinePositionLeft - defaultDimension - 1 - 1
		// }
	}

	nextLinePositionLeft := position + defaultDimension + 1 - 1
	nextLinePositionRight := position + defaultDimension + 1 + 1
	for counterBelowLines := numberOfLines - currentLine - 1; nextLinePositionRight < len(out) && nextLinePositionLeft < len(out) && counterBelowLines > 0; counterBelowLines-- {
		// if out[nextLinePositionRight] == emptyField || out[nextLinePositionRight] == attackPlace {
		diagNumbers = append(diagNumbers, nextLinePositionRight)
		nextLinePositionRight = nextLinePositionRight + defaultDimension + 1 + 1
		// }
		// if out[nextLinePositionLeft] == emptyField || out[nextLinePositionLeft] == attackPlace {
		diagNumbers = append(diagNumbers, nextLinePositionLeft)
		nextLinePositionLeft = nextLinePositionLeft + defaultDimension + 1 - 1
		// }
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
