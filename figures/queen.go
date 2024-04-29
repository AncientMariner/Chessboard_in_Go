package figures

type Queen struct {
	Figure
}

func (queen *Queen) Handle(board string) map[string]string {
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

			if isAnotherFigureNotPresent(out, i) {
				placeAttackPlacesHorizontally(out, i)
				placeAttackPlacesVertically(out, i)
				placeAttackPlacesDiagonallyAbove(out, i)
				placeAttackPlacesDiagonallyBelow(out, i)

				out[i] = queen.GetName()

				b := &BoardWithFigurePosition{}
				b.Board = string(out)
				// b.number = i
				// b.Hash()

				boards[b.Board] = b.Board
			}
		}
	}
	return boards
}

func isAnotherFigureNotPresent(out []rune, i int) bool {
	return !isAnotherFigurePresentOnTheLine(out, i) && !isAnotherFigurePresentOnTheColumn(out, i) && !isAnotherFigurePresentDiag(out, i)
}

func (*Queen) GetName() rune {
	return 'q'
}
