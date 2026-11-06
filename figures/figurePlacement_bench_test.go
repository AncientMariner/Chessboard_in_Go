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

	for b.Loop() {
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

	i := 0
	for b.Loop() {
		_ = GenerateHash(boards[i%len(boards)])
		i++
	}
}

// BenchmarkHashStrategy compares full rehash (GenerateHash/ZobristHash) against
// incremental update (ZobristUpdate) for a single-cell change on an 8x8 board.
// ZobristUpdate should be O(1) vs O(n) for the full-scan variants.
func BenchmarkHashStrategy(b *testing.B) {
	board := []byte(
		"________" +
			"________" +
			"________" +
			"____k___" +
			"________" +
			"________" +
			"________" +
			"________",
	)
	const changedCell = 36 // centre cell where the king sits
	baseHash := ZobristHash(board)

	b.Run("GenerateHashFull", func(b *testing.B) {
		for b.Loop() {
			_ = GenerateHash(board)
		}
	})

	b.Run("ZobristHashFull", func(b *testing.B) {
		for b.Loop() {
			_ = ZobristHash(board)
		}
	})

	b.Run("ZobristUpdateIncremental", func(b *testing.B) {
		for b.Loop() {
			_ = ZobristUpdate(baseHash, changedCell, 'k', '_')
		}
	})
}
