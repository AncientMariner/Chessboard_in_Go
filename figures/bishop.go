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

			bishop.placeAttackPlacesDiagonallyBelow(out, i)

			hashSetOfBoards.Insert(&FigurePosition{string(out), i})
		}
	}

	return hashSetOfBoards

}

func (bishop *Bishop) placeAttackPlacesDiagonallyBelow(out []rune, position int) {

	if position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	diagBelowRight := position + defaultDimension + 1 + 1
	diagBelowLeft := position + defaultDimension + 1 - 1
	isNotLastLine := position < len(out)-defaultDimension-1

	if isNotLastLine && diagBelowRight < len(out) && out[diagBelowRight] == emptyField {
		out[diagBelowRight] = attackPlace
	}
	if isNotLastLine && position%defaultDimension != 0 && diagBelowLeft < len(out) && out[diagBelowLeft] == emptyField {
		out[diagBelowLeft] = attackPlace
	}

}

func (*Bishop) GetName() rune {
	return 'b'
}
