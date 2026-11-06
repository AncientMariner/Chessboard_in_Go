package figures

import (
	"fmt"
	"runtime"
	"testing"
)

// Helper function to create multiple boards for testing
func createMultipleBoards(count int) map[uint64][]byte {
	boards := make(map[uint64][]byte, count)

	// Create variations of boards with a king placed at different positions
	emptyBoard := []byte("________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________" +
		"________")

	king := &King{}
	result := king.Handle(emptyBoard)

	// Use the first 'count' boards from the result
	i := 0
	for hash, board := range result {
		if i >= count {
			break
		}
		boards[hash] = board
		i++
	}

	return boards
}

// BenchmarkPlaceFigureSequentialSmall benchmarks sequential processing with few boards
func BenchmarkPlaceFigureSequentialSmall(b *testing.B) {
	boards := createMultipleBoards(5) // Below parallelThreshold
	placement := &Placement{}
	rook := &Rook{}

	for b.Loop() {
		result := placement.placeFigure(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigureSequentialMedium benchmarks with medium board count (at threshold)
func BenchmarkPlaceFigureSequentialMedium(b *testing.B) {
	boards := createMultipleBoards(10) // At parallelThreshold
	placement := &Placement{}
	rook := &Rook{}

	for b.Loop() {
		result := placement.placeFigure(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigureParallelLarge benchmarks parallel processing with many boards
func BenchmarkPlaceFigureParallelLarge(b *testing.B) {
	boards := createMultipleBoards(50) // Well above parallelThreshold
	placement := &Placement{}
	rook := &Rook{}

	for b.Loop() {
		result := placement.placeFigure(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigureParallelVeryLarge benchmarks parallel processing with very many boards
func BenchmarkPlaceFigureParallelVeryLarge(b *testing.B) {
	boards := createMultipleBoards(64) // Maximum for 8x8 board with 1 figure
	placement := &Placement{}
	rook := &Rook{}

	for b.Loop() {
		result := placement.placeFigure(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigureSequentialDirect benchmarks the sequential implementation directly
func BenchmarkPlaceFigureSequentialDirect(b *testing.B) {
	boards := createMultipleBoards(50)
	placement := &Placement{}
	rook := &Rook{}

	for b.Loop() {
		result := placement.placeFigureSequential(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigureParallelDirect benchmarks the parallel implementation directly
func BenchmarkPlaceFigureParallelDirect(b *testing.B) {
	boards := createMultipleBoards(50)
	placement := &Placement{}
	rook := &Rook{}

	for b.Loop() {
		result := placement.placeFigureParallel(boards, rook)
		PutMapToPool(result)
	}
}

func BenchmarkPlaceFigureSequential1000(b *testing.B) {
	boards := createMultipleBoards(1000)
	placement := &Placement{}
	rook := &Rook{}

	for b.Loop() {
		result := placement.placeFigureSequential(boards, rook)
		PutMapToPool(result)
	}
}

func BenchmarkPlaceFigureParallel1000(b *testing.B) {
	boards := createMultipleBoards(1000)
	placement := &Placement{}
	rook := &Rook{}

	for b.Loop() {
		result := placement.placeFigureParallel(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFiguresMultipleIterationsSequential benchmarks the full PlaceFigures workflow
func BenchmarkPlaceFiguresMultipleIterationsSequential(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	king := &King{}

	for b.Loop() {
		result := placement.PlaceFigures(2, king, make(map[uint64][]byte))
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigures3Kings benchmarks placing 3 kings (realistic scenario)
func BenchmarkPlaceFigures3Kings(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	king := &King{}

	for b.Loop() {
		result := placement.PlaceFigures(3, king, make(map[uint64][]byte))
		PutMapToPool(result)
	}
}

// BenchmarkComplexScenarioKingAndRook benchmarks a complex multi-figure scenario
func BenchmarkComplexScenarioKingAndRook(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	king := &King{}
	rook := &Rook{}

	for b.Loop() {
		boards := make(map[uint64][]byte)
		boards = placement.PlaceFigures(1, king, boards)
		boards = placement.PlaceFigures(1, rook, boards)
		PutMapToPool(boards)
	}
}

func BenchmarkPuzzle8Queens(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	queen := &Queen{}

	for b.Loop() {
		result := placement.PlaceFigures(8, queen, make(map[uint64][]byte))
		PutMapToPool(result)
	}
}

func BenchmarkPuzzle2Rooks2Kings(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	king := &King{}
	rook := &Rook{}

	for b.Loop() {
		boards := make(map[uint64][]byte)
		boards = placement.PlaceFigures(2, king, boards)
		boards = placement.PlaceFigures(2, rook, boards)
		PutMapToPool(boards)
	}
}

func BenchmarkPuzzleAllPieces(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	king := &King{}
	queen := &Queen{}
	rook := &Rook{}
	bishop := &Bishop{}
	knight := &Knight{}

	for b.Loop() {
		// runtime.GC() stabilises heap state before each iteration, preventing GC debt
		// from one pass bleeding into the next. Without it, B/op varies 2-3x across runs
		// due to the large working set (~4-9 GB) this benchmark allocates.
		runtime.GC()
		boards := make(map[uint64][]byte)
		boards = placement.PlaceFigures(1, king, boards)
		boards = placement.PlaceFigures(1, queen, boards)
		boards = placement.PlaceFigures(1, rook, boards)
		boards = placement.PlaceFigures(1, bishop, boards)
		boards = placement.PlaceFigures(1, knight, boards)
		PutMapToPool(boards)
	}
}

// BenchmarkBoardSizeScaling measures how placement cost grows with board dimension.
// Uses a single king to isolate board-size effect from piece interaction complexity.
func BenchmarkBoardSizeScaling(b *testing.B) {
	sizes := []int{4, 6, 8, 10, 12}

	for _, dim := range sizes {
		b.Run(fmt.Sprintf("%dx%d", dim, dim), func(b *testing.B) {
			placement := &Placement{}
			placement.SetDimension(dim)
			king := &King{}

			for b.Loop() {
				result := placement.PlaceFigures(1, king, make(map[uint64][]byte))
				PutMapToPool(result)
			}
		})
	}
}
