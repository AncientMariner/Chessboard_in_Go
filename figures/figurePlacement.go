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
	Board string
	// number int
	hash string
}

func (e *BoardWithFigurePosition) Hash() string {
	// algorithm := sha256.New()
	// algorithm.Write([]byte(e.Board))
	//
	// sum32 := algorithm.Sum(nil)
	//

	algorithm := fnv.New32a()
	algorithm.Write([]byte(e.Board))
	sum32 := algorithm.Sum32()

	e.hash = fmt.Sprintf("%x", sum32)
	return e.hash
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

		for u, _ := range boardsWithPlacement {
			resultingMap[u] = u
		}
	}
	return resultingMap
}

func (p *Placement) PlaceFiguresOnEmptyBoard(board string, behaviour FigureBehaviour) map[string]string {
	return behaviour.Handle(board)
}
