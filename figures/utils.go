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

// getBoardFromPool retrieves a board slice from the pool, resizing if necessary
func getBoardFromPool(dimension int) *[]byte {
	ptr := boardPool.Get().(*[]byte)
	board := *ptr
	requiredSize := (dimension + 1) * dimension

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
func getMapFromPool(capacityHint int) map[uint64][]byte {
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
