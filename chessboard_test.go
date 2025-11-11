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
		{"Test empty chessboard", &boardBuilder{chessboard: &Chessboard{}, figureQuantityMap: make(map[rune]int)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChessboard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChessboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardBuilder_Build(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[rune]int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Chessboard
	}{
		{"Test build", fields{
			chessboard:             &Chessboard{},
			currentFigureBehaviour: nil,
			figureQuantityMap:      make(map[rune]int),
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
		figureQuantityMap      map[rune]int
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
		figureQuantityMap      map[rune]int
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
				map[rune]int{'k': 1},
				&figures.King{},
				figures.Placement{},
			},
			&figures.Queen{},
			map[rune]int{'q': 1},
		}, args{figure: &figures.Queen{}},
			&boardBuilder{
				&Chessboard{
					map[rune]int{'k': 1},
					&figures.King{},
					figures.Placement{},
				},
				&figures.Queen{},
				map[rune]int{'q': 1},
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
	t.Skip("skipping test")

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
		{"Test empty board with 2 king", args{board: NewChessboard().withKing(2).Build()}, 1806}, // 3612 if black and white king
		{"Test empty board with 1 king 1 king", args{board: NewChessboard().withKing(1).withKing(1).Build()}, 1806},
		{"Test empty board with 2 rook", args{board: NewChessboard().withRook(2).Build()}, 1568},
		{"Test empty board with 1 rook 1 rook", args{board: NewChessboard().withRook(1).withRook(1).Build()}, 1568},
		{"Test empty board with 1 king 1 rook", args{board: NewChessboard().withKing(1).withRook(1).Build()}, 2952},
		{"Test empty board with 1 rook 1 king", args{board: NewChessboard().withRook(1).withKing(1).Build()}, 2952},
		{"Test empty board with 1 rook 1 king", args{board: NewChessboard().withRook(2).withKing(2).Build()}, 669496},
		// {"Test empty board with 8 rook", args{board: NewChessboard().withRook(8).Build()}, 40320},
		{"Test empty board with 1 king 1 bishop", args{board: NewChessboard().withKing(1).withBishop(1).Build()}, 3248},
		{"Test empty board with 2 king 2 bishop", args{board: NewChessboard().withKing(2).withBishop(2).Build()}, 1177824},
		{"Test empty board with 1 bishop 1 king", args{board: NewChessboard().withBishop(1).withKing(1).Build()}, 3178},
		// {"Test empty board with 14 bishop", args{board: NewChessboard().withBishop(14).Build()}, 1736},
		{"Test empty board with 2 bishop ", args{board: NewChessboard().withBishop(2).Build()}, 1736},
		{"Test empty board with 1 bishop 1 bishop", args{board: NewChessboard().withBishop(1).withBishop(1).Build()}, 1736},
		{"Test empty board with 1 bishop 1 rook", args{board: NewChessboard().withBishop(1).withRook(1).Build()}, 2506},
		{"Test empty board with 1 rook 1 bishop", args{board: NewChessboard().withRook(1).withBishop(1).Build()}, 2576},
		// {"Test empty board with 32 knight", args{board: NewChessboard().withKnight(32).Build()}, 3063828},
		{"Test empty board with 1 knight 1 knight", args{board: NewChessboard().withKnight(1).withKnight(1).Build()}, 1848},
		{"Test empty board with 1 king 1 knight", args{board: NewChessboard().withKing(1).withKnight(1).Build()}, 3456},
		{"Test empty board with 1 knight 1 king", args{board: NewChessboard().withKnight(1).withKing(1).Build()}, 3288},
		{"Test empty board with 1 queen 1 knight", args{board: NewChessboard().withQueen(1).withKnight(1).Build()}, 2338},
		{"Test empty board with 1 queen 1 bishop", args{board: NewChessboard().withQueen(1).withBishop(1).Build()}, 2506},
		{"Test empty board with 1 bishop 1 knight", args{board: NewChessboard().withRook(1).withKnight(1).Build()}, 2968},
		{"Test empty board with 1 queen 1 rook", args{board: NewChessboard().withQueen(1).withRook(1).Build()}, 2506},
		{"Test empty board with 1 queen 1 rook", args{board: NewChessboard().withRook(1).withKnight(1).Build()}, 2968},
		{"Test empty board with 1 queen 1 rook", args{board: NewChessboard().withBishop(1).withKnight(1).Build()}, 3234},
		// {"Test empty board with 16 king 14 bishop", args{board: NewChessboard().withKing(16).withBishop(14).Build()}, 3063828},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(1).Build()}, 64},
		{"Test empty board with 2 queen", args{board: NewChessboard().withQueen(2).Build()}, 1288},
		{"Test empty board with 2 queen", args{board: NewChessboard().withQueen(1).withQueen(1).Build()}, 1288},
		{"Test empty board with 1 king 1 queen", args{board: NewChessboard().withKing(1).withQueen(1).Build()}, 2576},
		{"Test empty board with 1 queen 1 king", args{board: NewChessboard().withQueen(1).withKing(1).Build()}, 2506},
		{"Test empty board with 3 queen", args{board: NewChessboard().withQueen(3).Build()}, 10320},
		{"Test empty board with 4 queen", args{board: NewChessboard().withQueen(4).Build()}, 34568},
		{"Test empty board with 5 queen", args{board: NewChessboard().withQueen(5).Build()}, 46736},
		{"Test empty board with 6 queen", args{board: NewChessboard().withQueen(6).Build()}, 22708},
		{"Test empty board with 7 queen", args{board: NewChessboard().withQueen(7).Build()}, 3192},
		{"Test empty board with 8 queen", args{board: NewChessboard().withQueen(8).Build()}, 92},
		{"Test empty board with 8 queen", args{board: NewChessboard().withQueen(1).withQueen(1).withQueen(1).withQueen(1).withQueen(1).withQueen(1).withQueen(1).withQueen(1).Build()}, 92},
		{"Test empty board with 9 queen, impossible case", args{board: NewChessboard().withQueen(9).Build()}, 0},
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

func Test_number_of_boards_with_1_figure_7x7(t *testing.T) {
	type args struct {
		board *Chessboard
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test empty board 7x7 with 1 king", args{board: NewChessboardWithSize(7).withKing(1).Build()}, 49},
		{"Test empty 7x7 board with 2 king 2 queen 2 bishop 1 knight", args{board: NewChessboardWithSize(7).withKing(2).withQueen(2).withBishop(2).withKnight(1).Build()}, 3761852},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.board.calculateBoards(); len(got) != tt.want {
				t.Errorf("calculateBoards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_of_boards_with_different_figure_variations(t *testing.T) {
	calculateBoards_R_K_R := NewChessboard().withRook(1).withKing(1).withRook(1).Build().calculateBoards()
	calculateBoards_R_R_K := NewChessboard().withRook(1).withRook(1).withKing(1).Build().calculateBoards()
	calculateBoards_K_R_R := NewChessboard().withKing(1).withRook(1).withRook(1).Build().calculateBoards()

	var unitedSet = make(map[string]string, len(calculateBoards_R_K_R)+len(calculateBoards_R_R_K)+len(calculateBoards_K_R_R))

	for u, s := range calculateBoards_R_K_R {
		unitedSet[u] = s
	}

	for u, s := range calculateBoards_R_R_K {
		unitedSet[u] = s
	}

	for u, s := range calculateBoards_K_R_R {
		unitedSet[u] = s
	}

	if len(calculateBoards_R_K_R) != 49887 {
		t.Errorf("calculateBoards() all possible variations = %v, want %v", len(calculateBoards_R_K_R), 49887)
	}
	if len(calculateBoards_R_K_R) != 49887 {
		t.Errorf("calculateBoards() all possible variations = %v, want %v", len(calculateBoards_R_K_R), 49887)
	}
	if len(calculateBoards_K_R_R) != 49887 {
		t.Errorf("calculateBoards() all possible variations = %v, want %v", len(calculateBoards_K_R_R), 49887)
	}
	if len(unitedSet) != 49887 {
		t.Errorf("calculateBoards() all possible variations = %v, want %v", len(unitedSet), 49887)
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
		figureQuantityMap      map[rune]int
		currentFigureBehaviour figures.FigureBehaviour
		figurePlacement        figures.Placement
	}
	type args struct {
		behaviour            figures.FigureBehaviour
		previousFigureBoards map[string]string
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
		{"Test empty board with 1 king", fields{map[rune]int{(&figures.King{}).GetName(): 1}, &figures.King{}, figures.Placement{}}, args{&figures.King{}, make(map[string]string)}, 64},
		{"Test empty board with 2 king", fields{map[rune]int{(&figures.King{}).GetName(): 2}, &figures.King{}, figures.Placement{}}, args{&figures.King{}, make(map[string]string)}, 1806},
		{"Test empty board with 1 king 1 rook", fields{map[rune]int{(&figures.King{}).GetName(): 1, (&figures.Rook{}).GetName(): 1}, figureBehaviour, figures.Placement{}}, args{figureBehaviour, make(map[string]string)}, 2952},
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

func Test_boardBuilder_withKing(t *testing.T) {
	type fields struct {
		chessboard             *Chessboard
		currentFigureBehaviour figures.FigureBehaviour
		figureQuantityMap      map[rune]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  rune
		want2  int
	}{
		{"Test builder with king", fields{
			&Chessboard{},
			nil,
			make(map[rune]int),
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
		figureQuantityMap      map[rune]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  rune
		want2  int
	}{
		{"Test builder with queen", fields{
			&Chessboard{},
			nil,
			make(map[rune]int),
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
		figureQuantityMap      map[rune]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  rune
		want2  int
	}{
		{"Test builder with bishop", fields{
			&Chessboard{},
			nil,
			make(map[rune]int),
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
		figureQuantityMap      map[rune]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  rune
		want2  int
	}{
		{"Test builder with knight", fields{
			&Chessboard{},
			nil,
			make(map[rune]int),
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
		figureQuantityMap      map[rune]int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  rune
		want2  int
	}{
		{"Test builder with rook", fields{
			&Chessboard{},
			nil,
			make(map[rune]int),
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
