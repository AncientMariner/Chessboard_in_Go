package main

import (
	"Chessboard_in_Go/figures"
	"Chessboard_in_Go/figuresPlacement"
	"fmt"
	"github.com/hashicorp/go-set/v2"
)

func main() {
	board := NewChessboard().withKing(1).withQueen(2).Build()
	// board := NewChessboard().withKing(1).withQueen(1).withKnight(2).withRook(2).withBishop(2).Build()
	boardsWithFigures := placeFigures(board)

	fmt.Printf("\nfigures %v", board)
	fmt.Printf("\nboardsWithFigures.Size() %d", boardsWithFigures.Size())
}

func placeFigures(board *Chessboard) *set.HashSet[*figuresPlacement.FigurePosition, string] {

	return placeFigure(board, board.currentFigureBehaviour, set.NewHashSet[*figuresPlacement.FigurePosition, string](0))
}

// add all combinations of figures placement
func placeFigure(board *Chessboard, behaviour figures.FigureBehaviour, previousFigureBoards *set.HashSet[*figuresPlacement.FigurePosition, string]) *set.HashSet[*figuresPlacement.FigurePosition, string] {
	numberOfFigures := board.figureQuantityMap[behaviour.GetName()]

	// extract no need to put board param here
	boards := board.figurePlacement.PlaceFigure(numberOfFigures, behaviour, previousFigureBoards)

	var result = set.NewHashSet[*figuresPlacement.FigurePosition, string](previousFigureBoards.Size() + boards.Size()) // check to calculate empty places in order to set proper size

	boards.ForEach(func(position *figuresPlacement.FigurePosition) bool {
		result.Insert(position)
		return true
	})

	if behaviour.GetNext() != nil {
		placeFigure(board, behaviour.GetNext(), result)
	}

	return result
}

type Chessboard struct {
	figureQuantityMap      map[rune]int
	currentFigureBehaviour figures.FigureBehaviour
	figurePlacement        figuresPlacement.Placement
}

type ChessboardBuilder interface {
	withKing(int) ChessboardBuilder
	withQueen(int) ChessboardBuilder
	withBishop(int) ChessboardBuilder
	withKnight(int) ChessboardBuilder
	withRook(int) ChessboardBuilder
	Build() *Chessboard
}

type boardBuilder struct {
	chessboard             *Chessboard
	currentFigureBehaviour figures.FigureBehaviour
	figureQuantityMap      map[rune]int
}

func (b *boardBuilder) withKing(quantity int) ChessboardBuilder {
	figure := &figures.King{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func (b *boardBuilder) withQueen(quantity int) ChessboardBuilder {
	figure := &figures.Queen{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func (b *boardBuilder) withBishop(quantity int) ChessboardBuilder {
	figure := &figures.Bishop{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func (b *boardBuilder) withKnight(quantity int) ChessboardBuilder {
	figure := &figures.Knight{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func (b *boardBuilder) withRook(quantity int) ChessboardBuilder {
	figure := &figures.Rook{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func NewChessboard() ChessboardBuilder {
	return &boardBuilder{chessboard: &Chessboard{}, figureQuantityMap: make(map[rune]int)}
}

func (b *boardBuilder) addToChain(figure figures.FigureBehaviour) ChessboardBuilder {
	if b.chessboard.currentFigureBehaviour == nil {
		b.chessboard = &Chessboard{b.figureQuantityMap, figure, b.chessboard.figurePlacement}
	} else {
		b.currentFigureBehaviour.SetNext(figure)
	}
	// is needed in order to have the recent added figure and add a link to it a
	b.currentFigureBehaviour = figure
	return b
}

func (b *boardBuilder) Build() *Chessboard {
	return b.chessboard
}
