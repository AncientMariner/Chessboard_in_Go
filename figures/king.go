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

			// place diagonally

			hashSetOfBoards.Insert(&FigurePosition{string(out), i})
		}
	}

	// if king.next != nil {
	// 	nextFigureBoards := king.next.Handle(board)
	// 	result := set.NewHashSet[*FigurePosition, string](nextFigureBoards.Size())
	//
	// 	nextFigureBoards.ForEach(func(position *FigurePosition) bool {
	// 		result.Insert(position)
	// 		return true
	// 	})
	// 	return result
	// }
	return hashSetOfBoards
}

func placeAttackPlacesVertically(out []rune, position int) {
	// place above
	if position >= defaultDimension+1 && out[position-defaultDimension-1] == '_' {
		out[position-defaultDimension-1] = 'x'
	}
	// place below(the one before last line is the last one where next line is accessible)
	if position < len(out)-defaultDimension-1 && out[position+defaultDimension+1] == '_' {
		out[position+defaultDimension+1] = 'x'
	}
}

func placeAttackPlacesHorizontally(out []rune, position int) {
	// place left
	if position-1 >= 0 && out[position-1] == '_' {
		out[position-1] = 'x'
	}
	// place right
	if (position+1) < len(out) && out[position+1] == '_' {
		out[position+1] = 'x'
	}
}

func (*King) GetName() rune {
	return 'k'
}
