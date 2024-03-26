package figuresPlacement

import (
	"Chessboard_in_Go/figures"
	"fmt"
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
	return fmt.Sprintf("%s:%d", e.Board, e.number)
}

const defaultDimension = 8

func drawEmptyBoard() string {

	var board strings.Builder

	for x := 0; x < defaultDimension; x++ {
		for y := 0; y < defaultDimension; y++ {
			board.WriteString("_")
		}
		board.WriteString("\n")
	}
	return board.String()
}

func (p *Placement) PlaceFigure(numberOfFigures int, behaviour figures.FigureBehaviour, boards *set.HashSet[*FigurePosition, string]) *set.HashSet[*FigurePosition, string] {
	for i := 0; i < numberOfFigures; i++ {
		if boards.Size() == 0 {
			boards = p.PlaceFiguresOnEmptyBoard(drawEmptyBoard(), behaviour)
		} else {
			boards = p.placeFiguresOnBoard(boards, behaviour)
		}
	}
	return boards
}

func (p *Placement) placeFiguresOnBoard(boards *set.HashSet[*FigurePosition, string], behaviour figures.FigureBehaviour) *set.HashSet[*FigurePosition, string] {
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

func (p *Placement) PlaceFiguresOnEmptyBoard(board string, behaviour figures.FigureBehaviour) *set.HashSet[*FigurePosition, string] {

	countOfEmptyPlaces := 0
	for i := 0; i < len(board); i++ {
		if board[i] == '_' {
			countOfEmptyPlaces++
		}
	}

	hashSetOfBoards := set.NewHashSet[*FigurePosition, string](countOfEmptyPlaces)

	for i := 0; i < len(board); i++ {
		if board[i] == '_' {
			out := []rune(board)
			out[i] = behaviour.GetName()
			boardWithFigure := string(out)

			hashSetOfBoards.Insert(&FigurePosition{boardWithFigure, i})
		}
	}
	return hashSetOfBoards
}
