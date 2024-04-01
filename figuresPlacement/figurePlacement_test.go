package figuresPlacement

import (
	"Chessboard_in_Go/figures"
	"github.com/hashicorp/go-set/v2"
	"reflect"
	"testing"
)

func Test_drawEmptyBoard(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Initial board test", "________\n________\n________\n________\n________\n________\n________\n________\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := drawEmptyBoard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("drawEmptyBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlacement_PlaceFigure(t *testing.T) {
	type fields struct {
		currentPlacement FigurePlacement
	}
	type args struct {
		numberOfFigures int
		behaviour       figures.FigureBehaviour
		boards          *set.HashSet[*FigurePosition, string]
	}

	setOfBoards := set.NewHashSet[*FigurePosition, string](1)
	setOfBoards.Insert(&FigurePosition{
		Board:  "________\n________\n________\n________\n________\n________\n________\n________\n",
		number: 1,
	})

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test place figures without figures", fields{nil}, args{
			0,
			&figures.King{},
			set.NewHashSet[*FigurePosition, string](0),
		}, 0},
		{"Test place figures on empty board", fields{nil}, args{
			1,
			&figures.King{},
			set.NewHashSet[*FigurePosition, string](0),
		}, 64},
		{"Test place figures on board", fields{nil}, args{
			1,
			&figures.King{},
			setOfBoards,
		}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{
				currentPlacement: tt.fields.currentPlacement,
			}
			if got := p.PlaceFigure(tt.args.numberOfFigures, tt.args.behaviour, tt.args.boards); got.Size() != tt.want {
				t.Errorf("PlaceFigure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlacement_PlaceFiguresOnEmptyBoard(t *testing.T) {
	type fields struct {
		currentPlacement FigurePlacement
	}
	type args struct {
		board     string
		behaviour figures.FigureBehaviour
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test placement on empty board", fields{nil}, args{"________\n________\n________\n________\n________\n________\n________\n________\n", &figures.King{}}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{
				currentPlacement: tt.fields.currentPlacement,
			}
			if got := p.PlaceFiguresOnEmptyBoard(tt.args.board, tt.args.behaviour); got.Size() != tt.want {
				t.Errorf("PlaceFiguresOnEmptyBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlacement_placeFiguresOnBoard(t *testing.T) {
	type fields struct {
		currentPlacement FigurePlacement
	}
	type args struct {
		boards    *set.HashSet[*FigurePosition, string]
		behaviour figures.FigureBehaviour
	}

	setOfBoards := set.NewHashSet[*FigurePosition, string](1)
	setOfBoards.Insert(&FigurePosition{
		Board:  "________\n________\n________\n________\n________\n________\n________\n________\n",
		number: 1,
	})

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test placement on board", fields{nil}, args{setOfBoards, &figures.King{}}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{
				currentPlacement: tt.fields.currentPlacement,
			}
			if got := p.placeFiguresOnBoard(tt.args.boards, tt.args.behaviour); got.Size() != tt.want {
				t.Errorf("placeFiguresOnBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
