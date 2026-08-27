package figures

import (
	"sync"
)
// getDimensionFromBoard calculates the dimension of the chessboard based on the length of the board slice
// using Newton method to find the integer square root. If the length is not a perfect square, it defaults to 8.
func getDimensionFromBoard(board []byte) int {
	n := len(board)
	if n <= 0 {
		return 8
	}
	d := n
	for {
		d1 := (d + n/d) / 2
		if d1 >= d {
			break
		}
		d = d1
	}
	if d*d == n {
		return d
	}
	return 8
}

// boardPool is a shared pool for all chess figures to reuse byte slices
// This reduces allocations by reusing temporary board representations
var boardPool = sync.Pool{
	New: func() interface{} {
		// Start with a default 8x8 board size
		// Will be resized as needed in getBoardFromPool
		b := make([]byte, 64)
		return &b
	},
}

// getBoardFromPool retrieves a board slice from the pool, resizing if necessary
func getBoardFromPool(dimension int) *[]byte {
	ptr := boardPool.Get().(*[]byte)
	board := *ptr
	requiredSize := dimension * dimension

	// If the pooled board is the wrong size, resize it
	if len(board) != requiredSize {
		board = make([]byte, requiredSize)
		*ptr = board
	}

	return ptr
}

// mapPool is a shared pool for reusing maps to reduce allocations
// Maps are cleared before being returned to the pool
var mapPool = sync.Pool{
	New: func() interface{} {
		return make(map[uint64][]byte, 64) // Pre-allocate with reasonable capacity
	},
}

// getMapFromPool retrieves a clean map from the pool with the given capacity hint
func getMapFromPool() map[uint64][]byte {
	m := mapPool.Get().(map[uint64][]byte)
	// Map should already be empty from putMapToPool, but this is a safety check
	if len(m) > 0 {
		// Clear the map if it somehow wasn't cleared
		for k := range m {
			delete(m, k)
		}
	}
	return m
}

// PutMapToPool clears the map and returns it to the pool for reuse (exported for chessboard.go)
func PutMapToPool(m map[uint64][]byte) {
	putMapToPool(m)
}

// putMapToPool clears the map and returns it to the pool for reuse (internal)
func putMapToPool(m map[uint64][]byte) {
	// Clear the map before returning to pool
	for k := range m {
		delete(m, k)
	}
	mapPool.Put(m)
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
