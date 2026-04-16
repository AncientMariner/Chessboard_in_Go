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
			// Check validity first before doing any allocation
			if !isAnotherFigurePresentOnTheLine(board, i, dimension) &&
				!isAnotherFigurePresentOnTheColumn(board, i, dimension) &&
				!isAnotherFigurePresentDiag(board, i, dimension) {
				// Get working buffer from pool
				outPtr := getBoardFromPool(dimension)
				out := *outPtr
				copy(out, board)

				placeAttackPlacesHorizontally(out, i, dimension)
				placeAttackPlacesVertically(out, i, dimension)
				placeAttackPlacesDiagonallyAbove(out, i, dimension)
				placeAttackPlacesDiagonallyBelow(out, i, dimension)
				out[i] = queen.GetName()

				// Make permanent copy for storage
				permanent := make([]byte, len(out))
				copy(permanent, out)
				boards[GenerateHash(permanent)] = permanent

				// Return working buffer to pool
				boardPool.Put(outPtr)
			}
		}
	}
	return boards
}

func (*Queen) GetName() byte {
	return 'q'
}
