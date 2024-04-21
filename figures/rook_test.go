package figures

import (
	"reflect"
	"testing"
)

func TestRook_GetName(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{"Test get name", fields{Figure{}}, 'r'},
		{"Test get name with empty figure", fields{}, 'r'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ro := &Rook{
				Figure: tt.fields.Figure,
			}
			if got := ro.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRook_Handle(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	type args struct {
		board string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// {"Test handle board size 8", fields{Figure{next: nil}}, args{board: "____\n"}, 4},
		{"Test handle board size 8", fields{Figure{next: nil}}, args{board: "________\n"}, 8},
		{"Test handle board size 64", fields{Figure{next: nil}}, args{board: "________\n________\n________\n________\n________\n________\n________\n________\n"}, 64},
		{"Test handle board size 50", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxx_____\n________\n________\n________\n________\n________\n"}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rook := &Rook{
				Figure: tt.fields.Figure,
			}
			if got := rook.Handle(tt.args.board); got.Size() != tt.want {
				t.Errorf("Handle() size = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRook_placeAttackPlacesHorizontally(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name                      string
		args                      args
		wantBoard                 []rune
		wantPossibleToPlaceFigure bool
	}{
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, []rune{'_', 'x', 'x', 'x', 'x', 'x', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			1}, []rune{'x', '_', 'x', 'x', 'x', 'x', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			2}, []rune{'x', 'x', '_', 'x', 'x', 'x', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			3}, []rune{'x', 'x', 'x', '_', 'x', 'x', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			4}, []rune{'x', 'x', 'x', 'x', '_', 'x', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			5}, []rune{'x', 'x', 'x', 'x', 'x', '_', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			6}, []rune{'x', 'x', 'x', 'x', 'x', 'x', '_', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			7}, []rune{'x', 'x', 'x', 'x', 'x', 'x', 'x', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			8}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			9}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', 'x', 'x', 'x', 'x', 'x', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', 'x', 'x', 'x', 'x', 'x', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			11}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', 'x', '_', 'x', 'x', 'x', 'x', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			12}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', 'x', 'x', '_', 'x', 'x', 'x', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			13}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', 'x', 'x', 'x', '_', 'x', 'x', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			14}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', 'x', 'x', 'x', 'x', '_', 'x', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			15}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', 'x', 'x', 'x', 'x', 'x', '_', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			16}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', 'x', 'x', 'x', 'x', 'x', 'x', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			17}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			18}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			18}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', 'x', 'x', 'x', 'x', 'x', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			14}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', 'x', 'x', 'x', 'x', '_', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			19}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', 'x', 'x', 'x', 'x', 'x', 'x', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			25}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', 'x', 'x', 'x', 'x', 'x', 'x', '_', '\n'}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			26}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			27}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isHorizontalPlacementPossiblely := (&Rook{}).placeAttackPlacesHorizontally(tt.args.out, tt.args.position)

			if isHorizontalPlacementPossiblely != tt.wantPossibleToPlaceFigure {
				t.Errorf("isHorizontalPlacementPossiblely = %v, wantPossibleToPlaceFigure %v", isHorizontalPlacementPossiblely, tt.wantPossibleToPlaceFigure)
			}
			if !reflect.DeepEqual(tt.args.out, tt.wantBoard) {
				t.Errorf("placeAttackPlacesDiagonallyBelow() = %v, wantBoard %v", tt.args.out, tt.wantBoard)
			}
		})
	}
}

func TestRook_placeAttackPlacesVertically(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name      string
		args      args
		wantBoard []rune
	}{
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			7}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			8}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			9}, []rune{'x', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, []rune{'_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			16}, []rune{'_', '_', '_', '_', '_', '_', '_', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			17}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			18}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			18}, []rune{'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			19}, []rune{'_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			21}, []rune{'_', '_', '_', 'x', '_', '_', '_', '_', '\n', '_', '_', '_', 'x', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', 'x', '_', '_', '_', '_', '\n', '_', '_', '_', 'x', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			24}, []rune{'_', '_', '_', '_', '_', '_', 'x', '_', '\n', '_', '_', '_', '_', '_', '_', 'x', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			25}, []rune{'_', '_', '_', '_', '_', '_', '_', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n'}},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			26}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#",
			args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
				27},
			[]rune{'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			(&Rook{}).placeAttackPlacesVertically(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.wantBoard) {
				t.Errorf("placeAttackPlacesVertically() = %v, wantBoard %v", tt.args.out, tt.wantBoard)
			}
		})
	}
}

func Test_isAnotherFigurePresentOnTheLine(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, false},
		{"Test position#", args{[]rune{'_', 'k', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'k', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'k', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			14}, true},
		{"Test position#", args{[]rune{'_', '_', 'q', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, true},
		{"Test position#", args{[]rune{'_', '_', '_', 'b', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', 'n', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', 'r', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, true},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', 'x', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, false},
		{"Test position#", args{[]rune{'_', 'x', 'x', 'x', 'x', 'x', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, false},
		{"Test position#", args{[]rune{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnotherFigurePresentOnTheLineHorizontally(tt.args.out, tt.args.position); got != tt.want {
				t.Errorf("isAnotherFigurePresentOnTheLineHorizontally() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnotherFigurePresentOnTheLineVertically(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			1}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			2}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			3}, false},

		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			4}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			5}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			6}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			7}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'k', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			8}, true},

		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			14}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			19}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			23}, false},

		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			28}, false},
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'k', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnotherFigurePresentOnTheLineVertically(tt.args.out, tt.args.position); got != tt.want {
				t.Errorf("isAnotherFigurePresentOnTheLineVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}
