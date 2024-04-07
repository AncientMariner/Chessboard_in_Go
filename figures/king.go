package figures

import (
	"github.com/hashicorp/go-set/v2"
)

type King struct {
	Figure
}

func (king *King) Handle(board string) *set.HashSet[*FigurePosition, string] {

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

			out[i] = king.GetName()

			placeAttackPlacesHorizontally(out, i)
			placeAttackPlacesVertically(out, i)
			placeDiagonallyAbove(out, i)
			placeDiagonallyBelow(out, i)

			hashSetOfBoards.Insert(&FigurePosition{string(out), i})
		}
	}
	return hashSetOfBoards
}

func placeDiagonallyAbove(out []rune, position int) {
	if position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	positionOneLineAbove := position - defaultDimension - 1

	diagAboveRight := positionOneLineAbove + 1
	previousLineExists := position >= defaultDimension+1

	if previousLineExists && out[diagAboveRight] == emptyField {
		out[diagAboveRight] = attackPlace
	}
	diagAboveLeft := positionOneLineAbove - 1
	if previousLineExists && (position-1)%defaultDimension != 0 && diagAboveLeft >= 0 && out[diagAboveLeft] == emptyField {
		out[diagAboveLeft] = attackPlace
	}
}

func placeDiagonallyBelow(out []rune, position int) {
	if position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	diagBelowRight := position + defaultDimension + 1 + 1
	diagBelowLeft := position + defaultDimension + 1 - 1
	isOneBeforeLastLine := position < len(out)-defaultDimension-1

	if isOneBeforeLastLine && diagBelowRight < len(out) && out[diagBelowRight] == emptyField {
		out[diagBelowRight] = attackPlace
	}
	if isOneBeforeLastLine && position%defaultDimension != 0 && diagBelowLeft < len(out) && out[diagBelowLeft] == emptyField {
		out[diagBelowLeft] = attackPlace
	}
}

func placeAttackPlacesVertically(out []rune, position int) {
	positionAbove := position - defaultDimension - 1
	if position >= defaultDimension+1 && out[positionAbove] == emptyField {
		out[positionAbove] = attackPlace
	}
	positionBelow := position + defaultDimension + 1
	if position < len(out)-defaultDimension-1 && out[positionBelow] == emptyField {
		out[positionBelow] = attackPlace
	}
}

func placeAttackPlacesHorizontally(out []rune, position int) {
	if position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	previousPosition := position - 1
	if previousPosition >= 0 && out[previousPosition] == emptyField {
		out[previousPosition] = attackPlace
	}
	nextPosition := position + 1
	if nextPosition < len(out) && out[nextPosition] == emptyField {
		out[nextPosition] = attackPlace
	}
}

func (*King) GetName() rune {
	return 'k'
}
