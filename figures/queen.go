package figures

var _ FigureBehaviour = (*Queen)(nil)

type Queen struct {
	Figure
}

func (queen *Queen) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool(getCountOfEmptyPlaces(board))

	for i := 0; i < len(board) && len(board) == ((defaultDimension+1)*defaultDimension); i++ {
		if board[i] == emptyField {
			// Get from pool
			outPtr := boardPool.Get().(*[]byte)
			out := *outPtr
			copy(out, board)

			if !isAnotherFigurePresentOnTheLine(out, i) && !isAnotherFigurePresentOnTheColumn(out, i) && !isAnotherFigurePresentDiag(out, i) {
				placeAttackPlacesHorizontally(out, i)
				placeAttackPlacesVertically(out, i)
				placeAttackPlacesDiagonallyAbove(out, i)
				placeAttackPlacesDiagonallyBelow(out, i)

				out[i] = queen.GetName()

				// Make permanent copy for the map
				permanent := make([]byte, len(out))
				copy(permanent, out)
				boards[GenerateHash(permanent)] = permanent
			}

			// Always return to pool
			boardPool.Put(outPtr)
		}
	}
	return boards
}

func (*Queen) GetName() byte {
	return 'q'
}
