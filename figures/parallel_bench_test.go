package figures

import (
	"testing"
)

// Helper function to create multiple boards for testing
func createMultipleBoards(count int) map[uint64][]byte {
	boards := make(map[uint64][]byte, count)

	// Create variations of boards with a king placed at different positions
	emptyBoard := []byte("________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n" +
		"________\n")

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

// BenchmarkPlaceFigure_Sequential_Small benchmarks sequential processing with few boards
func BenchmarkPlaceFigure_Sequential_Small(b *testing.B) {
	boards := createMultipleBoards(5) // Below parallelThreshold
	placement := &Placement{}
	rook := &Rook{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.placeFigure(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigure_Sequential_Medium benchmarks with medium board count (at threshold)
func BenchmarkPlaceFigure_Sequential_Medium(b *testing.B) {
	boards := createMultipleBoards(10) // At parallelThreshold
	placement := &Placement{}
	rook := &Rook{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.placeFigure(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigure_Parallel_Large benchmarks parallel processing with many boards
func BenchmarkPlaceFigure_Parallel_Large(b *testing.B) {
	boards := createMultipleBoards(50) // Well above parallelThreshold
	placement := &Placement{}
	rook := &Rook{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.placeFigure(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigure_Parallel_VeryLarge benchmarks parallel processing with very many boards
func BenchmarkPlaceFigure_Parallel_VeryLarge(b *testing.B) {
	boards := createMultipleBoards(64) // Maximum for 8x8 board with 1 figure
	placement := &Placement{}
	rook := &Rook{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.placeFigure(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigureSequential_Direct benchmarks the sequential implementation directly
func BenchmarkPlaceFigureSequential_Direct(b *testing.B) {
	boards := createMultipleBoards(50)
	placement := &Placement{}
	rook := &Rook{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.placeFigureSequential(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigureParallel_Direct benchmarks the parallel implementation directly
func BenchmarkPlaceFigureParallel_Direct(b *testing.B) {
	boards := createMultipleBoards(50)
	placement := &Placement{}
	rook := &Rook{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.placeFigureParallel(boards, rook)
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigures_MultipleIterations benchmarks the full PlaceFigures workflow
func BenchmarkPlaceFigures_MultipleIterations_Sequential(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	king := &King{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.PlaceFigures(2, king, make(map[uint64][]byte))
		PutMapToPool(result)
	}
}

// BenchmarkPlaceFigures_3Kings benchmarks placing 3 kings (realistic scenario)
func BenchmarkPlaceFigures_3Kings(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	king := &King{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := placement.PlaceFigures(3, king, make(map[uint64][]byte))
		PutMapToPool(result)
	}
}

// BenchmarkComplexScenario benchmarks a complex multi-figure scenario
func BenchmarkComplexScenario_KingAndRook(b *testing.B) {
	placement := &Placement{}
	placement.SetDimension(8)
	king := &King{}
	rook := &Rook{}
	king.SetNext(rook)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boards := make(map[uint64][]byte)
		boards = placement.PlaceFigures(1, king, boards)
		boards = placement.PlaceFigures(1, rook, boards)
		PutMapToPool(boards)
	}
}
