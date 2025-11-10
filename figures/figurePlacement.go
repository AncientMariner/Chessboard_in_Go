package figures

import (
	"crypto/sha512"
	"fmt"
	"strings"
)

type Placement struct {
}

func GenerateHash(s string) string {
	algorithm := sha512.New512_256()
	algorithm.Write([]byte(s))
	return fmt.Sprintf("%x", algorithm.Sum(nil))
}

var defaultDimension = 8

const emptyField = '_'
const attackPlace = 'x'

func drawEmptyBoard() string {

	var board strings.Builder

	for x := 0; x < defaultDimension; x++ {
		for y := 0; y < defaultDimension; y++ {
			board.WriteString(string(emptyField))
		}
		board.WriteString("\n")
	}
	return board.String()
}

func (p *Placement) SetDimension(value int) {
	defaultDimension = value
}

func (p *Placement) PlaceFigures(numberOfFigures int, behaviour FigureBehaviour, boards map[string]string) map[string]string {
	if len(boards) == 0 {
		boards = p.placeFigureOnBoard(drawEmptyBoard(), behaviour)
	}
	for i := 0; i < numberOfFigures; i++ {
		boards = p.placeFigure(boards, behaviour)
	}
	return boards
}

func (p *Placement) placeFigure(boards map[string]string, behaviour FigureBehaviour) map[string]string {
	var resultingMap = make(map[string]string)

	for _, board := range boards {

		boardsWithPlacement := p.placeFigureOnBoard(board, behaviour)

		for hash, element := range boardsWithPlacement {
			resultingMap[hash] = element
		}
	}
	return resultingMap
}

func (p *Placement) placeFigureOnBoard(board string, behaviour FigureBehaviour) map[string]string {
	return behaviour.Handle(board)
}
