package figuresPlacement

import (
    "fmt"
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
