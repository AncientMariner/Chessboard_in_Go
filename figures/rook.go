package figures

import "github.com/hashicorp/go-set/v2"

type Rook struct {
	Figure
}

func (rook *Rook) Handle(board string) *set.HashSet[*FigurePosition, string] {

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

			out[i] = rook.GetName()

			rook.placeAttackPlacesHorizontally(out, i)
			// placeAttackPlacesVertically(out, i)

			hashSetOfBoards.Insert(&FigurePosition{string(out), i})
		}
	}

	return set.NewHashSet[*FigurePosition, string](0)
}

func (rook *Rook) placeAttackPlacesHorizontally(out []rune, position int) {
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

func (*Rook) GetName() rune {
	return 'r'
}
