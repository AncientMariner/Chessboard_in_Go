package figures

import (
	"crypto/sha512"
	"fmt"
	"strings"
)

type Placement struct {
}

type BoardWithFigurePosition struct {
	Board string
	hash  string
}

func (e *BoardWithFigurePosition) Hash() string {
	algorithm := sha512.New512_256()
	algorithm.Write([]byte(e.Board))
	e.hash = fmt.Sprintf("%x", algorithm.Sum(nil))
	return e.hash
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

func (p *Placement) PlaceFigure(numberOfFigures int, behaviour FigureBehaviour, boards map[string]string) map[string]string {
	for i := 0; i < numberOfFigures; i++ {
		if len(boards) == 0 {
			boards = p.PlaceFiguresOnEmptyBoard(drawEmptyBoard(), behaviour)
		} else {
			boards = p.placeFiguresOnBoard(boards, behaviour)
		}
	}
	return boards
}

func (p *Placement) placeFiguresOnBoard(boards map[string]string, behaviour FigureBehaviour) map[string]string {
	var resultingMap = make(map[string]string)

	for _, board := range boards {

		boardsWithPlacement := p.PlaceFiguresOnEmptyBoard(board, behaviour)

		for hash, element := range boardsWithPlacement {
			resultingMap[hash] = element
		}
	}
	return resultingMap
}

func (p *Placement) PlaceFiguresOnEmptyBoard(board string, behaviour FigureBehaviour) map[string]string {
	return behaviour.Handle(board)
}
