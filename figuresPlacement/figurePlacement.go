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
	// use different figures
	// use boars provided
	return p.test(numberOfFigures)
}

func (p *Placement) test(numberOfFigures int) *set.HashSet[*FigurePosition, string] {
	var boards *set.HashSet[*FigurePosition, string]

	for i := 0; i < numberOfFigures; i++ {
		if boards == nil {
			boards = p.placeInitialFiguresOnEmptyBoard()
		} else {
			boards = p.placeFiguresOnNotEmptyBoard(boards)
		}
	}
	return boards
}

func (p *Placement) placeInitialFiguresOnEmptyBoard() *set.HashSet[*FigurePosition, string] {
	return p.PlaceFiguresOnBoard(drawEmptyBoard())
}

func (p *Placement) placeFiguresOnNotEmptyBoard(boards *set.HashSet[*FigurePosition, string]) *set.HashSet[*FigurePosition, string] {
	var resultingBoards = set.NewHashSet[*FigurePosition, string](boards.Size() * boards.Size())

	boards.ForEach(func(position *FigurePosition) bool {
		boardsWithPlacement := p.PlaceFiguresOnBoard(position.Board)

		boardsWithPlacement.ForEach(func(position *FigurePosition) bool {
			resultingBoards.Insert(position)
			return true
		})
		return true
	})
	return resultingBoards
}

func (p *Placement) PlaceFiguresOnBoard(board string) *set.HashSet[*FigurePosition, string] {

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
			out[i] = 'k' // get figure
			boardWithFigure := string(out)

			hashSetOfBoards.Insert(&FigurePosition{boardWithFigure, i})
		}
	}
	return hashSetOfBoards
}
