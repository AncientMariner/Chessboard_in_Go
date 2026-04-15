package figures

var _ FigureBehaviour = (*Rook)(nil)

type Rook struct {
	Figure
}

func (rook *Rook) Handle(board []byte) map[uint64][]byte {
	boards := make(map[uint64][]byte, getCountOfEmptyPlaces(board))

	for i := 0; i < len(board) && len(board) == ((defaultDimension+1)*defaultDimension); i++ {
		if board[i] == emptyField {
			// Get from pool
			outPtr := boardPool.Get().(*[]byte)
			out := *outPtr
			copy(out, board)

			if !isAnotherFigurePresentOnTheLine(out, i) && !isAnotherFigurePresentOnTheColumn(out, i) {
				placeAttackPlacesHorizontally(out, i)
				placeAttackPlacesVertically(out, i)
				out[i] = rook.GetName()

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

func placeAttackPlacesHorizontally(out []byte, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}

	var counterOfLeftPositions = (position) % (defaultDimension + 1)
	var counterOfRightPositions = defaultDimension - ((position) % (defaultDimension + 1)) - 1

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

func isAnotherFigurePresentOnTheLine(out []byte, position int) bool {
	var counterOfLeftPositions = (position) % (defaultDimension + 1)
	var counterOfRightPositions = defaultDimension - ((position) % (defaultDimension + 1)) - 1

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
	return len(previousPositionNumbers)+len(nextPositionNumbers) < defaultDimension-1
}

func placeAttackPlacesVertically(out []byte, position int) {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}

	abovePosition := position - defaultDimension - 1
	currentLine := position/(defaultDimension+1) + 1
	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		lineOfTheAbovePosition := abovePosition/(defaultDimension+1) + 1
		if lineOfTheAbovePosition == lineAbove && out[abovePosition] == emptyField {
			out[abovePosition] = attackPlace
		}
		abovePosition = abovePosition - defaultDimension - 1
	}

	belowPosition := position + defaultDimension + 1
	for lineBelow := currentLine + 1; lineBelow <= defaultDimension; lineBelow++ {
		lineOfTheBelowPosition := belowPosition/(defaultDimension+1) + 1

		if lineOfTheBelowPosition == lineBelow && out[belowPosition] == emptyField {
			out[belowPosition] = attackPlace
		}
		belowPosition = belowPosition + defaultDimension + 1
	}
}

func isAnotherFigurePresentOnTheColumn(out []byte, position int) bool {
	currentLine := position/(defaultDimension+1) + 1

	var aboveLineNumbers []int
	abovePosition := position - defaultDimension - 1

	for lineAbove := currentLine - 1; lineAbove > 0; lineAbove-- {
		lineOfTheAbovePosition := abovePosition/(defaultDimension+1) + 1

		if lineOfTheAbovePosition == lineAbove && position >= defaultDimension+1 && out[abovePosition] == emptyField || out[abovePosition] == attackPlace {
			aboveLineNumbers = append(aboveLineNumbers, abovePosition)
		}
		abovePosition = abovePosition - defaultDimension - 1
	}

	var belowLineNumbers []int
	belowPosition := position + defaultDimension + 1

	for lineBelow := currentLine + 1; lineBelow <= defaultDimension; lineBelow++ {
		lineOfTheBelowPosition := belowPosition/(defaultDimension+1) + 1

		if lineBelow == lineOfTheBelowPosition && out[belowPosition] == emptyField || out[belowPosition] == attackPlace {
			belowLineNumbers = append(belowLineNumbers, belowPosition)
		}
		belowPosition = belowPosition + defaultDimension + 1
	}
	return len(aboveLineNumbers)+len(belowLineNumbers) < defaultDimension-1
}

func (*Rook) GetName() byte {
	return 'r'
}
