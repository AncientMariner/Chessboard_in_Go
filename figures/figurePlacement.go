package figures

import (
	"fmt"
	"hash/fnv"
	"strings"
)

type Placement struct {
	currentPlacement FigurePlacement
}

type FigurePlacement interface {
	PlaceFiguresOnBoard([]string) FigurePlacement
}

type BoardWithFigurePosition struct {
	Board  string
	number int
	hash   uint32
}

func (e *BoardWithFigurePosition) Hash() string {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(e.Board))
	sum32 := algorithm.Sum32()
	e.hash = sum32
	return fmt.Sprintf("%s:%d", e.Board, sum32)
}

const defaultDimension = 8
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

func (p *Placement) PlaceFigure(numberOfFigures int, behaviour FigureBehaviour, boards map[uint32]string) map[uint32]string {
	for i := 0; i < numberOfFigures; i++ {
		if len(boards) == 0 {
			boards = p.PlaceFiguresOnEmptyBoard(drawEmptyBoard(), behaviour)
		} else {
			boards = p.placeFiguresOnBoard(boards, behaviour)
		}
	}
	return boards
}

func (p *Placement) placeFiguresOnBoard(boards map[uint32]string, behaviour FigureBehaviour) map[uint32]string {
	var resultingMap = make(map[uint32]string)

	for _, ss := range boards {
		boardsWithPlacement := p.PlaceFiguresOnEmptyBoard(ss, behaviour)

		for u, s := range boardsWithPlacement {
			resultingMap[u] = s
		}
	}
	return resultingMap
}

func (p *Placement) PlaceFiguresOnEmptyBoard(board string, behaviour FigureBehaviour) map[uint32]string {
	return behaviour.Handle(board)
}
