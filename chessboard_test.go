package main

import (
    "reflect"
    "testing"
)

func Test_drawEmptyBoard(t *testing.T) {
    tests := []struct {
        name string
        want string
    }{
        {"Initial board test", "________\n________\n________\n________\n________\n________\n________\n________\n"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := drawEmptyBoard(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("drawEmptyBoard() = %v, want %v", got, tt.want)
            }
        })
    }
}
