package main

import "strings"

func main() {
	drawEmptyBoard()
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
