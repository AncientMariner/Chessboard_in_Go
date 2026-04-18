package main

import (
	"reflect"
	"testing"

	"Chessboard_in_Go/figures"
)

func TestNewChessBoard(t *testing.T) {
	tests := []struct {
		name string
		want ChessboardBuilder
	}{
		{"Test empty chessboard", &boardBuilder{chessboard: &Chessboard{}, figureQuantityMap: make(map[byte]int)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChessboard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChessboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewChessBoardWithSize(t *testing.T) {
	tests := []struct {
		name string
		want ChessboardBuilder
	}{
		{"Test empty chessboard size 10", &boardBuilder{chessboard: &Chessboard{}, figureQuantityMap: make(map[byte]int)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChessboardWithSize(10); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChessboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardBuilder_Build(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[byte]int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Chessboard
	}{
		{"Test build", fields{
			chessboard:             &Chessboard{},
			currentFigureBehaviour: nil,
			figureQuantityMap:      make(map[byte]int),
		}, &Chessboard{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &boardBuilder{
				chessboard:             tt.fields.chessboard,
				currentFigureBehaviour: tt.fields.currentFigureBehaviour,
				figureQuantityMap:      tt.fields.figureQuantityMap,
			}
			if got := b.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardBuilder_addToEmptyChain(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[byte]int
	}
	type args struct {
		figure figures.FigureBehaviour
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ChessboardBuilder
	}{
		{"Test add to empty chain", fields{
			&Chessboard{
				nil,
				nil,
				figures.Placement{},
			},
			nil,
			nil,
		}, args{figure: &figures.Queen{}},
			&boardBuilder{
				&Chessboard{
					nil,
					&figures.Queen{},
					figures.Placement{},
				},
				&figures.Queen{},
				nil,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &boardBuilder{
				tt.fields.chessboard,
				tt.fields.currentFigureBehaviour,
				tt.fields.figureQuantityMap,
			}

			if got := b.addToChain(tt.args.figure); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToEmptyChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardBuilder_addToNonEmptyChain(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[byte]int
	}
	type args struct {
		figure figures.FigureBehaviour
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ChessboardBuilder
	}{
		{"Test add to chain", fields{
			&Chessboard{
				map[byte]int{'k': 1},
				&figures.King{},
				figures.Placement{},
			},
			&figures.Queen{},
			map[byte]int{'q': 1},
		}, args{figure: &figures.Queen{}},
			&boardBuilder{
				&Chessboard{
					map[byte]int{'k': 1},
					&figures.King{},
					figures.Placement{},
				},
				&figures.Queen{},
				map[byte]int{'q': 1},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &boardBuilder{
				tt.fields.chessboard,
				tt.fields.currentFigureBehaviour,
				tt.fields.figureQuantityMap,
			}
			if got := b.addToChain(tt.args.figure); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToNonEmptyChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_different_combinations(t *testing.T) {
	type args struct {
		board *Chessboard
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test empty board with 1 king", args{board: NewChessboard().withKing(1).Build()}, 64},
		// {"Test empty board with 16 king", args{board: NewChessboard().withKing(16).Build()}, 1806},
		{"Test empty board with 2 king", args{board: NewChessboard().withKing(2).Build()}, 1806},                    // 3612 if black and white king (distinguishable)
		{"Test empty board with 1 king 1 king", args{board: NewChessboard().withKing(1).withKing(1).Build()}, 1806}, // Now fixed - no longer adds duplicate handler, same as withKing(2)
		{"Test empty board with 2 rook", args{board: NewChessboard().withRook(2).Build()}, 1568},
		{"Test empty board with 1 rook 1 rook", args{board: NewChessboard().withRook(1).withRook(1).Build()}, 1568}, // Now fixed - no longer adds duplicate handler, same as withRook(2)
		{"Test empty board with 1 king 1 rook", args{board: NewChessboard().withKing(1).withRook(1).Build()}, 2940},
		{"Test empty board with 1 rook 1 king", args{board: NewChessboard().withRook(1).withKing(1).Build()}, 2940},
		{"Test empty board with 1 rook 1 king", args{board: NewChessboard().withRook(2).withKing(2).Build()}, 657390},
		// {"Test empty board with 8 rook", args{board: NewChessboard().withRook(8).Build()}, 40320},
		{"Test empty board with 1 king 1 bishop", args{board: NewChessboard().withKing(1).withBishop(1).Build()}, 3248},
		{"Test empty board with 2 king 2 bishop", args{board: NewChessboard().withKing(2).withBishop(2).Build()}, 1177824},
		{"Test empty board with 1 bishop 1 king", args{board: NewChessboard().withBishop(1).withKing(1).Build()}, 3248},
		// {"Test empty board with 14 bishop", args{board: NewChessboard().withBishop(14).Build()}, 1736},
		{"Test empty board with 2 bishop ", args{board: NewChessboard().withBishop(2).Build()}, 1736},
		{"Test empty board with 1 bishop 1 bishop", args{board: NewChessboard().withBishop(1).withBishop(1).Build()}, 1736}, // Now fixed - no longer adds duplicate handler, same as withBishop(2)
		{"Test empty board with 1 bishop 1 rook", args{board: NewChessboard().withBishop(1).withRook(1).Build()}, 2576},
		{"Test empty board with 1 rook 1 bishop", args{board: NewChessboard().withRook(1).withBishop(1).Build()}, 2576},
		// {"Test empty board with 32 knight", args{board: NewChessboard().withKnight(32).Build()}, 3063828},
		{"Test empty board with 1 knight 1 knight", args{board: NewChessboard().withKnight(1).withKnight(1).Build()}, 1848}, // Now fixed - no longer adds duplicate handler, same as withKnight(2)
		{"Test empty board with 1 king 1 knight", args{board: NewChessboard().withKing(1).withKnight(1).Build()}, 3276},
		{"Test empty board with 1 knight 1 king", args{board: NewChessboard().withKnight(1).withKing(1).Build()}, 3276},
		{"Test empty board with 1 queen 1 knight", args{board: NewChessboard().withQueen(1).withKnight(1).Build()}, 2240},
		{"Test empty board with 1 knight 1 queen", args{board: NewChessboard().withKnight(1).withQueen(1).Build()}, 2240},
		{"Test empty board with 1 queen 1 bishop", args{board: NewChessboard().withQueen(1).withBishop(1).Build()}, 2576},
		{"Test empty board with 1 bishop 1 queen", args{board: NewChessboard().withBishop(1).withQueen(1).Build()}, 2576},
		{"Test empty board with 1 rook 1 knight", args{board: NewChessboard().withRook(1).withKnight(1).Build()}, 2800},
		{"Test empty board with 1 knight 1 rook", args{board: NewChessboard().withKnight(1).withRook(1).Build()}, 2800},
		{"Test empty board with 1 queen 1 rook", args{board: NewChessboard().withQueen(1).withRook(1).Build()}, 2576},
		{"Test empty board with 1 rook 1 queen", args{board: NewChessboard().withRook(1).withQueen(1).Build()}, 2576},
		{"Test empty board with 1 bishop 1 knight", args{board: NewChessboard().withBishop(1).withKnight(1).Build()}, 3136},
		{"Test empty board with 1 knight 1 bishop", args{board: NewChessboard().withKnight(1).withBishop(1).Build()}, 3136},
		// {"Test empty board with 16 king 14 bishop", args{board: NewChessboard().withKing(16).withBishop(14).Build()}, 3063828},
		{"Test empty board with 1 king 1 queen", args{board: NewChessboard().withKing(1).withQueen(1).Build()}, 2576},
		{"Test empty board with 1 queen 1 king", args{board: NewChessboard().withQueen(1).withKing(1).Build()}, 2576},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.board.calculateBoards(); len(got) != tt.want {
				t.Errorf("calculateBoards() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_queen_combinations(t *testing.T) {
	type args struct {
		board *Chessboard
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test empty board with 6 queen", args{board: NewChessboard().withQueen(6).Build()}, 22708},
		{"Test empty board with 7 queen", args{board: NewChessboard().withQueen(7).Build()}, 3192},
		{"Test empty board with 8 queen", args{board: NewChessboard().withQueen(8).Build()}, 92},
		{"Test empty board with 8 queen", args{board: NewChessboard().withQueen(1).withQueen(1).withQueen(1).withQueen(1).withQueen(1).withQueen(1).withQueen(1).withQueen(1).Build()}, 92}, // Now fixed - no longer adds duplicate handlers, same as withQueen(8)
		// {"Test empty board with 9 queen, impossible case", args{board: NewChessboard().withQueen(9).Build()}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.board.calculateBoards(); len(got) != tt.want {
				t.Errorf("calculateBoards() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_different_combinations_long_running(t *testing.T) {
	type args struct {
		board *Chessboard
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"Test empty board with 16 king", args{board: NewChessboard().withKing(16).Build()}, 1806},
		{"Test empty board with 8 rook", args{board: NewChessboard().withRook(8).Build()}, 40320},
		// {"Test empty board with 14 bishop", args{board: NewChessboard().withBishop(14).Build()}, 1736},
		// {"Test empty board with 32 knight", args{board: NewChessboard().withKnight(32).Build()}, 3063828},
		// {"Test empty board with 16 king 14 bishop", args{board: NewChessboard().withKing(16).withBishop(14).Build()}, 3063828},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.board.calculateBoards(); len(got) != tt.want {
				t.Errorf("calculateBoards() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

// When all the tests are run in parallel, and there are different board sizes, a race condition happens
// func Test_number_of_boards_with_1_figure_7x7(t *testing.T) {
// 	type args struct {
// 		board *Chessboard
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"Test empty board 7x7 with 1 king", args{board: NewChessboardWithSize(7).withKing(1).Build()}, 49},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.args.board.calculateBoards(); len(got) != tt.want {
// 				t.Errorf("calculateBoards() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_number_of_boards_with_1_figure_7x7_long(t *testing.T) {
// 	type args struct {
// 		board *Chessboard
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"Test empty 7x7 board with 2 king 2 queen 2 bishop 1 knight", args{board: NewChessboardWithSize(7).withKing(2).withQueen(2).withBishop(2).withKnight(1).Build()}, 3761852},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.args.board.calculateBoards(); len(got) != tt.want {
// 				t.Errorf("calculateBoards() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_number_of_boards_with_different_figure_variations(t *testing.T) {
	// After fixing the king placement bug and duplicate handler bug, all three approaches now produce identical results
	calculateBoards_R_K := NewChessboard().withRook(2).withKing(1).Build().calculateBoards()
	calculateBoards_R_K_alt := NewChessboard().withRook(1).withRook(1).withKing(1).Build().calculateBoards()
	calculateBoards_K_R := NewChessboard().withKing(1).withRook(1).withRook(1).Build().calculateBoards()

	var unitedSet = make(map[uint64][]byte, len(calculateBoards_R_K)+len(calculateBoards_R_K_alt)+len(calculateBoards_K_R))

	for u, s := range calculateBoards_R_K {
		unitedSet[u] = s
	}

	for u, s := range calculateBoards_R_K_alt {
		unitedSet[u] = s
	}

	for u, s := range calculateBoards_K_R {
		unitedSet[u] = s
	}

	// After fixing both bugs, all three approaches now produce identical results
	if len(calculateBoards_R_K) != 49464 {
		t.Errorf("calculateBoards() R(2)-K(1) = %v, want %v", len(calculateBoards_R_K), 49464)
	}
	if len(calculateBoards_R_K_alt) != 49464 {
		t.Errorf("calculateBoards() R(1)-R(1)-K(1) = %v, want %v (now fixed - no longer adds duplicate handler)", len(calculateBoards_R_K_alt), 49464)
	}
	if len(calculateBoards_K_R) != 49464 {
		t.Errorf("calculateBoards() K(1)-R(2) = %v, want %v", len(calculateBoards_K_R), 49464)
	}
	// United set should contain 49464 unique boards since all three approaches produce identical results
	if len(unitedSet) != 49464 {
		t.Errorf("calculateBoards() united set = %v, want %v", len(unitedSet), 49464)
	}
}

// Test_placement_order_independence_for_all_figure_pairs verifies that placing figures in different orders
// produces identical results for all 10 possible pairs of different figure types.
// This ensures the placement algorithm is truly order-independent after bug fixes.
func Test_placement_order_independence_for_all_figure_pairs(t *testing.T) {
	testCases := []struct {
		name          string
		figure1Name   string
		figure2Name   string
		builder1      func() *Chessboard // figure1 then figure2
		builder2      func() *Chessboard // figure2 then figure1
		expectedCount int
	}{
		{
			name:          "King-Rook order independence",
			figure1Name:   "King",
			figure2Name:   "Rook",
			builder1:      func() *Chessboard { return NewChessboard().withKing(1).withRook(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withRook(1).withKing(1).Build() },
			expectedCount: 2940,
		},
		{
			name:          "King-Bishop order independence",
			figure1Name:   "King",
			figure2Name:   "Bishop",
			builder1:      func() *Chessboard { return NewChessboard().withKing(1).withBishop(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withBishop(1).withKing(1).Build() },
			expectedCount: 3248,
		},
		{
			name:          "King-Knight order independence",
			figure1Name:   "King",
			figure2Name:   "Knight",
			builder1:      func() *Chessboard { return NewChessboard().withKing(1).withKnight(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withKnight(1).withKing(1).Build() },
			expectedCount: 3276,
		},
		{
			name:          "King-Queen order independence",
			figure1Name:   "King",
			figure2Name:   "Queen",
			builder1:      func() *Chessboard { return NewChessboard().withKing(1).withQueen(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withQueen(1).withKing(1).Build() },
			expectedCount: 2576,
		},
		{
			name:          "Rook-Bishop order independence",
			figure1Name:   "Rook",
			figure2Name:   "Bishop",
			builder1:      func() *Chessboard { return NewChessboard().withRook(1).withBishop(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withBishop(1).withRook(1).Build() },
			expectedCount: 2576,
		},
		{
			name:          "Rook-Knight order independence",
			figure1Name:   "Rook",
			figure2Name:   "Knight",
			builder1:      func() *Chessboard { return NewChessboard().withRook(1).withKnight(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withKnight(1).withRook(1).Build() },
			expectedCount: 2800,
		},
		{
			name:          "Rook-Queen order independence",
			figure1Name:   "Rook",
			figure2Name:   "Queen",
			builder1:      func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withQueen(1).withRook(1).Build() },
			expectedCount: 2576,
		},
		{
			name:          "Bishop-Knight order independence",
			figure1Name:   "Bishop",
			figure2Name:   "Knight",
			builder1:      func() *Chessboard { return NewChessboard().withBishop(1).withKnight(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withKnight(1).withBishop(1).Build() },
			expectedCount: 3136,
		},
		{
			name:          "Bishop-Queen order independence",
			figure1Name:   "Bishop",
			figure2Name:   "Queen",
			builder1:      func() *Chessboard { return NewChessboard().withBishop(1).withQueen(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withQueen(1).withBishop(1).Build() },
			expectedCount: 2576,
		},
		{
			name:          "Knight-Queen order independence",
			figure1Name:   "Knight",
			figure2Name:   "Queen",
			builder1:      func() *Chessboard { return NewChessboard().withKnight(1).withQueen(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withQueen(1).withKnight(1).Build() },
			expectedCount: 2240,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			boards1 := tc.builder1().calculateBoards()
			boards2 := tc.builder2().calculateBoards()

			// Check that both orders produce the same count
			if len(boards1) != tc.expectedCount {
				t.Errorf("%s-%s: got %d boards, want %d", tc.figure1Name, tc.figure2Name, len(boards1), tc.expectedCount)
			}
			if len(boards2) != tc.expectedCount {
				t.Errorf("%s-%s: got %d boards, want %d", tc.figure2Name, tc.figure1Name, len(boards2), tc.expectedCount)
			}

			// Check that both orders produce identical board configurations
			if len(boards1) != len(boards2) {
				t.Errorf("Order matters! %s-%s produced %d boards, but %s-%s produced %d boards",
					tc.figure1Name, tc.figure2Name, len(boards1),
					tc.figure2Name, tc.figure1Name, len(boards2))
			}

			// Verify that the actual board configurations are identical (not just the count)
			for hash, board1 := range boards1 {
				board2, exists := boards2[hash]
				if !exists {
					t.Errorf("Board configuration %v exists in %s-%s but not in %s-%s",
						hash, tc.figure1Name, tc.figure2Name, tc.figure2Name, tc.figure1Name)
					break
				}
				// Compare board contents
				if len(board1) != len(board2) {
					t.Errorf("Board %v has different contents: %d vs %d bytes",
						hash, len(board1), len(board2))
					break
				}
				for i := range board1 {
					if board1[i] != board2[i] {
						t.Errorf("Board %v differs at position %d: %v vs %v",
							hash, i, board1[i], board2[i])
						break
					}
				}
			}

			// Check for boards in boards2 that aren't in boards1
			for hash := range boards2 {
				if _, exists := boards1[hash]; !exists {
					t.Errorf("Board configuration %v exists in %s-%s but not in %s-%s",
						hash, tc.figure2Name, tc.figure1Name, tc.figure1Name, tc.figure2Name)
					break
				}
			}
		})
	}
}

// Test_placement_order_independence_for_three_figures verifies that placing 3 figures in different orders
// produces identical results. Tests all meaningful 3-figure combinations (not all 3 same).
// Covers: 1 each of 3 different types + 2 of one type + 1 of another.
func Test_placement_order_independence_for_three_figures(t *testing.T) {
	t.Skip("Long-running test for 3-figure combinations in all orders. Run separately when needed.")
	testCases := []struct {
		name          string
		description   string
		builders      []func() *Chessboard // All permutations to test
		expectedCount int
	}{
		// ========== 1 of each of 3 different types (10 combinations) ==========
		{
			name:        "K(1)-R(1)-B(1) all orders",
			description: "1 King + 1 Rook + 1 Bishop in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withBishop(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withBishop(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKing(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withRook(1).withKing(1).Build() },
			},
			expectedCount: 92768,
		},
		{
			name:        "K(1)-R(1)-N(1) all orders",
			description: "1 King + 1 Rook + 1 Knight in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withKnight(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKnight(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withKing(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withRook(1).withKing(1).Build() },
			},
			expectedCount: 100552,
		},
		{
			name:        "K(1)-R(1)-Q(1) all orders",
			description: "1 King + 1 Rook + 1 Queen in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withQueen(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKing(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withRook(1).withKing(1).Build() },
			},
			expectedCount: 71320,
		},
		{
			name:        "K(1)-B(1)-N(1) all orders",
			description: "1 King + 1 Bishop + 1 Knight in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withBishop(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withKnight(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKing(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKnight(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withKing(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withBishop(1).withKing(1).Build() },
			},
			expectedCount: 126920,
		},
		{
			name:        "K(1)-B(1)-Q(1) all orders",
			description: "1 King + 1 Bishop + 1 Queen in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withBishop(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withQueen(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKing(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withQueen(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKing(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withBishop(1).withKing(1).Build() },
			},
			expectedCount: 80560,
		},
		{
			name:        "K(1)-N(1)-Q(1) all orders",
			description: "1 King + 1 Knight + 1 Queen in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withKnight(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withQueen(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withKing(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withQueen(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKing(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKnight(1).withKing(1).Build() },
			},
			expectedCount: 70384,
		},
		{
			name:        "R(1)-B(1)-N(1) all orders",
			description: "1 Rook + 1 Bishop + 1 Knight in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(1).withBishop(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKnight(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withRook(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKnight(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withRook(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withBishop(1).withRook(1).Build() },
			},
			expectedCount: 87056,
		},
		{
			name:        "R(1)-B(1)-Q(1) all orders",
			description: "1 Rook + 1 Bishop + 1 Queen in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(1).withBishop(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withRook(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withQueen(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withRook(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withBishop(1).withRook(1).Build() },
			},
			expectedCount: 61920,
		},
		{
			name:        "R(1)-N(1)-Q(1) all orders",
			description: "1 Rook + 1 Knight + 1 Queen in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(1).withKnight(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withRook(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withQueen(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withRook(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKnight(1).withRook(1).Build() },
			},
			expectedCount: 59032,
		},
		{
			name:        "B(1)-N(1)-Q(1) all orders",
			description: "1 Bishop + 1 Knight + 1 Queen in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withBishop(1).withKnight(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withQueen(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withBishop(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withQueen(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withBishop(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKnight(1).withBishop(1).Build() },
			},
			expectedCount: 68344,
		},

		// ========== 2 of one type + 1 of another (20 combinations) ==========
		{
			name:        "K(2)-R(1) both orders",
			description: "2 Kings + 1 Rook in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(2).Build() },
			},
			expectedCount: 58344,
		},
		{
			name:        "K(2)-B(1) both orders",
			description: "2 Kings + 1 Bishop in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKing(2).Build() },
			},
			expectedCount: 71952,
		},
		{
			name:        "K(2)-N(1) both orders",
			description: "2 Kings + 1 Knight in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withKing(2).Build() },
			},
			expectedCount: 73080,
		},
		{
			name:        "K(2)-Q(1) both orders",
			description: "2 Kings + 1 Queen in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKing(2).Build() },
			},
			expectedCount: 44980,
		},
		{
			name:        "R(2)-K(1) both orders",
			description: "2 Rooks + 1 King in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(2).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withRook(2).Build() },
			},
			expectedCount: 49464,
		},
		{
			name:        "R(2)-B(1) both orders",
			description: "2 Rooks + 1 Bishop in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(2).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withRook(2).Build() },
			},
			expectedCount: 38264,
		},
		{
			name:        "R(2)-N(1) both orders",
			description: "2 Rooks + 1 Knight in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(2).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withRook(2).Build() },
			},
			expectedCount: 44932,
		},
		{
			name:        "R(2)-Q(1) both orders",
			description: "2 Rooks + 1 Queen in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(2).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withRook(2).Build() },
			},
			expectedCount: 38264,
		},
		{
			name:        "B(2)-K(1) both orders",
			description: "2 Bishops + 1 King in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withBishop(2).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withBishop(2).Build() },
			},
			expectedCount: 68816,
		},
		{
			name:        "B(2)-R(1) both orders",
			description: "2 Bishops + 1 Rook in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withBishop(2).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withBishop(2).Build() },
			},
			expectedCount: 43360,
		},
		{
			name:        "B(2)-N(1) both orders",
			description: "2 Bishops + 1 Knight in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withBishop(2).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withBishop(2).Build() },
			},
			expectedCount: 64640,
		},
		{
			name:        "B(2)-Q(1) both orders",
			description: "2 Bishops + 1 Queen in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withBishop(2).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withBishop(2).Build() },
			},
			expectedCount: 43360,
		},
		{
			name:        "N(2)-K(1) both orders",
			description: "2 Knights + 1 King in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKnight(2).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withKnight(2).Build() },
			},
			expectedCount: 75616,
		},
		{
			name:        "N(2)-R(1) both orders",
			description: "2 Knights + 1 Rook in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKnight(2).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKnight(2).Build() },
			},
			expectedCount: 55084,
		},
		{
			name:        "N(2)-B(1) both orders",
			description: "2 Knights + 1 Bishop in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKnight(2).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKnight(2).Build() },
			},
			expectedCount: 69628,
		},
		{
			name:        "N(2)-Q(1) both orders",
			description: "2 Knights + 1 Queen in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKnight(2).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKnight(2).Build() },
			},
			expectedCount: 35156,
		},
		{
			name:        "Q(2)-K(1) both orders",
			description: "2 Queens + 1 King in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withQueen(2).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withQueen(2).Build() },
			},
			expectedCount: 30960,
		},
		{
			name:        "Q(2)-R(1) both orders",
			description: "2 Queens + 1 Rook in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withQueen(2).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(2).Build() },
			},
			expectedCount: 30960,
		},
		{
			name:        "Q(2)-B(1) both orders",
			description: "2 Queens + 1 Bishop in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withQueen(2).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withQueen(2).Build() },
			},
			expectedCount: 30960,
		},
		{
			name:        "Q(2)-N(1) both orders",
			description: "2 Queens + 1 Knight in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withQueen(2).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withQueen(2).Build() },
			},
			expectedCount: 23216,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Calculate boards for all permutations
			var allResults []map[uint64][]byte
			for i, builder := range tc.builders {
				boards := builder().calculateBoards()
				allResults = append(allResults, boards)

				// Check that each permutation produces the expected count
				if len(boards) != tc.expectedCount {
					t.Errorf("Permutation %d: got %d boards, want %d", i+1, len(boards), tc.expectedCount)
				}
			}

			// Verify all permutations produce identical results
			if len(allResults) > 1 {
				firstResult := allResults[0]
				for i := 1; i < len(allResults); i++ {
					if len(allResults[i]) != len(firstResult) {
						t.Errorf("Permutation %d count differs: got %d, want %d",
							i+1, len(allResults[i]), len(firstResult))
					}

					// Check that board configurations are identical
					for hash, board1 := range firstResult {
						board2, exists := allResults[i][hash]
						if !exists {
							t.Errorf("Permutation %d missing board hash %v", i+1, hash)
							break
						}
						if len(board1) != len(board2) {
							t.Errorf("Permutation %d board %v length differs: %d vs %d",
								i+1, hash, len(board1), len(board2))
							break
						}
						for j := range board1 {
							if board1[j] != board2[j] {
								t.Errorf("Permutation %d board %v differs at position %d",
									i+1, hash, j)
								break
							}
						}
					}

					// Check reverse: boards in result[i] but not in firstResult
					for hash := range allResults[i] {
						if _, exists := firstResult[hash]; !exists {
							t.Errorf("Permutation %d has extra board hash %v", i+1, hash)
							break
						}
					}
				}
			}
		})
	}
}

