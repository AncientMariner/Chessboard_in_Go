package figures

import (
	"github.com/hashicorp/go-set/v2"
)

type Rook struct {
	Figure
}

func (rook *Rook) Handle(board string) *set.HashSet[*FigurePosition, string] {

	countOfEmptyPlaces := 0
	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			countOfEmptyPlaces++
		}
	}

	hashSetOfBoards := set.NewHashSet[*FigurePosition, string](countOfEmptyPlaces)

	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			out := []rune(board)

			isPossibleToPlaceHorizontally := rook.placeAttackPlacesHorizontally(out, i)
			isPossibleToPlaceVertically := rook.placeAttackPlacesVertically(out, i)
			if isPossibleToPlaceHorizontally && isPossibleToPlaceVertically {
				out[i] = rook.GetName()
				hashSetOfBoards.Insert(&FigurePosition{string(out), i})
			}
		}
	}
	return hashSetOfBoards
}

func (rook *Rook) placeAttackPlacesHorizontally(out []rune, position int) bool {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return false
	}

	if isAnotherFigurePresentOnTheLine(out, position) {
		return false
	} else {
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
		return true
	}
}

func isAnotherFigurePresentOnTheLine(out []rune, position int) bool {
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
	return len(previousPositionNumbers)+len(nextPositionNumbers) < 7
}

func (rook *Rook) placeAttackPlacesVertically(out []rune, position int) bool {
	if position >= len(out) || position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return false
	}

	if isAnotherFigurePresentOnTheColumn(out, position) {
		return false
	} else {
		positionAbove := position - defaultDimension - 1

		for linesAbove := position / (defaultDimension + 1); linesAbove > 0; linesAbove-- {
			if position >= defaultDimension+1 && positionAbove >= 0 && out[positionAbove] == emptyField {
				out[positionAbove] = attackPlace
				positionAbove = positionAbove - defaultDimension - 1
			}
		}

		positionBelow := position + defaultDimension + 1

		for linesBelow := defaultDimension - position/(defaultDimension+1); linesBelow > 0; linesBelow-- {
			if positionBelow < len(out) && position < len(out)-defaultDimension-1 && out[positionBelow] == emptyField {
				out[positionBelow] = attackPlace
				positionBelow = positionBelow + defaultDimension + 1
			}
		}
		return true
	}
}

func isAnotherFigurePresentOnTheColumn(out []rune, position int) bool {
	numberOfLines := len(out) / (defaultDimension + 1)
	currentLine := position / (defaultDimension + 1)

	var aboveLineNumbers []int
	counterAboveLines := currentLine
	for previousLinePosition := position - defaultDimension - 1; previousLinePosition >= 0 && counterAboveLines > 0; counterAboveLines-- {
		if out[previousLinePosition] == emptyField || out[previousLinePosition] == attackPlace {
			aboveLineNumbers = append(aboveLineNumbers, previousLinePosition)
		}
		previousLinePosition = previousLinePosition - defaultDimension - 1
	}

	var belowLineNumbers []int
	var counterBelowLines = numberOfLines - currentLine - 1
	for nextLinePosition := position + defaultDimension + 1; nextLinePosition < len(out) && counterBelowLines > 0; counterBelowLines-- {
		if out[nextLinePosition] == emptyField || out[nextLinePosition] == attackPlace {
			belowLineNumbers = append(belowLineNumbers, nextLinePosition)
		}
		nextLinePosition = nextLinePosition + defaultDimension + 1
	}
	return len(aboveLineNumbers)+len(belowLineNumbers) < 7
}

func (*Rook) GetName() rune {
	return 'r'
}
