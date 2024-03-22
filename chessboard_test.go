package main

import (
    "testing"
)

func Test_placeFiguresOnBoard(t *testing.T) {
    type args struct {
        board *Chessboard
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {"Test empty board with 1 king", args{board: NewChessBoard().withKing(1).Build()}, 64},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := placeFigures(tt.args.board); got.Size() != tt.want {
                t.Errorf("placeFigures() = %v, want %v", got, tt.want)
            }
        })
    }
}
