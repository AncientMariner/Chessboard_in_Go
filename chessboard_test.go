package main

import (
	"Chessboard_in_Go/figures"
	"reflect"
	"testing"
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

// todo add more variations
func Test_number_of_boards_with_1_figure(t *testing.T) {
	type args struct {
		board *Chessboard
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test empty board with 1 king", args{board: NewChessboard().withKing(1).Build()}, 64},
		{"Test empty board with 2 king", args{board: NewChessboard().withKing(2).Build()}, 1806}, // 3612 if black and white king
		{"Test empty board with 1 king 1 king", args{board: NewChessboard().withKing(1).withKing(1).Build()}, 1806},
		{"Test empty board with 2 rook", args{board: NewChessboard().withRook(2).Build()}, 3080},
		{"Test empty board with 1 rook 1 rook", args{board: NewChessboard().withRook(1).withRook(1).Build()}, 3080},
		{"Test empty board with 2 bishop ", args{board: NewChessboard().withBishop(2).Build()}, 2436},
		{"Test empty board with 1 bishop 1 bishop", args{board: NewChessboard().withBishop(1).withBishop(1).Build()}, 2436},
		{"Test empty board with 1 king 1 rook", args{board: NewChessboard().withKing(1).withBishop(1).Build()}, 3251},
		{"Test empty board with 1 bishop 1 king", args{board: NewChessboard().withBishop(1).withKing(1).Build()}, 3248},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(1).Build()}, 64},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(2).Build()}, 1226},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(1).withQueen(1).Build()}, 1226},
		{"Test empty board with 1 queen", args{board: NewChessboard().withKing(1).withQueen(1).Build()}, 2471},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(1).withKing(1).Build()}, 2535},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(7).Build()}, 727},
		// todo check hashcode
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(4).Build()}, 25813},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(5).Build()}, 26116},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(6).Build()}, 4},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(7).Build()}, 727},
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(8).Build()}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.board.placeFigures(); len(got) != tt.want {
				t.Errorf("placeFigures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_of_boards_with_different_figure_variations(t *testing.T) {
	placeFigures_R_K_R := NewChessboard().withRook(1).withKing(1).withRook(1).Build().placeFigures()
	placeFigures_R_R_K := NewChessboard().withRook(1).withRook(1).withKing(1).Build().placeFigures()
	placeFigures_K_R_R := NewChessboard().withKing(1).withRook(1).withRook(1).Build().placeFigures()

	var unitedSet = make(map[uint32]string, len(placeFigures_R_K_R)+len(placeFigures_R_R_K)+len(placeFigures_K_R_R))

	for u, s := range placeFigures_R_K_R {
		unitedSet[u] = s
	}

	for u, s := range placeFigures_R_R_K {
		unitedSet[u] = s
	}

	for u, s := range placeFigures_K_R_R {
		unitedSet[u] = s
	}

	// counterOfNotUniqueItemsInSet := 0
	// unitedSet.ForEach(func(position *figures.BoardWithFigurePosition) bool {
	// 	if contains(unitedBoards, position.Board) {
	// 		counterOfNotUniqueItemsInSet++
	// 	}
	// 	unitedBoards = append(unitedBoards, position.Board)
	// 	return true
	// })
	if len(unitedSet) != 113022 {
		t.Errorf("placeFigures() all possible variations = %v, want %v", len(unitedSet), 113022)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func Test_board_with_1_figure(t *testing.T) {
	type fields struct {
		figureQuantityMap      map[rune]int
		currentFigureBehaviour figures.FigureBehaviour
		figurePlacement        figures.Placement
	}
	type args struct {
		behaviour            figures.FigureBehaviour
		previousFigureBoards map[uint32]string
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
		{"Test empty board with 1 king", fields{map[rune]int{(&figures.King{}).GetName(): 1}, &figures.King{}, figures.Placement{}}, args{&figures.King{}, make(map[uint32]string)}, 64},
		{"Test empty board with 2 king", fields{map[rune]int{(&figures.King{}).GetName(): 2}, &figures.King{}, figures.Placement{}}, args{&figures.King{}, make(map[uint32]string)}, 3612},
		{"Test empty board with 1 king 1 rook", fields{map[rune]int{(&figures.King{}).GetName(): 1, (&figures.Rook{}).GetName(): 1}, figureBehaviour, figures.Placement{}}, args{figureBehaviour, make(map[uint32]string)}, 2952},
		// {"Test empty board with 1 king 1 queen", fields{map[rune]int{(&figures.King{}).GetName(): 1, (&figures.Queen{}).GetName(): 1}, behaviour, figures.Placement{}}, args{behaviour, set.NewHashSet[*figures.BoardWithFigurePosition, string](0)}, 4032},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &Chessboard{
				figureQuantityMap:      tt.fields.figureQuantityMap,
				currentFigureBehaviour: tt.fields.currentFigureBehaviour,
				figurePlacement:        tt.fields.figurePlacement,
			}
			if got := board.placeFigure(tt.args.behaviour, tt.args.previousFigureBoards); len(got) != tt.want {
				t.Errorf("placeFigure() = %v, want %v", got, tt.want)
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
