package main

import (
	"Chessboard_in_Go/figures"
	"Chessboard_in_Go/figuresPlacement"
	"fmt"
	"github.com/hashicorp/go-set/v2"
)

func main() {
	board := NewChessBoard().withKing(1).withQueen(1).Build()
	// board := NewChessBoard().withKing(1).withQueen(1).withKnight(2).withRook(2).withBishop(2).Build()
	boardsWithFigures := placeFigures(board)

	fmt.Printf("\nfigures %v", board)
	fmt.Printf("\nboardsWithFigures.Size() %d", boardsWithFigures.Size())
}

func placeFigures(board *Chessboard) *set.HashSet[*figuresPlacement.FigurePosition, string] {
	numberOfKings := board.figureQuantityMap['k']

	boards := placeFigure(board, numberOfKings, board.currentFigureBehaviour, set.NewHashSet[*figuresPlacement.FigurePosition, string](0))
	return boards
}

func placeFigure(board *Chessboard, numberOrFigures int, behaviour figures.FigureBehaviour, previousFigureBoards *set.HashSet[*figuresPlacement.FigurePosition, string]) *set.HashSet[*figuresPlacement.FigurePosition, string] {

	// extract no need to put board param here
	boards := board.figurePlacement.PlaceFigure(numberOrFigures, behaviour, previousFigureBoards)

	var result = set.NewHashSet[*figuresPlacement.FigurePosition, string](previousFigureBoards.Size() + boards.Size()) // check to calculate empty places in order to set proper size

	boards.ForEach(func(position *figuresPlacement.FigurePosition) bool {
		result.Insert(position)
		return true
	})

	if behaviour.GetNext() != nil {
		placeFigure(board, numberOrFigures, behaviour.GetNext(), result)
	}

	return result
}

type Chessboard struct {
	figureQuantityMap      map[rune]int
	currentFigureBehaviour figures.FigureBehaviour
	figurePlacement        figuresPlacement.Placement
}

type ChessBoardBuilder interface {
	withKing(int) ChessBoardBuilder
	withQueen(int) ChessBoardBuilder
	withBishop(int) ChessBoardBuilder
	withKnight(int) ChessBoardBuilder
	withRook(int) ChessBoardBuilder
	Build() *Chessboard
}

type boardBuilder struct {
	chessboard             *Chessboard
	currentFigureBehaviour figures.FigureBehaviour
	figureQuantityMap      map[rune]int
}

func (b *boardBuilder) withKing(quantity int) ChessBoardBuilder {
	figure := &figures.King{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func (b *boardBuilder) withQueen(quantity int) ChessBoardBuilder {
	figure := &figures.Queen{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func (b *boardBuilder) withBishop(quantity int) ChessBoardBuilder {
	figure := &figures.Bishop{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func (b *boardBuilder) withKnight(quantity int) ChessBoardBuilder {
	figure := &figures.Knight{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func (b *boardBuilder) withRook(quantity int) ChessBoardBuilder {
	figure := &figures.Rook{}
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

func NewChessBoard() ChessBoardBuilder {
	return &boardBuilder{chessboard: &Chessboard{}, figureQuantityMap: make(map[rune]int)}
}

func (b *boardBuilder) addToChain(figure figures.FigureBehaviour) ChessBoardBuilder {
	if b.chessboard.currentFigureBehaviour == nil {
		b.chessboard = &Chessboard{currentFigureBehaviour: figure, figureQuantityMap: b.figureQuantityMap}
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
