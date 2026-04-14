package figures

import (
	"fmt"
	"strings"
	"hash/fnv"
)

type Placement struct {
}

func GenerateHash(s []byte) string {
	h := fnv.New64a()
    h.Write(s)
	return fmt.Sprintf("%x", h.Sum64())
}

var defaultDimension = 8

const emptyField = '_'
const attackPlace = 'x'

func drawEmptyBoard() []byte {

	var board strings.Builder
	board.Grow(defaultDimension*defaultDimension + defaultDimension)

	for x := 0; x < defaultDimension; x++ {
		for y := 0; y < defaultDimension; y++ {
			board.WriteByte(emptyField)
		}
		board.WriteByte('\n')
	}
	return []byte(board.String())
}

func (p *Placement) SetDimension(value int) {
	defaultDimension = value
}

func (p *Placement) PlaceFigures(numberOfFigures int, behaviour FigureBehaviour, boards map[string][]byte) map[string][]byte {
	for i := 0; i < numberOfFigures; i++ {
		if len(boards) == 0 {
			boards = p.placeFigureOnBoard(drawEmptyBoard(), behaviour)
		} else {
			boards = p.placeFigure(boards, behaviour)
		}
	}
	return boards
}

func (p *Placement) placeFigure(boards map[string][]byte, behaviour FigureBehaviour) map[string][]byte{
	var resultingMap = make(map[string][]byte)

	for _, board := range boards {

		boardsWithPlacement := p.placeFigureOnBoard(board, behaviour)

		for hash, element := range boardsWithPlacement {
			resultingMap[hash] = element
		}
	}
	return resultingMap
}

func (p *Placement) placeFigureOnBoard(board []byte, behaviour FigureBehaviour) map[string][]byte{
	return behaviour.Handle(board)
}
