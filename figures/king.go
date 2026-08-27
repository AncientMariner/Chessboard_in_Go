package figures

var _ FigureBehaviour = (*King)(nil)

type King struct {
	Figure
}

func (king *King) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool()
	dimension := getDimensionFromBoard(board)

	for i := 0; i < len(board) && len(board) == (dimension*dimension); i++ {
		if board[i] == emptyField {
			// Check validity first before doing any allocation
			if !isAnotherFigurePresent(board, i, dimension) {
				// Get working buffer from pool
				outPtr := getBoardFromPool(dimension)
				out := *outPtr
				copy(out, board)

				king.placeAttackPlacesHorizontally(out, i, dimension)
				king.placeAttackPlacesVertically(out, i, dimension)
				king.placeDiagonallyAbove(out, i, dimension)
				king.placeDiagonallyBelow(out, i, dimension)
				out[i] = king.GetName()

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

func isAnotherFigurePresent(out []byte, position int, dimension int) bool {
	isPiece := func(i int) bool { return out[i] != emptyField && out[i] != attackPlace }

	prevLine := position >= dimension
	nextLine := position < len(out)-dimension
	leftCol  := position%dimension != 0
	rightCol := position%dimension != dimension-1

	if prevLine {
		above := position - dimension
		if isPiece(above) { return true }
		if leftCol  && isPiece(above-1) { return true }
		if rightCol && isPiece(above+1) { return true }
	}
	if nextLine {
		below := position + dimension
		if isPiece(below) { return true }
		if leftCol  && isPiece(below-1) { return true }
		if rightCol && isPiece(below+1) { return true }
	}
	if leftCol  && isPiece(position-1) { return true }
	if rightCol && isPiece(position+1) { return true }
	return false
}

func (king *King) placeDiagonallyAbove(out []byte, position int, dimension int) {
	positionOneLineAbove := position - dimension

	diagAboveRight := positionOneLineAbove + 1
	previousLineExists := position >= dimension
	rightColumnExists := position%dimension != dimension-1
	if previousLineExists && rightColumnExists && out[diagAboveRight] == emptyField {
		out[diagAboveRight] = attackPlace
	}
	diagAboveLeft := positionOneLineAbove - 1
	leftColumnExists := position%dimension != 0
	if previousLineExists && leftColumnExists && diagAboveLeft >= 0 && out[diagAboveLeft] == emptyField {
		out[diagAboveLeft] = attackPlace
	}
}

func (king *King) placeDiagonallyBelow(out []byte, position int, dimension int) {
	diagBelowRight := position + dimension + 1
	diagBelowLeft := position + dimension - 1
	isNotLastLine := position < len(out)-dimension
	rightColumnExists := position%dimension != dimension-1

	if isNotLastLine && rightColumnExists && diagBelowRight < len(out) && out[diagBelowRight] == emptyField {
		out[diagBelowRight] = attackPlace
	}
	if isNotLastLine && position%dimension != 0 && diagBelowLeft < len(out) && out[diagBelowLeft] == emptyField {
		out[diagBelowLeft] = attackPlace
	}
}

func (king *King) placeAttackPlacesVertically(out []byte, position int, dimension int) {
	positionAbove := position - dimension
	if position >= dimension && out[positionAbove] == emptyField {
		out[positionAbove] = attackPlace
	}
	positionBelow := position + dimension
	if position < len(out)-dimension && out[positionBelow] == emptyField {
		out[positionBelow] = attackPlace
	}
}

func (king *King) placeAttackPlacesHorizontally(out []byte, position int, dimension int) {
	previousPosition := position - 1
	leftColumnExists := position%dimension != 0
	if previousPosition >= 0 && leftColumnExists && out[previousPosition] == emptyField {
		out[previousPosition] = attackPlace
	}
	nextPosition := position + 1
	rightColumnExists := position%dimension != dimension-1
	if nextPosition < len(out) && rightColumnExists && out[nextPosition] == emptyField {
		out[nextPosition] = attackPlace
	}
}

func (*King) GetName() byte {
	return 'k'
}
