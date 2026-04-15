package figures

import "sync"

// boardPool is a shared pool for all chess figures to reuse byte slices
// This reduces allocations by reusing temporary board representations
var boardPool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, (defaultDimension+1)*defaultDimension)
		return &b
	},
}

func getCountOfEmptyPlaces(board []byte) int {
	counter := 0
	for i := 0; i < len(board); i++ {
		if board[i] == emptyField {
			counter++
		}
	}
	return counter
}