// func contains(s []string, e string) bool {
// 	for _, a := range s {
// 		if a == e {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// // difference returns the elements in `a` that aren't in `b`.
// func difference(a, b []string) []string {
// 	mb := make(map[string]struct{}, len(b))
// 	for _, x := range b {
// 		mb[x] = struct{}{}
// 	}
// 	var diff []string
// 	for _, x := range a {
// 		if _, found := mb[x]; !found {
// 			diff = append(diff, x)
// 		}
// 	}
// 	return diff
// }

func Test_board_with_1_figure(t *testing.T) {
	type fields struct {
		figureQuantityMap      map[byte]int
		currentFigureBehaviour figures.FigureBehaviour
		figurePlacement        figures.Placement
	}
	type args struct {
		behaviour            figures.FigureBehaviour
		previousFigureBoards map[uint64][]byte
	}
	behaviour := &figures.King{}
	behaviour.SetNext(&figures.Queen{})

	figureBehaviour := &figures.King{}
	figureBehaviour.SetNext(&figures.Rook{})

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test empty board with 1 king", fields{map[byte]int{(&figures.King{}).GetName(): 1}, &figures.King{}, figures.Placement{}}, args{&figures.King{}, make(map[uint64][]byte)}, 64},
		{"Test empty board with 2 king", fields{map[byte]int{(&figures.King{}).GetName(): 2}, &figures.King{}, figures.Placement{}}, args{&figures.King{}, make(map[uint64][]byte)}, 1806},
		{"Test empty board with 1 king 1 rook", fields{map[byte]int{(&figures.King{}).GetName(): 1, (&figures.Rook{}).GetName(): 1}, figureBehaviour, figures.Placement{}}, args{figureBehaviour, make(map[uint64][]byte)}, 2940},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &Chessboard{
				figureQuantityMap:      tt.fields.figureQuantityMap,
				currentFigureBehaviour: tt.fields.currentFigureBehaviour,
				figurePlacement:        tt.fields.figurePlacement,
			}
			if got := board.calculateBoard(tt.args.behaviour, tt.args.previousFigureBoards); len(got) != tt.want {
				t.Errorf("calculateBoard() = %d, want %d", len(got), tt.want)
			}
		})
	}
}

