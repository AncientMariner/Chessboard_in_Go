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
		[]byte(
			"________" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________"),
		[]byte(
			"k_______" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________"),
		[]byte(
			"xxx_____" +
				"xkx_____" +
				"xxx_____" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________"),
		[]byte(
			"kxkxkxkx" +
				"xxxxxxxx" +
				"kxkxkxkx" +
				"xxxxxxxx" +
				"kxkxkxkx" +
				"xxxxxxxx" +
				"kxkxkxkx" +
				"xxxxxxxx"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateHash(boards[i%len(boards)])
	}
}
