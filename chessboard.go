package main

import (
	"Chessboard_in_Go/figures"
	"fmt"
	"strings"
)

func main() {
	drawEmptyBoard()
	boardWithRookAndBishop := NewBuilder(map[string]int{"rook": 1, "bishop": 1}).withRook().withBishop().Build()

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
	withKing() ChessBoardBuilder
	withRook() ChessBoardBuilder
	withBishop() ChessBoardBuilder
	Build() *Chessboard
}

type boardBuilder struct {
	currentFigureBehaviour figures.FigureBehaviour
	figureQuantityMap      map[string]int
}

func (bb *boardBuilder) withRook() ChessBoardBuilder {
	return bb.addToChain(&figures.Rook{})
}

func (bb *boardBuilder) withBishop() ChessBoardBuilder {
	return bb.addToChain(&figures.Bishop{})
}

func (bb *boardBuilder) withKing() ChessBoardBuilder {
	return bb.addToChain(&figures.King{})
}

func NewBuilder(figureQuantityMap map[string]int) ChessBoardBuilder {
	return &boardBuilder{figureQuantityMap: figureQuantityMap}
}

func (bb *boardBuilder) addToChain(figure figures.FigureBehaviour) ChessBoardBuilder {
	if bb.currentFigureBehaviour == nil {
		bb.currentFigureBehaviour = figure
	} else {
		bb.currentFigureBehaviour.SetNext(figure)
	}
	return bb
}

func (bb *boardBuilder) Build() *Chessboard {
	return &Chessboard{
		figureQuantityMap:      bb.figureQuantityMap,
		currentFigureBehaviour: bb.currentFigureBehaviour,
	}
}
