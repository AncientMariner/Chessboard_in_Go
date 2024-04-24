package figures

import (
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/hashicorp/go-set/v2"
)

type Placement struct {
	currentPlacement FigurePlacement
}

type FigurePlacement interface {
	PlaceFiguresOnBoard([]string) FigurePlacement
}

type FigurePosition struct {
	Board  string
	number int
}

func (e *FigurePosition) Hash() string {
	algorithm := fnv.New64()
	algorithm.Write([]byte(e.Board + string(algorithm.Sum64())))
	return fmt.Sprintf("%s:%d", e.Board, algorithm.Sum64())
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

func (p *Placement) PlaceFigure(numberOfFigures int, behaviour FigureBehaviour, boards *set.HashSet[*FigurePosition, string]) *set.HashSet[*FigurePosition, string] {
	for i := 0; i < numberOfFigures; i++ {
		if boards.Size() == 0 {
			boards = p.PlaceFiguresOnEmptyBoard(drawEmptyBoard(), behaviour)
		} else {
			boards = p.placeFiguresOnBoard(boards, behaviour)
		}
	}
	return boards
}

func (p *Placement) placeFiguresOnBoard(boards *set.HashSet[*FigurePosition, string], behaviour FigureBehaviour) *set.HashSet[*FigurePosition, string] {
	var resultingBoards = set.NewHashSet[*FigurePosition, string](boards.Size() * boards.Size())

	boards.ForEach(func(position *FigurePosition) bool {
		boardsWithPlacement := p.PlaceFiguresOnEmptyBoard(position.Board, behaviour)

		boardsWithPlacement.ForEach(func(position *FigurePosition) bool {
			resultingBoards.Insert(position)
			return true
		})
		return true
	})
	return resultingBoards
}

func (p *Placement) PlaceFiguresOnEmptyBoard(board string, behaviour FigureBehaviour) *set.HashSet[*FigurePosition, string] {
	return behaviour.Handle(board)
}
