package figures

type Queen struct {
	Figure
}

func (queen *Queen) Handle(board string) map[uint32]string {
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

			if !isAnotherFigurePresentOnTheLine(out, i) && !isAnotherFigurePresentOnTheColumn(out, i) && !isAnotherFigurePresentDiag(out, i) {
				placeAttackPlacesHorizontally(out, i)
				placeAttackPlacesVertically(out, i)
				placeAttackPlacesDiagonallyAbove(out, i)
				placeAttackPlacesDiagonallyBelow(out, i)

				out[i] = queen.GetName()

				b := &BoardWithFigurePosition{}
				b.Board = string(out)
				b.number = i
				b.Hash()

				hashSetOfBoards[b.hash] = b.Board
			}
		}
	}
	return hashSetOfBoards
}

func (*Queen) GetName() rune {
	return 'q'
}
