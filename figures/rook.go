package figures

var _ FigureBehaviour = (*Rook)(nil)

type Rook struct {
	Figure
}

func (rook *Rook) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool(getCountOfEmptyPlaces(board))
	dimension := getDimensionFromBoard(board)

	for i := 0; i < len(board) && len(board) == ((dimension+1)*dimension); i++ {
		if board[i] == emptyField {
			// Check validity first before doing any allocation
			if !isAnotherFigurePresentOnTheLine(board, i, dimension) && !isAnotherFigurePresentOnTheColumn(board, i, dimension) {
				// Get working buffer from pool
				outPtr := getBoardFromPool(dimension)
				out := *outPtr
				copy(out, board)

				placeAttackPlacesHorizontally(out, i, dimension)
				placeAttackPlacesVertically(out, i, dimension)
				out[i] = rook.GetName()

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

func placeAttackPlacesHorizontally(out []byte, position int, dimension int) {
	if position >= len(out) || position == dimension || position%(dimension+1) == dimension {
		return
	}

	var counterOfLeftPositions = (position) % (dimension + 1)
	var counterOfRightPositions = dimension - ((position) % (dimension + 1)) - 1

	for previousPosition := position - 1; counterOfLeftPositions >= 0 && previousPosition >= 0; counterOfLeftPositions-- {
		if out[previousPosition] == emptyField {
			out[previousPosition] = attackPlace
		}
		previousPosition--
	}

	for nextPosition := position + 1; counterOfRightPositions >= 0 && nextPosition < len(out); counterOfRightPositions-- {
		if out[nextPosition] == emptyField {
			out[nextPosition] = attackPlace
		}
		nextPosition++
	}
}

func isAnotherFigurePresentOnTheLine(out []byte, position int, dimension int) bool {
	var counterOfLeftPositions = (position) % (dimension + 1)
	var counterOfRightPositions = dimension - ((position) % (dimension + 1)) - 1

	var previousPositionNumbers []int
	var nextPositionNumbers []int

	for previousPosition := position - 1; counterOfLeftPositions >= 0 && previousPosition >= 0; counterOfLeftPositions-- {
		if out[previousPosition] == emptyField || out[previousPosition] == attackPlace {
			previousPositionNumbers = append(previousPositionNumbers, previousPosition)
		}
		previousPosition--
	}

	for nextPosition := position + 1; counterOfRightPositions >= 0 && nextPosition < len(out); counterOfRightPositions-- {
		if out[nextPosition] == emptyField || out[nextPosition] == attackPlace {
			nextPositionNumbers = append(nextPositionNumbers, nextPosition)
		}
		nextPosition++
	}
	return len(previousPositionNumbers)+len(nextPositionNumbers) < dimension-1
}

func placeAttackPlacesVertically(out []byte, position int, dimension int) {
	if position >= len(out) || position == dimension || position%(dimension+1) == dimension {
		return
	}

	abovePosition := position - dimension - 1
	currentLine := position/(dimension+1) + 1
	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		lineOfTheAbovePosition := abovePosition/(dimension+1) + 1
		if lineOfTheAbovePosition == lineAbove && out[abovePosition] == emptyField {
			out[abovePosition] = attackPlace
		}
		abovePosition = abovePosition - dimension - 1
	}

	belowPosition := position + dimension + 1
	for lineBelow := currentLine + 1; lineBelow <= dimension; lineBelow++ {
		lineOfTheBelowPosition := belowPosition/(dimension+1) + 1

		if lineOfTheBelowPosition == lineBelow && out[belowPosition] == emptyField {
			out[belowPosition] = attackPlace
		}
		belowPosition = belowPosition + dimension + 1
	}
}

func isAnotherFigurePresentOnTheColumn(out []byte, position int, dimension int) bool {
	currentLine := position/(dimension+1) + 1

	var aboveLineNumbers []int
	abovePosition := position - dimension - 1

	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		lineOfTheAbovePosition := abovePosition/(dimension+1) + 1

		if lineOfTheAbovePosition == lineAbove && position >= dimension+1 && out[abovePosition] == emptyField || out[abovePosition] == attackPlace {
			aboveLineNumbers = append(aboveLineNumbers, abovePosition)
		}
		abovePosition = abovePosition - dimension - 1
	}

	var belowLineNumbers []int
	belowPosition := position + dimension + 1

	for lineBelow := currentLine + 1; lineBelow <= dimension; lineBelow++ {
		lineOfTheBelowPosition := belowPosition/(dimension+1) + 1

		if lineBelow == lineOfTheBelowPosition && out[belowPosition] == emptyField || out[belowPosition] == attackPlace {
			belowLineNumbers = append(belowLineNumbers, belowPosition)
		}
		belowPosition = belowPosition + dimension + 1
	}
	return len(aboveLineNumbers)+len(belowLineNumbers) < dimension-1
}

func (*Rook) GetName() byte {
	return 'r'
}
