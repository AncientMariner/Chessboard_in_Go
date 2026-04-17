package figures

import (
	"testing"
)

// BenchmarkGenerateHash benchmarks the hash generation with pooled hashers
func BenchmarkGenerateHash(b *testing.B) {
	board := []byte(
		"________" +
			"________" +
			"________" +
			"________" +
			"________" +
			"________" +
			"________" +
			"________",
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateHash(board)
	}
}

// BenchmarkGenerateHashParallel benchmarks hash generation under concurrent load
func BenchmarkGenerateHashParallel(b *testing.B) {
	board := []byte(
		"________" +
			"________" +
			"________" +
			"________" +
			"________" +
			"________" +
			"________" +
			"________",
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
		[]byte("________________________________________________________________"),
		[]byte("k_______________________________________________________________"),
		[]byte("xxx_____xkx_____xxx_____________________________________________"),
		[]byte("kxkxkxkxxxxxxxxxkxkxkxkxxxxxxxxxkxkxkxkxxxxxxxxxkxkxkxkxxxxxxxxx"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateHash(boards[i%len(boards)])
	}
}