// Test_placement_order_independence_for_four_plus_figures verifies that placing 4 or more figures
// in different orders produces identical results. This tests complex chain interactions.
// Note: Tests a sample of permutations due to performance (4! = 24, 5! = 120 permutations possible)
func Test_placement_order_independence_for_four_plus_figures(t *testing.T) {
	t.Skip("Long-running test for 4+ figure combinations in multiple orders. Run separately when needed.")
	testCases := []struct {
		name          string
		description   string
		builders      []func() *Chessboard
		expectedCount int
	}{
		// ========== 4 Figures: 1 of each of 4 different types (5 combinations) ==========
		{
			name:        "K(1)-R(1)-B(1)-N(1) sample permutations",
			description: "1 King + 1 Rook + 1 Bishop + 1 Knight - testing 4 key permutations",
			builders: []func() *Chessboard{
				// Test 4 representative permutations to verify order independence
				func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withBishop(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withBishop(1).withRook(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(1).withBishop(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKnight(1).withKing(1).withRook(1).Build() },
			},
			expectedCount: 2444016,
		},
		{
			name:        "K(1)-R(1)-B(1)-Q(1) sample permutations",
			description: "1 King + 1 Rook + 1 Bishop + 1 Queen - testing 4 key permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withBishop(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withBishop(1).withRook(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).withKing(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKing(1).withQueen(1).withRook(1).Build() },
			},
			expectedCount: 1309152,
		},
		{
			name:        "K(1)-R(1)-N(1)-Q(1) sample permutations",
			description: "1 King + 1 Rook + 1 Knight + 1 Queen - testing 4 key permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withKnight(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKnight(1).withRook(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).withKing(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withKing(1).withQueen(1).withRook(1).Build() },
			},
			expectedCount: 1237560,
		},
		{
			name:        "K(1)-B(1)-N(1)-Q(1) sample permutations",
			description: "1 King + 1 Bishop + 1 Knight + 1 Queen - testing 4 key permutations",
			builders: []func() *Chessboard{
				func() *Chessboard {
					return NewChessboard().withKing(1).withBishop(1).withKnight(1).withQueen(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withQueen(1).withKnight(1).withBishop(1).withKing(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withBishop(1).withQueen(1).withKing(1).withKnight(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withKnight(1).withKing(1).withQueen(1).withBishop(1).Build()
				},
			},
			expectedCount: 1657656,
		},
		{
			name:        "R(1)-B(1)-N(1)-Q(1) sample permutations",
			description: "1 Rook + 1 Bishop + 1 Knight + 1 Queen - testing 4 key permutations",
			builders: []func() *Chessboard{
				func() *Chessboard {
					return NewChessboard().withRook(1).withBishop(1).withKnight(1).withQueen(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withQueen(1).withKnight(1).withBishop(1).withRook(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withBishop(1).withQueen(1).withRook(1).withKnight(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withKnight(1).withRook(1).withQueen(1).withBishop(1).Build()
				},
			},
			expectedCount: 1082800,
		},
		{
			name:        "K(1)-R(1)-B(1)-Q(1) sample permutations",
			description: "1 King + 1 Rook + 1 Bishop + 1 Queen - testing 8 of 24 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withBishop(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withBishop(1).withRook(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).withKing(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKing(1).withQueen(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withBishop(1).withRook(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(1).withQueen(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withRook(1).withBishop(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withQueen(1).withKing(1).withRook(1).Build() },
			},
			expectedCount: 1309152,
		},
		{
			name:        "K(1)-R(1)-N(1)-Q(1) sample permutations",
			description: "1 King + 1 Rook + 1 Knight + 1 Queen - testing 6 of 24 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withKnight(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKnight(1).withRook(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).withKing(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withKing(1).withQueen(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withKnight(1).withRook(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(1).withQueen(1).withKnight(1).Build() },
			},
			expectedCount: 1237560,
		},
		{
			name:        "K(1)-B(1)-N(1)-Q(1) sample permutations",
			description: "1 King + 1 Bishop + 1 Knight + 1 Queen - testing 6 of 24 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard {
					return NewChessboard().withKing(1).withBishop(1).withKnight(1).withQueen(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withQueen(1).withKnight(1).withBishop(1).withKing(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withBishop(1).withQueen(1).withKing(1).withKnight(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withKnight(1).withKing(1).withQueen(1).withBishop(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withKing(1).withKnight(1).withBishop(1).withQueen(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withBishop(1).withKing(1).withQueen(1).withKnight(1).Build()
				},
			},
			expectedCount: 1657656,
		},
		{
			name:        "R(1)-B(1)-N(1)-Q(1) sample permutations",
			description: "1 Rook + 1 Bishop + 1 Knight + 1 Queen - testing 6 of 24 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard {
					return NewChessboard().withRook(1).withBishop(1).withKnight(1).withQueen(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withQueen(1).withKnight(1).withBishop(1).withRook(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withBishop(1).withQueen(1).withRook(1).withKnight(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withKnight(1).withRook(1).withQueen(1).withBishop(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withRook(1).withKnight(1).withBishop(1).withQueen(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withBishop(1).withRook(1).withQueen(1).withKnight(1).Build()
				},
			},
			expectedCount: 1082800,
		},

		// ========== 4 Figures: 2 of one + 1 of each of 2 others ==========
		{
			name:        "K(2)-R(1)-B(1) all permutations",
			description: "2 Kings + 1 Rook + 1 Bishop in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withRook(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(2).withBishop(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(2).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withBishop(1).withKing(2).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKing(2).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withRook(1).withKing(2).Build() },
			},
			expectedCount: 1422376,
		},
		{
			name:        "K(2)-R(1)-N(1) all permutations",
			description: "2 Kings + 1 Rook + 1 Knight in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withRook(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(2).withKnight(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(2).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKnight(1).withKing(2).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withKing(2).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withRook(1).withKing(2).Build() },
			},
			expectedCount: 1543368,
		},
		{
			name:        "K(2)-R(1)-Q(1) all permutations",
			description: "2 Kings + 1 Rook + 1 Queen in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withRook(1).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(2).withQueen(1).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withKing(2).withQueen(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(1).withQueen(1).withKing(2).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withKing(2).withRook(1).Build() },
				func() *Chessboard { return NewChessboard().withQueen(1).withRook(1).withKing(2).Build() },
			},
			expectedCount: 833568,
		},
		{
			name:        "R(2)-K(1)-B(1) all permutations",
			description: "2 Rooks + 1 King + 1 Bishop in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(2).withKing(1).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(2).withBishop(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withRook(2).withBishop(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withBishop(1).withRook(2).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withRook(2).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withBishop(1).withKing(1).withRook(2).Build() },
			},
			expectedCount: 930496,
		},
		{
			name:        "R(2)-K(1)-N(1) all permutations",
			description: "2 Rooks + 1 King + 1 Knight in all 6 permutations",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(2).withKing(1).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withRook(2).withKnight(1).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withRook(2).withKnight(1).Build() },
				func() *Chessboard { return NewChessboard().withKing(1).withKnight(1).withRook(2).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withRook(2).withKing(1).Build() },
				func() *Chessboard { return NewChessboard().withKnight(1).withKing(1).withRook(2).Build() },
			},
			expectedCount: 1079012,
		},

		// ========== 4 Figures: 2 of one + 2 of another ==========
		{
			name:        "K(2)-R(2) both orders",
			description: "2 Kings + 2 Rooks in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withRook(2).Build() },
				func() *Chessboard { return NewChessboard().withRook(2).withKing(2).Build() },
			},
			expectedCount: 657390,
		},
		{
			name:        "K(2)-B(2) both orders",
			description: "2 Kings + 2 Bishops in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withKing(2).withBishop(2).Build() },
				func() *Chessboard { return NewChessboard().withBishop(2).withKing(2).Build() },
			},
			expectedCount: 1177824,
		},
		{
			name:        "R(2)-B(2) both orders",
			description: "2 Rooks + 2 Bishops in both orders",
			builders: []func() *Chessboard{
				func() *Chessboard { return NewChessboard().withRook(2).withBishop(2).Build() },
				func() *Chessboard { return NewChessboard().withBishop(2).withRook(2).Build() },
			},
			expectedCount: 384084,
		},

		// ========== 5 Figures: 1 of each type ==========
		{
			name:        "K(1)-R(1)-B(1)-N(1)-Q(1) sample permutations",
			description: "All 5 figure types - testing 4 key permutations",
			builders: []func() *Chessboard{
				func() *Chessboard {
					return NewChessboard().withKing(1).withRook(1).withBishop(1).withKnight(1).withQueen(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withQueen(1).withKnight(1).withBishop(1).withRook(1).withKing(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withBishop(1).withKing(1).withQueen(1).withRook(1).withKnight(1).Build()
				},
				func() *Chessboard {
					return NewChessboard().withKnight(1).withBishop(1).withKing(1).withQueen(1).withRook(1).Build()
				},
			},
			expectedCount: 17265360,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var allResults []map[uint64][]byte
			for i, builder := range tc.builders {
				boards := builder().calculateBoards()
				allResults = append(allResults, boards)

				if len(boards) != tc.expectedCount {
					t.Errorf("Permutation %d: got %d boards, want %d", i+1, len(boards), tc.expectedCount)
				}
			}

			// Verify all permutations produce identical results
			if len(allResults) > 1 {
				firstResult := allResults[0]
				for i := 1; i < len(allResults); i++ {
					if len(allResults[i]) != len(firstResult) {
						t.Errorf("Permutation %d count differs: got %d, want %d",
							i+1, len(allResults[i]), len(firstResult))
					}

					// Check board hash identity (spot check - full comparison is too slow)
					checkedCount := 0
					for hash := range firstResult {
						if _, exists := allResults[i][hash]; !exists {
							t.Errorf("Permutation %d missing board hash %v", i+1, hash)
							break
						}
						checkedCount++
						if checkedCount >= 100 { // Spot check first 100 boards for performance
							break
						}
					}
				}
			}
		})
	}
}

// Test_mixed_quantities_same_figure_type verifies that calling withX() multiple times
// produces the same result as calling withX(n) once. Tests the duplicate handler fix.
func Test_mixed_quantities_same_figure_type(t *testing.T) {
	testCases := []struct {
		name          string
		description   string
		builder1      func() *Chessboard // Mixed calls: withX(1).withX(1)...
		builder2      func() *Chessboard // Single call: withX(n)
		expectedCount int
	}{
		{
			name:          "K(1)+K(1)+K(1) equals K(3)",
			description:   "3 separate King(1) calls should equal King(3)",
			builder1:      func() *Chessboard { return NewChessboard().withKing(1).withKing(1).withKing(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withKing(3).Build() },
			expectedCount: 29708,
		},
		{
			name:          "R(2)+R(1) equals R(3)",
			description:   "Rook(2) then Rook(1) should equal Rook(3)",
			builder1:      func() *Chessboard { return NewChessboard().withRook(2).withRook(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withRook(3).Build() },
			expectedCount: 18816,
		},
		{
			name:          "R(1)+R(2) equals R(3)",
			description:   "Rook(1) then Rook(2) should equal Rook(3) (order reversed)",
			builder1:      func() *Chessboard { return NewChessboard().withRook(1).withRook(2).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withRook(3).Build() },
			expectedCount: 18816,
		},
		{
			name:        "N(1)+N(1)+N(1)+N(1) equals N(4)",
			description: "4 separate Knight(1) calls should equal Knight(4)",
			builder1: func() *Chessboard {
				return NewChessboard().withKnight(1).withKnight(1).withKnight(1).withKnight(1).Build()
			},
			builder2:      func() *Chessboard { return NewChessboard().withKnight(4).Build() },
			expectedCount: 376560,
		},
		{
			name:          "Q(4)+Q(4) equals Q(8)",
			description:   "Queen(4) then Queen(4) should equal Queen(8)",
			builder1:      func() *Chessboard { return NewChessboard().withQueen(4).withQueen(4).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withQueen(8).Build() },
			expectedCount: 92,
		},
		{
			name:          "K(1)+R(1)+K(1) equals K(2)+R(1)",
			description:   "Mixed King calls with Rook in between",
			builder1:      func() *Chessboard { return NewChessboard().withKing(1).withRook(1).withKing(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withKing(2).withRook(1).Build() },
			expectedCount: 58344,
		},
		{
			name:          "R(1)+K(1)+R(1) equals R(2)+K(1)",
			description:   "Mixed Rook calls with King in between",
			builder1:      func() *Chessboard { return NewChessboard().withRook(1).withKing(1).withRook(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withRook(2).withKing(1).Build() },
			expectedCount: 49464,
		},
		{
			name:          "B(1)+N(1)+B(1) equals B(2)+N(1)",
			description:   "Mixed Bishop calls with Knight in between",
			builder1:      func() *Chessboard { return NewChessboard().withBishop(1).withKnight(1).withBishop(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withBishop(2).withKnight(1).Build() },
			expectedCount: 64640,
		},
		{
			name:          "Q(2)+R(1)+Q(2)+B(1) equals Q(4)+R(1)+B(1)",
			description:   "Complex mixed quantities",
			builder1:      func() *Chessboard { return NewChessboard().withQueen(2).withRook(1).withQueen(2).withBishop(1).Build() },
			builder2:      func() *Chessboard { return NewChessboard().withQueen(4).withRook(1).withBishop(1).Build() },
			expectedCount: 681240,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			boards1 := tc.builder1().calculateBoards()
			boards2 := tc.builder2().calculateBoards()

			// Check counts match
			if len(boards1) != tc.expectedCount {
				t.Errorf("Mixed quantity approach: got %d boards, want %d", len(boards1), tc.expectedCount)
			}
			if len(boards2) != tc.expectedCount {
				t.Errorf("Single call approach: got %d boards, want %d", len(boards2), tc.expectedCount)
			}

			// Check that both approaches produce identical results
			if len(boards1) != len(boards2) {
				t.Errorf("Approaches differ! Mixed=%d, Single=%d", len(boards1), len(boards2))
			}

			// Verify board configurations are identical
			for hash, board1 := range boards1 {
				board2, exists := boards2[hash]
				if !exists {
					t.Errorf("Board hash %v exists in mixed approach but not in single call", hash)
					break
				}
				if len(board1) != len(board2) {
					t.Errorf("Board %v has different contents: %d vs %d bytes", hash, len(board1), len(board2))
					break
				}
			}

			// Check reverse
			for hash := range boards2 {
				if _, exists := boards1[hash]; !exists {
					t.Errorf("Board hash %v exists in single call but not in mixed approach", hash)
					break
				}
			}
		})
	}
}

func Test_boardBuilder_withKing(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[byte]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  byte
		want2  int
	}{
		{"Test builder with king", fields{
			&Chessboard{},
			nil,
			make(map[byte]int),
		}, args{quantity: 1}, 'k', 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			build := (&boardBuilder{
				tt.fields.chessboard,
				tt.fields.currentFigureBehaviour,
				tt.fields.figureQuantityMap,
			}).withKing(tt.args.quantity).Build()

			if reflect.TypeOf(build.currentFigureBehaviour).String() != "*figures.King" {
				t.Error("figure type is not king")
			}

			actualNumberOfFigures := build.figureQuantityMap[tt.want1]

			if actualNumberOfFigures == 0 {
				t.Errorf("board does not contain king, actual number is 0, want %d", tt.want1)
			}
			if actualNumberOfFigures != tt.want2 {
				t.Errorf("actualNumberOfFigures = %v, want %v", actualNumberOfFigures, tt.want2)
			}
		})
	}
}

func Test_boardBuilder_withQueen(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[byte]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  byte
		want2  int
	}{
		{"Test builder with queen", fields{
			&Chessboard{},
			nil,
			make(map[byte]int),
		}, args{quantity: 1}, 'q', 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			build := (&boardBuilder{
				tt.fields.chessboard,
				tt.fields.currentFigureBehaviour,
				tt.fields.figureQuantityMap,
			}).withQueen(tt.args.quantity).Build()

			if reflect.TypeOf(build.currentFigureBehaviour).String() != "*figures.Queen" {
				t.Error("figure type is not queen")
			}

			actualNumberOfFigures := build.figureQuantityMap[tt.want1]

			if actualNumberOfFigures == 0 {
				t.Errorf("board does not contain queen, actual number is 0, want %d", tt.want1)
			}
			if actualNumberOfFigures != tt.want2 {
				t.Errorf("actualNumberOfFigures = %v, want %v", actualNumberOfFigures, tt.want2)
			}
		})
	}
}

func Test_boardBuilder_withBishop(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[byte]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  byte
		want2  int
	}{
		{"Test builder with bishop", fields{
			&Chessboard{},
			nil,
			make(map[byte]int),
		}, args{quantity: 1}, 'b', 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			build := (&boardBuilder{
				tt.fields.chessboard,
				tt.fields.currentFigureBehaviour,
				tt.fields.figureQuantityMap,
			}).withBishop(tt.args.quantity).Build()

			if reflect.TypeOf(build.currentFigureBehaviour).String() != "*figures.Bishop" {
				t.Error("figure type is not bishop")
			}

			actualNumberOfFigures := build.figureQuantityMap[tt.want1]

			if actualNumberOfFigures == 0 {
				t.Errorf("board does not contain bishop, actual number is 0, want %d", tt.want1)
			}
			if actualNumberOfFigures != tt.want2 {
				t.Errorf("actualNumberOfFigures = %v, want %v", actualNumberOfFigures, tt.want2)
			}
		})
	}
}

func Test_boardBuilder_withKnight(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[byte]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  byte
		want2  int
	}{
		{"Test builder with knight", fields{
			&Chessboard{},
			nil,
			make(map[byte]int),
		}, args{quantity: 1}, 'n', 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			build := (&boardBuilder{
				tt.fields.chessboard,
				tt.fields.currentFigureBehaviour,
				tt.fields.figureQuantityMap,
			}).withKnight(tt.args.quantity).Build()

			if reflect.TypeOf(build.currentFigureBehaviour).String() != "*figures.Knight" {
				t.Error("figure type is not knight")
			}

			actualNumberOfFigures := build.figureQuantityMap[tt.want1]

			if actualNumberOfFigures == 0 {
				t.Errorf("board does not contain knight, actual number is 0, want %d", tt.want1)
			}
			if actualNumberOfFigures != tt.want2 {
				t.Errorf("actualNumberOfFigures = %v, want %v", actualNumberOfFigures, tt.want2)
			}
		})
	}
}

func Test_boardBuilder_withRook(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[byte]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  byte
		want2  int
	}{
		{"Test builder with rook", fields{
			&Chessboard{},
			nil,
			make(map[byte]int),
		}, args{quantity: 1}, 'r', 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			build := (&boardBuilder{
				tt.fields.chessboard,
				tt.fields.currentFigureBehaviour,
				tt.fields.figureQuantityMap,
			}).withRook(tt.args.quantity).Build()

			if reflect.TypeOf(build.currentFigureBehaviour).String() != "*figures.Rook" {
				t.Error("figure type is not rook")
			}

			actualNumberOfFigures := build.figureQuantityMap[tt.want1]

			if actualNumberOfFigures == 0 {
				t.Errorf("board does not contain rook, actual number is 0, want %d", tt.want1)
			}
			if actualNumberOfFigures != tt.want2 {
				t.Errorf("actualNumberOfFigures = %v, want %v", actualNumberOfFigures, tt.want2)
			}
		})
	}
}
