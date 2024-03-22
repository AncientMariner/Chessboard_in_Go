package main

import (
    "Chessboard_in_Go/figures"
    "Chessboard_in_Go/figuresPlacement"
    "fmt"
    "github.com/hashicorp/go-set/v2"
)

func main() {
    board := NewChessBoard().withKing(1).withQueen(1).withKnight(2).withRook(2).withBishop(2).Build()
    boardsWithFigures := placeFiguresOnBoard(board)

    fmt.Printf("\nfigues %v", board)
    fmt.Printf("\nboardsWithFigures.Size() %d", boardsWithFigures.Size())
}

func placeFiguresOnBoard(board *Chessboard) *set.HashSet[*figuresPlacement.FigurePosition, string] {
    numberOfKings := board.figureQuantityMap["king"]

    // go over chain, place figure on board

    return board.figurePlacement.PlaceFigure(numberOfKings)
}

type Chessboard struct {
    figureQuantityMap      map[string]int
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
    figureQuantityMap      map[string]int
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
    return &boardBuilder{chessboard: &Chessboard{}, figureQuantityMap: make(map[string]int)}
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
