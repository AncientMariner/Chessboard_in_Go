package figures

import (
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
		behaviour       FigureBehaviour
		boards          map[string]string
	}

	item := &BoardWithFigurePosition{}
	item.Board = "________\n________\n________\n________\n________\n________\n________\n________\n"
	// item.number = 1
	item.Hash()

	newMap := make(map[string]string)
	newMap[item.hash] = item.Board

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test place figures without figures", fields{nil}, args{
			0,
			&King{},
			make(map[string]string),
		}, 0},
		{"Test place figures on empty board", fields{nil}, args{
			1,
			&King{},
			make(map[string]string),
		}, 64},
		{"Test place figures on board", fields{nil}, args{
			1,
			&King{},
			newMap,
		}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{
				currentPlacement: tt.fields.currentPlacement,
			}
			if got := p.PlaceFigure(tt.args.numberOfFigures, tt.args.behaviour, tt.args.boards); len(got) != tt.want {
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
		behaviour FigureBehaviour
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test placement on empty board", fields{nil}, args{"________\n________\n________\n________\n________\n________\n________\n________\n", &King{}}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{
				currentPlacement: tt.fields.currentPlacement,
			}
			if got := p.PlaceFiguresOnEmptyBoard(tt.args.board, tt.args.behaviour); len(got) != tt.want {
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
		boards    map[string]string
		behaviour FigureBehaviour
	}

	item := &BoardWithFigurePosition{}
	item.Board = string("________\n________\n________\n________\n________\n________\n________\n________\n")
	// item.number = 1
	item.Hash()

	newMap := make(map[string]string)
	newMap[item.hash] = item.Board

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test placement on board", fields{nil}, args{newMap, &King{}}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{
				currentPlacement: tt.fields.currentPlacement,
			}
			if got := p.placeFiguresOnBoard(tt.args.boards, tt.args.behaviour); len(got) != tt.want {
				t.Errorf("placeFiguresOnBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
