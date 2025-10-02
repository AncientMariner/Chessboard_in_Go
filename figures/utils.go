package figures

func getCountOfEmptyPlaces(board string) (countOfEmptyPlaces int) {
	countOfEmptyPlaces = 0
	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			countOfEmptyPlaces++
		}
	}
	return
}
