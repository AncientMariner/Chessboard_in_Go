package figures

var _ FigureBehaviour = (*Queen)(nil)

type Queen struct {
	Figure
}

func (queen *Queen) Handle(board []byte) map[string][]byte{
	boards := make(map[string][]byte, getCountOfEmptyPlaces(board))

	for i := 0; i < len(board) && len(board) == ((defaultDimension+1)*defaultDimension); i++ {
		if board[i] == emptyField {
			out := make([]byte, len(board))
            copy(out, board) 

			if !isAnotherFigurePresentOnTheLine(out, i) && !isAnotherFigurePresentOnTheColumn(out, i) && !isAnotherFigurePresentDiag(out, i) {
				placeAttackPlacesHorizontally(out, i)
				placeAttackPlacesVertically(out, i)
				placeAttackPlacesDiagonallyAbove(out, i)
				placeAttackPlacesDiagonallyBelow(out, i)

				out[i] = queen.GetName()

				boards[GenerateHash(out)] = out
			}
		}
	}
	return boards
}

func (*Queen) GetName() byte {
	return 'q'
}
