package main

import (
	"Chessboard_in_Go/figures"
)

func (board *Chessboard) calculateBoards() map[uint64][]byte {

	return board.calculateBoard(board.currentFigureBehaviour, make(map[uint64][]byte))
}

func (board *Chessboard) calculateBoard(behaviour figures.FigureBehaviour, previousFigureBoards map[uint64][]byte) map[uint64][]byte {
	result := board.figurePlacement.PlaceFigures(board.figureQuantityMap[behaviour.GetName()], behaviour, previousFigureBoards)

	// Return the previous map to pool if it's not the initial empty map
	if len(previousFigureBoards) > 0 {
		figures.PutMapToPool(previousFigureBoards)
	}

	if behaviour.GetNext() != nil {
		result = board.calculateBoard(behaviour.GetNext(), result)
	}
	return result
}

type Chessboard struct {
	figureQuantityMap      map[byte]int
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
	figureQuantityMap      map[byte]int
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
	b.figureQuantityMap[figure.GetName()] += quantity
	return b.addToChain(figure)
}

// NewChessboard with default size 8
func NewChessboard() ChessboardBuilder {
	chessboard := &Chessboard{}
	chessboard.figurePlacement.SetDimension(8)
	return &boardBuilder{chessboard: chessboard, figureQuantityMap: make(map[byte]int)}
}

// NewChessboardWithSize custom default size
func NewChessboardWithSize(size int) ChessboardBuilder {
	chessboard := &Chessboard{}
	chessboard.figurePlacement.SetDimension(size)
	return &boardBuilder{chessboard: chessboard, figureQuantityMap: make(map[byte]int)}
}

func (b *boardBuilder) addToChain(figure figures.FigureBehaviour) ChessboardBuilder {
	if b.chessboard.currentFigureBehaviour == nil {
		b.chessboard = &Chessboard{b.figureQuantityMap, figure, b.chessboard.figurePlacement}
		b.currentFigureBehaviour = figure
	} else {
		// Check if this figure type already exists in the chain
		figureAlreadyInChain := false
		current := b.chessboard.currentFigureBehaviour
		for current != nil {
			if current.GetName() == figure.GetName() {
				figureAlreadyInChain = true
				break
			}
			current = current.GetNext()
		}

		// Only add to chain if this figure type is not already present
		if !figureAlreadyInChain {
			b.currentFigureBehaviour.SetNext(figure)
			b.currentFigureBehaviour = figure
		}
	}
	return b
}

func (b *boardBuilder) Build() *Chessboard {
	return b.chessboard
}
