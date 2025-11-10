package main

import (
	"Chessboard_in_Go/figures"
)

func main() {
	// test example
	// board := NewChessboard().withKing(1).withQueen(2).Build()
	// 	board := NewChessboard().withKing(1).Build()
	// board := NewChessboard().withKing(1).withQueen(1).withKnight(2).withRook(2).withBishop(2).Build()
	// 	boardsWithFigures := board.calculateBoards()
	//
	// 	fmt.Printf("\nfigures %v", board)
	// 	fmt.Printf("\nboardsWithFigures.Size() %d", len(boardsWithFigures))
}

func (board *Chessboard) calculateBoards() map[string]string {

	return board.calculateBoard(board.currentFigureBehaviour, make(map[string]string))
}

func (board *Chessboard) calculateBoard(behaviour figures.FigureBehaviour, previousFigureBoards map[string]string) map[string]string {
	result := board.figurePlacement.PlaceFigures(board.figureQuantityMap[behaviour.GetName()], behaviour, previousFigureBoards)

	if behaviour.GetNext() != nil {
		result = board.calculateBoard(behaviour.GetNext(), result)
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
	addFigure(figures.FigureBehaviour, int) ChessboardBuilder
	Build() *Chessboard
}

type boardBuilder struct {
	chessboard             *Chessboard
	currentFigureBehaviour figures.FigureBehaviour
	figureQuantityMap      map[rune]int
}

func (b *boardBuilder) withKing(quantity int) ChessboardBuilder {
	return b.addFigure(&figures.King{}, quantity)
}

func (b *boardBuilder) withQueen(quantity int) ChessboardBuilder {
	return b.addFigure(&figures.Queen{}, quantity)
}

func (b *boardBuilder) withBishop(quantity int) ChessboardBuilder {
	return b.addFigure(&figures.Bishop{}, quantity)
}

func (b *boardBuilder) withKnight(quantity int) ChessboardBuilder {
	return b.addFigure(&figures.Knight{}, quantity)
}

func (b *boardBuilder) withRook(quantity int) ChessboardBuilder {
	return b.addFigure(&figures.Rook{}, quantity)
}

func (b *boardBuilder) addFigure(figure figures.FigureBehaviour, quantity int) ChessboardBuilder {
	b.figureQuantityMap[figure.GetName()] = quantity
	return b.addToChain(figure)
}

// NewChessboard with default size 8
func NewChessboard() ChessboardBuilder {
	chessboard := &Chessboard{}
	chessboard.figurePlacement.SetDimension(8)
	return &boardBuilder{chessboard: chessboard, figureQuantityMap: make(map[rune]int)}
}

// NewChessboardWithSize custom default size
func NewChessboardWithSize(size int) ChessboardBuilder {
	chessboard := &Chessboard{}
	chessboard.figurePlacement.SetDimension(size)
	return &boardBuilder{chessboard: chessboard, figureQuantityMap: make(map[rune]int)}
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
