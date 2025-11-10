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

func TestPlacement_PlaceFigures(t *testing.T) {
	type args struct {
		numberOfFigures int
		behaviour       FigureBehaviour
		boards          map[string]string
	}

	board := "________\n________\n________\n________\n________\n________\n________\n________\n"
	hash := GenerateHash(board)

	newMap := make(map[string]string)
	newMap[hash] = board

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test place figures without figures", args{
			0,
			&King{},
			make(map[string]string),
		}, 0},
		{"Test place figures on empty board", args{
			1,
			&King{},
			make(map[string]string),
		}, 64},
		{"Test place figures on board", args{
			1,
			&King{},
			newMap,
		}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{}
			if got := p.PlaceFigures(tt.args.numberOfFigures, tt.args.behaviour, tt.args.boards); len(got) != tt.want {
				t.Errorf("PlaceFigures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlacement_placeFiguresOnBoard(t *testing.T) {
	type fields struct {
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
		{"Test placement on empty board", fields{}, args{"________\n________\n________\n________\n________\n________\n________\n________\n", &King{}}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{}
			if got := p.placeFigureOnBoard(tt.args.board, tt.args.behaviour); len(got) != tt.want {
				t.Errorf("PlaceFiguresOnEmptyBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlacement_placeFigure(t *testing.T) {
	type fields struct {
	}
	type args struct {
		boards    map[string]string
		behaviour FigureBehaviour
	}

	board := "________\n________\n________\n________\n________\n________\n________\n________\n"
	hash := GenerateHash(board)

	newMap := make(map[string]string)
	newMap[hash] = board

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test placement on board", fields{}, args{newMap, &King{}}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Placement{}
			if got := p.placeFigure(tt.args.boards, tt.args.behaviour); len(got) != tt.want {
				t.Errorf("placeFigure() = %v, want %v", got, tt.want)
			}
		})
	}
}
