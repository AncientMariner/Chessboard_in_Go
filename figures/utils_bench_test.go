package figures

import (
	"testing"
)

// BenchmarkMapPoolingWithPool benchmarks map operations using the pool
func BenchmarkMapPoolingWithPool(b *testing.B) {
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	i := 0
	for b.Loop() {
		m := getMapFromPool(0)
		// Simulate typical usage: add some entries.
		// i varies keys across iterations so the map must insert, not just overwrite.
		for j := 0; j < 10; j++ {
			hash := uint64(i*10 + j)
			m[hash] = board
		}
		putMapToPool(m)
		i++
	}
}

// BenchmarkMapPoolingWithoutPool benchmarks map operations without pooling
func BenchmarkMapPoolingWithoutPool(b *testing.B) {
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	i := 0
	for b.Loop() {
		m := make(map[uint64][]byte, 64)
		// Simulate typical usage: add some entries
		for j := 0; j < 10; j++ {
			hash := uint64(i*10 + j)
			m[hash] = board
		}
		// Let GC handle it
		i++
	}
}

// BenchmarkMapPoolingWithPoolParallel benchmarks pooled map operations under parallel load
func BenchmarkMapPoolingWithPoolParallel(b *testing.B) {
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m := getMapFromPool(0)
			for j := 0; j < 10; j++ {
				hash := uint64(i*10 + j)
				m[hash] = board
			}
			putMapToPool(m)
			i++
		}
	})
}

// BenchmarkMapPoolingWithoutPoolParallel benchmarks non-pooled map operations under parallel load
func BenchmarkMapPoolingWithoutPoolParallel(b *testing.B) {
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m := make(map[uint64][]byte, 64)
			for j := 0; j < 10; j++ {
				hash := uint64(i*10 + j)
				m[hash] = board
			}
			i++
		}
	})
}

// BenchmarkKingHandleWithPooling benchmarks the King figure Handle method with pooling
func BenchmarkKingHandleWithPooling(b *testing.B) {
	king := King{}
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	for b.Loop() {
		result := king.Handle(board)
		// Return the result map to pool
		PutMapToPool(result)
	}
}

func BenchmarkQueenHandle(b *testing.B) {
	queen := Queen{}
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	for b.Loop() {
		result := queen.Handle(board)
		PutMapToPool(result)
	}
}

func BenchmarkRookHandle(b *testing.B) {
	rook := Rook{}
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	for b.Loop() {
		result := rook.Handle(board)
		PutMapToPool(result)
	}
}

func BenchmarkBishopHandle(b *testing.B) {
	bishop := Bishop{}
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	for b.Loop() {
		result := bishop.Handle(board)
		PutMapToPool(result)
	}
}

func BenchmarkKnightHandle(b *testing.B) {
	knight := Knight{}
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	for b.Loop() {
		result := knight.Handle(board)
		PutMapToPool(result)
	}
}

// BenchmarkPlacement_WithPooling benchmarks Placement.placeFigure with map pooling
func BenchmarkPlacement_WithPooling(b *testing.B) {
	boards := make(map[uint64][]byte)
	emptyBoard := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")
	boards[GenerateHash(emptyBoard)] = emptyBoard

	placement := &Placement{}
	king := &King{}

	for b.Loop() {
		result := placement.placeFigure(boards, king)
		// Clean up result map
		PutMapToPool(result)
	}
}

// BenchmarkMapPoolingLargeCapacity benchmarks the large-capacity bypass path in getMapFromPool.
// When capacityHint > 64, the pooled map is discarded and a fresh allocation is made.
// Compare with BenchmarkMapPoolingWithPool for the pool-hit cost (0 B/op).
func BenchmarkMapPoolingLargeCapacity(b *testing.B) {
	board := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	i := 0
	for b.Loop() {
		m := getMapFromPool(500) // Larger capacity: exceeds pool default (64), triggers new alloc
		// i varies keys across iterations so the map must insert, not just overwrite.
		for j := 0; j < 50; j++ {
			hash := uint64(i*50 + j)
			m[hash] = board
		}
		putMapToPool(m)
		i++
	}
}

// BenchmarkGetBoardFromPool compares pool-hit (dimension=8, matches default pool size)
// against pool-miss (dimension=4 and dimension=12, require a resize on first call).
// In practice all three converge to ~8 ns/0 B after warmup: the pool retains the
// resized slice, so subsequent calls are hits regardless of dimension.
func BenchmarkGetBoardFromPool(b *testing.B) {
	sizes := []struct {
		name      string
		dimension int
	}{
		{"4x4Miss", 4},
		{"8x8Hit", 8},
		{"12x12Miss", 12},
	}

	for _, s := range sizes {
		b.Run(s.name, func(b *testing.B) {
			for b.Loop() {
				ptr := getBoardFromPool(s.dimension)
				boardPool.Put(ptr)
			}
		})
	}
}

