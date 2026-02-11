package figures

import "strings"

func getCountOfEmptyPlaces(board string) int {
	return strings.Count(board, string(emptyField))
}
