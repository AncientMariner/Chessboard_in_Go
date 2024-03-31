package main

import (
	"Chessboard_in_Go/figures"
	"Chessboard_in_Go/figuresPlacement"
	"reflect"
	"testing"

	"github.com/hashicorp/go-set/v2"
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
				figuresPlacement.Placement{},
			},
			nil,
			nil,
		}, args{figure: &figures.Queen{}},
			&boardBuilder{
				&Chessboard{
					nil,
					&figures.Queen{},
					figuresPlacement.Placement{},
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
				figuresPlacement.Placement{},
			},
			&figures.Queen{},
			map[rune]int{'q': 1},
		}, args{figure: &figures.Queen{}},
			&boardBuilder{
				&Chessboard{
					map[rune]int{'k': 1},
					&figures.King{},
					figuresPlacement.Placement{},
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
		{"Test empty board with 1 queen", args{board: NewChessboard().withQueen(1).Build()}, 64},
		{"Test empty board with 1 rook", args{board: NewChessboard().withRook(1).Build()}, 64},
		{"Test empty board with 1 knight", args{board: NewChessboard().withKnight(1).Build()}, 64},
		{"Test empty board with 1 bishop", args{board: NewChessboard().withBishop(1).Build()}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.board.placeFigures(); got.Size() != tt.want {
				t.Errorf("placeFigures() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_board_with_1_figure(t *testing.T) {
	type args struct {
		board *Chessboard
		figuresBehaviour figures.FigureBehaviour
	}
	king := &figures.King{}
	queen := &figures.Queen{}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{"Test empty board", args{&Chessboard{}}, 0},
		{"Test empty board with 1 king", args{&Chessboard{
			map[rune]int{king.GetName(): 1},
			king,
			figuresPlacement.Placement{},
		}, king}, 64},
		{"Test empty board with 1 king", args{&Chessboard{
			map[rune]int{queen.GetName(): 1},
			queen,
			figuresPlacement.Placement{},
		}, queen}, 64},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.board.placeFigure(tt.args.figuresBehaviour,  set.NewHashSet[*figuresPlacement.FigurePosition, string](0) ); got.Size() != tt.want {
				t.Errorf("placeFigures() = %v, want %v", got, tt.want)
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
