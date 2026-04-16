package figures

var _ FigureBehaviour = (*Queen)(nil)

type Queen struct {
	Figure
}

func (queen *Queen) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool(getCountOfEmptyPlaces(board))
	dimension := getDimensionFromBoard(board)

	for i := 0; i < len(board) && len(board) == ((dimension+1)*dimension); i++ {
		if board[i] == emptyField {
			// Get from pool
			outPtr := getBoardFromPool(dimension)
			out := *outPtr
			copy(out, board)

			if !isAnotherFigurePresentOnTheLine(out, i, dimension) && !isAnotherFigurePresentOnTheColumn(out, i, dimension) && !isAnotherFigurePresentDiag(out, i, dimension) {
				placeAttackPlacesHorizontally(out, i, dimension)
				placeAttackPlacesVertically(out, i, dimension)
				placeAttackPlacesDiagonallyAbove(out, i, dimension)
				placeAttackPlacesDiagonallyBelow(out, i, dimension)

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
