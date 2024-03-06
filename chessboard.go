package main

import (
	"Chessboard_in_Go/figures"
	"fmt"
	"strings"
)

func main() {
	drawEmptyBoard()
	boardWithRookAndBishop := NewBuilder().withKing(1).withRook(1).withBishop(1).Build()

	fmt.Println(boardWithRookAndBishop)
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

type Chessboard struct {
	figureQuantityMap      map[string]int
	currentFigureBehaviour figures.FigureBehaviour
}

type ChessBoardBuilder interface {
	withKing(int) ChessBoardBuilder
	withRook(int) ChessBoardBuilder
	withBishop(int) ChessBoardBuilder
	Build() *Chessboard
}

type boardBuilder struct {
	chessboard             *Chessboard
	currentFigureBehaviour figures.FigureBehaviour
	figureQuantityMap      map[string]int
}

func (bb *boardBuilder) withRook(quantity int) ChessBoardBuilder {
	figure := &figures.Rook{}
	bb.figureQuantityMap[figure.GetName()] = quantity
	return bb.addToChain(figure)
}

func (bb *boardBuilder) withBishop(quantity int) ChessBoardBuilder {
	figure := &figures.Bishop{}
	bb.figureQuantityMap[figure.GetName()] = quantity
	return bb.addToChain(figure)
}

func (bb *boardBuilder) withKing(quantity int) ChessBoardBuilder {
	figure := &figures.King{}
	bb.figureQuantityMap[figure.GetName()] = quantity
	return bb.addToChain(figure)
}

func NewBuilder() ChessBoardBuilder {
	return &boardBuilder{chessboard: &Chessboard{}, figureQuantityMap: make(map[string]int)}
}

func (bb *boardBuilder) addToChain(figure figures.FigureBehaviour) ChessBoardBuilder {
	if bb.chessboard.currentFigureBehaviour == nil {
		bb.chessboard = &Chessboard{currentFigureBehaviour: figure, figureQuantityMap: bb.figureQuantityMap}
	} else {
		bb.currentFigureBehaviour.SetNext(figure)
	}
	// needed in order to have the recent added figure and add a link to it a
	bb.currentFigureBehaviour = figure
	return bb
}

func (bb *boardBuilder) Build() *Chessboard {
	var chessboard = bb.chessboard
	return chessboard
}
