package figures

func getCountOfEmptyPlaces(board []byte) int {
	counter := 0
	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			counter++
		}
	}
	return counter
}
