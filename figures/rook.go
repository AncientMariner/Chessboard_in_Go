package figures

var _ FigureBehaviour = (*Rook)(nil)

type Rook struct {
	Figure
}

func (rook *Rook) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool()
	dimension := getDimensionFromBoard(board)
	boardHash := ZobristHash(board)

	for i := 0; i < len(board) && len(board) == (dimension*dimension); i++ {
		if board[i] == emptyField {
			// Check validity first before doing any allocation
			if !isAnotherFigurePresentOnTheLine(board, i, dimension) && !isAnotherFigurePresentOnTheColumn(board, i, dimension) {
				// Get working buffer from pool
				outPtr := getBoardFromPool(dimension)
				out := *outPtr
				copy(out, board)

				h := boardHash
				h = placeAttackPlacesHorizontally(out, i, dimension, h)
				h = placeAttackPlacesVertically(out, i, dimension, h)
				h = ZobristUpdate(h, i, emptyField, rook.GetName())
				out[i] = rook.GetName()

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

func placeAttackPlacesHorizontally(out []byte, position int, dimension int, h uint64) uint64 {
	if position >= len(out) {
		return h
	}

	var counterOfLeftPositions = (position) % dimension
	var counterOfRightPositions = dimension - ((position) % dimension) - 1

	for previousPosition := position - 1; counterOfLeftPositions > 0 && previousPosition >= 0; counterOfLeftPositions-- {
		if out[previousPosition] == emptyField {
			out[previousPosition] = attackPlace
			h = ZobristUpdate(h, previousPosition, emptyField, attackPlace)
		}
		previousPosition--
	}

	for nextPosition := position + 1; counterOfRightPositions > 0 && nextPosition < len(out); counterOfRightPositions-- {
		if out[nextPosition] == emptyField {
			out[nextPosition] = attackPlace
			h = ZobristUpdate(h, nextPosition, emptyField, attackPlace)
		}
		nextPosition++
	}
	return h
}

func isAnotherFigurePresentOnTheLine(out []byte, position int, dimension int) bool {
	left := position % dimension
	right := dimension - left - 1

	for p := position - 1; left > 0; left-- {
		if out[p] != emptyField && out[p] != attackPlace {
			return true
		}
		p--
	}
	for p := position + 1; right > 0; right-- {
		if out[p] != emptyField && out[p] != attackPlace {
			return true
		}
		p++
	}
	return false
}

func placeAttackPlacesVertically(out []byte, position int, dimension int, h uint64) uint64 {
	if position >= len(out) {
		return h
	}

	abovePosition := position - dimension
	currentLine := position/dimension + 1
	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		lineOfTheAbovePosition := abovePosition/dimension + 1
		if lineOfTheAbovePosition == lineAbove && out[abovePosition] == emptyField {
			out[abovePosition] = attackPlace
			h = ZobristUpdate(h, abovePosition, emptyField, attackPlace)
		}
		abovePosition = abovePosition - dimension
	}

	belowPosition := position + dimension
	for lineBelow := currentLine + 1; lineBelow <= dimension; lineBelow++ {
		lineOfTheBelowPosition := belowPosition/dimension + 1

		if lineOfTheBelowPosition == lineBelow && out[belowPosition] == emptyField {
			out[belowPosition] = attackPlace
			h = ZobristUpdate(h, belowPosition, emptyField, attackPlace)
		}
		belowPosition = belowPosition + dimension
	}
	return h
}

func isAnotherFigurePresentOnTheColumn(out []byte, position int, dimension int) bool {
	currentLine := position/dimension + 1

	abovePosition := position - dimension
	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		if abovePosition/dimension+1 == lineAbove && out[abovePosition] != emptyField && out[abovePosition] != attackPlace {
			return true
		}
		abovePosition -= dimension
	}

	belowPosition := position + dimension
	for lineBelow := currentLine + 1; lineBelow <= dimension; lineBelow++ {
		if belowPosition/dimension+1 == lineBelow && belowPosition < len(out) && out[belowPosition] != emptyField && out[belowPosition] != attackPlace {
			return true
		}
		belowPosition += dimension
	}
	return false
}

func (*Rook) GetName() byte {
	return 'r'
}
