package figures

var _ FigureBehaviour = (*Queen)(nil)

type Queen struct {
	Figure
}

func (queen *Queen) Handle(board string) map[string]string {
	boards := make(map[string]string, getCountOfEmptyPlaces(board))

	for i := 0; i < len(board) && len(board) == ((defaultDimension+1)*defaultDimension); i++ {
		if board[i] == emptyField {
			out := []rune(board)

			if !isAnotherFigurePresentOnTheLine(out, i) && !isAnotherFigurePresentOnTheColumn(out, i) && !isAnotherFigurePresentDiag(out, i) {
				placeAttackPlacesHorizontally(out, i)
				placeAttackPlacesVertically(out, i)
				placeAttackPlacesDiagonallyAbove(out, i)
				placeAttackPlacesDiagonallyBelow(out, i)

				out[i] = queen.GetName()

				boards[Hash(string(out))] = string(out)
			}
		}
	}
	return boards
}

func (*Queen) GetName() rune {
	return 'q'
}
