package main

import (
    "Chessboard_in_Go/figures"
    "Chessboard_in_Go/figuresPlacement"
    "fmt"
    "github.com/hashicorp/go-set/v2"
    "strings"
)

func main() {
    drawEmptyBoard()
    board := NewChessBoard().withKing(2).withQueen(1).withKnight(2).withRook(2).withBishop(2).Build()
    placeFiguresOnBoard(board)

    fmt.Printf("figues %v", board)
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

func placeFiguresOnBoard(board *Chessboard) *set.HashSet[*figuresPlacement.FigurePosition, string] {
    numberOfKings := board.figureQuantityMap["king"]

    // go over chain, place figure on board

    var boards *set.HashSet[*figuresPlacement.FigurePosition, string]
    var resultingBoards *set.HashSet[*figuresPlacement.FigurePosition, string]

    for i := 0; i < numberOfKings; i++ {

        // if board is empty call with it
        // else call with each element from set
        if boards == nil {
            boardsAfterInitialPlacement := board.figurePlacement.PlaceFiguresOnBoard(drawEmptyBoard())
            boards = set.NewHashSet[*figuresPlacement.FigurePosition, string](boardsAfterInitialPlacement.Size())

            boardsAfterInitialPlacement.ForEach(func(position *figuresPlacement.FigurePosition) bool {
                boards.Insert(position)
                return true
            })
        } else {
            resultingBoards = set.NewHashSet[*figuresPlacement.FigurePosition, string](boards.Size() * boards.Size())

            boards.ForEach(func(position *figuresPlacement.FigurePosition) bool {
                boardsWithPlacement := board.figurePlacement.PlaceFiguresOnBoard(position.Board)

                boardsWithPlacement.ForEach(func(position *figuresPlacement.FigurePosition) bool {
                    resultingBoards.Insert(position)
                    return true
                })
                return true
            })
        }
    }
    return resultingBoards
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

func (bb *boardBuilder) withKing(quantity int) ChessBoardBuilder {
    figure := &figures.King{}
    bb.figureQuantityMap[figure.GetName()] = quantity
    return bb.addToChain(figure)
}

func (bb *boardBuilder) withQueen(quantity int) ChessBoardBuilder {
    figure := &figures.Queen{}
    bb.figureQuantityMap[figure.GetName()] = quantity
    return bb.addToChain(figure)
}

func (bb *boardBuilder) withBishop(quantity int) ChessBoardBuilder {
    figure := &figures.Bishop{}
    bb.figureQuantityMap[figure.GetName()] = quantity
    return bb.addToChain(figure)
}

func (bb *boardBuilder) withKnight(quantity int) ChessBoardBuilder {
    figure := &figures.Knight{}
    bb.figureQuantityMap[figure.GetName()] = quantity
    return bb.addToChain(figure)
}

func (bb *boardBuilder) withRook(quantity int) ChessBoardBuilder {
    figure := &figures.Rook{}
    bb.figureQuantityMap[figure.GetName()] = quantity
    return bb.addToChain(figure)
}

func NewChessBoard() ChessBoardBuilder {
    return &boardBuilder{chessboard: &Chessboard{}, figureQuantityMap: make(map[string]int)}
}

func (bb *boardBuilder) addToChain(figure figures.FigureBehaviour) ChessBoardBuilder {
    if bb.chessboard.currentFigureBehaviour == nil {
        bb.chessboard = &Chessboard{currentFigureBehaviour: figure, figureQuantityMap: bb.figureQuantityMap}
    } else {
        bb.currentFigureBehaviour.SetNext(figure)
    }
    // is needed in order to have the recent added figure and add a link to it a
    bb.currentFigureBehaviour = figure
    return bb
}

func (bb *boardBuilder) Build() *Chessboard {
    return bb.chessboard
}
