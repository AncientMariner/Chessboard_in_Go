package figures

var _ FigureBehaviour = (*Queen)(nil)

type Queen struct {
	Figure
}

func (queen *Queen) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool()
	dimension := getDimensionFromBoard(board)
	boardHash := ZobristHash(board)

	for i := 0; i < len(board) && len(board) == (dimension*dimension); i++ {
		if board[i] == emptyField {
			// Check validity first before doing any allocation
			if !isAnotherFigurePresentOnTheLine(board, i, dimension) &&
				!isAnotherFigurePresentOnTheColumn(board, i, dimension) &&
				!isAnotherFigurePresentDiag(board, i, dimension) {
				// Get working buffer from pool
				outPtr := getBoardFromPool(dimension)
				out := *outPtr
				copy(out, board)

				h := boardHash
				h = placeAttackPlacesHorizontally(out, i, dimension, h)
				h = placeAttackPlacesVertically(out, i, dimension, h)
				h = placeAttackPlacesDiagonallyAbove(out, i, dimension, h)
				h = placeAttackPlacesDiagonallyBelow(out, i, dimension, h)
				h = ZobristUpdate(h, i, emptyField, queen.GetName())
				out[i] = queen.GetName()

				// Make permanent copy for storage
				permanent := make([]byte, len(out))
				copy(permanent, out)
				boards[h] = permanent

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
