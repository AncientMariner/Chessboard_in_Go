package main

import (
	"Chessboard_in_Go/figures"
	"fmt"
	// "github.com/hashicorp/go-set/v2"
)

func main() {
	// board := NewChessboard().withKing(1).withQueen(2).Build()
	board := NewChessboard().withKing(1).Build()
	// board := NewChessboard().withKing(1).withQueen(1).withKnight(2).withRook(2).withBishop(2).Build()
	boardsWithFigures := board.placeFigures()

	fmt.Printf("\nfigures %v", board)
	fmt.Printf("\nboardsWithFigures.Size() %d", len(boardsWithFigures))
}

func (board *Chessboard) placeFigures() map[uint32]string {

	return board.placeFigure(board.currentFigureBehaviour, make(map[uint32]string))
}

func (board *Chessboard) placeFigure(behaviour figures.FigureBehaviour, previousFigureBoards map[uint32]string) map[uint32]string {
	result := board.figurePlacement.PlaceFigure(board.figureQuantityMap[behaviour.GetName()], behaviour, previousFigureBoards)

	if behaviour.GetNext() != nil {
		result = board.placeFigure(behaviour.GetNext(), result)
	}
	return result
}

type Chessboard struct {
	figureQuantityMap      map[rune]int
	currentFigureBehaviour figures.FigureBehaviour
	figurePlacement        figures.Placement
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
	// is needed in order to have the recent added figure and add a link to it
	b.currentFigureBehaviour = figure
	return b
}

func (b *boardBuilder) Build() *Chessboard {
	return b.chessboard
}
