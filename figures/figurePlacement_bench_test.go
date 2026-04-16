package figures

import (
	"testing"
)

// BenchmarkGenerateHash benchmarks the hash generation with pooled hashers
func BenchmarkGenerateHash(b *testing.B) {
	board := []byte(
		"________\n" +
			"________\n" +
			"________\n" +
			"________\n" +
			"________\n" +
			"________\n" +
			"________\n" +
			"________\n",
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateHash(board)
	}
}

// BenchmarkGenerateHashParallel benchmarks hash generation under concurrent load
func BenchmarkGenerateHashParallel(b *testing.B) {
	board := []byte(
		"________\n" +
			"________\n" +
			"________\n" +
			"________\n" +
			"________\n" +
			"________\n" +
			"________\n" +
			"________\n",
	)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GenerateHash(board)
		}
	})
}

// BenchmarkGenerateHashVariedBoards benchmarks with different board states
func BenchmarkGenerateHashVariedBoards(b *testing.B) {
	boards := [][]byte{
		[]byte("________\n________\n________\n________\n________\n________\n________\n________\n"),
		[]byte("k_______\n________\n________\n________\n________\n________\n________\n________\n"),
		[]byte("xxx_____\nxkx_____\nxxx_____\n________\n________\n________\n________\n________\n"),
		[]byte("kxkxkxkx\nxxxxxxxx\nkxkxkxkx\nxxxxxxxx\nkxkxkxkx\nxxxxxxxx\nkxkxkxkx\nxxxxxxxx\n"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateHash(boards[i%len(boards)])
	}
}
