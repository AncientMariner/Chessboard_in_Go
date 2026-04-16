package figures

import (
	"testing"
)

// BenchmarkMapPooling_WithPool benchmarks map operations using the pool
func BenchmarkMapPooling_WithPool(b *testing.B) {
	board := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := getMapFromPool(64)
		// Simulate typical usage: add some entries
		for j := 0; j < 10; j++ {
			hash := uint64(i*10 + j)
			m[hash] = board
		}
		putMapToPool(m)
	}
}

// BenchmarkMapPooling_WithoutPool benchmarks map operations without pooling
func BenchmarkMapPooling_WithoutPool(b *testing.B) {
	board := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[uint64][]byte, 64)
		// Simulate typical usage: add some entries
		for j := 0; j < 10; j++ {
			hash := uint64(i*10 + j)
			m[hash] = board
		}
		// Let GC handle it
	}
}

// BenchmarkMapPooling_WithPool_Parallel benchmarks pooled map operations under parallel load
func BenchmarkMapPooling_WithPool_Parallel(b *testing.B) {
	board := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m := getMapFromPool(64)
			for j := 0; j < 10; j++ {
				hash := uint64(i*10 + j)
				m[hash] = board
			}
			putMapToPool(m)
			i++
		}
	})
}

// BenchmarkMapPooling_WithoutPool_Parallel benchmarks non-pooled map operations under parallel load
func BenchmarkMapPooling_WithoutPool_Parallel(b *testing.B) {
	board := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")

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

// BenchmarkKingHandle_WithPooling benchmarks the King figure Handle method with pooling
func BenchmarkKingHandle_WithPooling(b *testing.B) {
	king := King{}
	board := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := king.Handle(board)
		// Return the result map to pool
		PutMapToPool(result)
	}
}

// BenchmarkPlacement_WithPooling benchmarks Placement.placeFigure with map pooling
func BenchmarkPlacement_WithPooling(b *testing.B) {
	boards := make(map[uint64][]byte)
	emptyBoard := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")
	boards[GenerateHash(emptyBoard)] = emptyBoard

	placement := &Placement{}
	king := &King{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.placeFigure(boards, king)
		// Clean up result map
		PutMapToPool(result)
	}
}

// BenchmarkMapPooling_LargeCapacity benchmarks pooled maps with larger capacity hints
func BenchmarkMapPooling_LargeCapacity(b *testing.B) {
	board := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := getMapFromPool(256) // Larger capacity
		// Simulate typical usage with more entries
		for j := 0; j < 50; j++ {
			hash := uint64(i*50 + j)
			m[hash] = board
		}
		putMapToPool(m)
	}
}

// BenchmarkMapPooling_SmallCapacity benchmarks pooled maps with small capacity hints
func BenchmarkMapPooling_SmallCapacity(b *testing.B) {
	board := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := getMapFromPool(8) // Small capacity
		// Simulate typical usage with few entries
		for j := 0; j < 3; j++ {
			hash := uint64(i*3 + j)
			m[hash] = board
		}
		putMapToPool(m)
	}
}
