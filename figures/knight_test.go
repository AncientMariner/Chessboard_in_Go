package figures

import (
	"reflect"
	"testing"
)

func TestKnight_GetName(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{"Test get name", fields{Figure{}}, 'n'},
		{"Test get name with empty figure", fields{}, 'n'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kn := &Knight{
				Figure: tt.fields.Figure,
			}
			if got := kn.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKnight_Handle(t *testing.T) {
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
		{"Test handle board size 8 is not possible", fields{Figure{next: nil}}, args{board: "________\n"}, 0},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: "________\n________\n________\n________\n________\n________\n________\n________\n"}, 64},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxx_____\n________\n________\n________\n________\n________\n"}, 54},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxx_____\n___xxx__\n___xkx__\n___xxx__\n________\n________\n"}, 41},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxx_____\n___xxx__\n___xkx__\n___xxx__\nxxx__xxx\nxkx__xkx\n"}, 24},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: "xxx__xxx\nxkx__xkx\nxxx__xxx\nxxxxxx__\nxkxxkx__\nxxxxxx__\nxxx__xxx\nxkx__xkx\n"}, 9},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: "xxx__xxx\nxkx__xkx\nxxx__xxx\nxxxxxx__\nxkxxkxxx\nxxxxxxxk\nxxxxxxxx\nxkxkxxkx\n"}, 4},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: "xxxxxxxx\nxkxkxxkx\nxxxxxxxx\nxxxxxx__\nxkxxkxxx\nxxxxxxxk\nxxxxxxxx\nxkxkxxkx\n"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			knight := &Knight{
				Figure: tt.fields.Figure,
			}
			if got := knight.Handle(tt.args.board); len(got) != tt.want {
				t.Errorf("Handle() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_isAnotherFigurePresentAbove(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test vertically 1 above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 5,
		}, false},
		{"Test vertically 1 above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 10,
		}, false},
		{"Test vertically 1 above", args{
			out:      []rune{'b', '_', '_', 'k', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 10,
		}, true},
		{"Test vertically 1 above", args{
			out:      []rune{'b', '_', '_', '_', 'k', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 11,
		}, true},
		{"Test vertically 1 above", args{
			out:      []rune{'b', '_', '_', '_', 'k', 'q', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 16,
		}, true},
		{"Test vertically 1 above", args{
			out:      []rune{'_', 'b', '_', 'b', '_', 'q', '_', '_', '\n', 'k', '_', '_', '_', 'k', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 20,
		}, true},
		{"Test vertically 1 above", args{
			out:      []rune{'_', 'b', 'k', 'b', '_', 'q', '_', '_', '\n', 'k', '_', 'r', '_', 'k', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 18,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnotherFigurePresentAbove(tt.args.out, tt.args.position); got != tt.want {
				t.Errorf("isAnotherFigurePresentAbove() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

func Test_isAnotherFigurePresentBelow(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test vertically 1 below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 5,
		}, false},
		{"Test vertically 1 below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 10,
		}, false},
		{"Test vertically 1 below", args{
			out:      []rune{'b', '_', '_', 'k', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 10,
		}, false},
		{"Test vertically 1 below", args{
			out:      []rune{'b', '_', '_', '_', 'k', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 11,
		}, false},
		{"Test vertically 1 below", args{
			out:      []rune{'b', '_', '_', '_', 'k', 'q', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 16,
		}, false},
		{"Test vertically 1 below", args{
			out:      []rune{'_', 'b', '_', 'b', '_', 'q', '_', '_', '\n', 'k', '_', '_', '_', 'k', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 20,
		}, false},
		{"Test vertically 1 below", args{
			out:      []rune{'_', 'b', 'k', 'b', '_', 'q', '_', '_', '\n', 'k', '_', 'r', '_', 'k', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 18,
		}, false},
		{"Test vertically 1 below", args{
			out:      []rune{'_', 'b', 'k', 'b', '_', '_', '_', '_', '\n', '_', '_', 'b', '_', '_', '_', 'b', '_', '\n', '_', '_', '_', 'b', '_', 'b', '_', '_', '\n'},
			position: 4,
		}, true},
		{"Test vertically 1 below", args{
			out:      []rune{'_', 'b', 'k', 'b', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', 'b', '_', '\n', '_', '_', '_', 'b', '_', 'b', '_', '_', '\n'},
			position: 4,
		}, true},
		{"Test vertically 1 below", args{
			out:      []rune{'_', 'b', 'k', 'b', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', 'b', '_', 'b', '_', '_', '\n'},
			position: 4,
		}, true},
		{"Test vertically 1 below", args{
			out:      []rune{'_', 'b', 'k', 'b', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', 'b', '_', '_', '\n'},
			position: 4,
		}, true},
		{"Test vertically 1 below", args{
			out:      []rune{'_', 'b', 'k', 'b', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 4,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnotherFigurePresentBelow(tt.args.out, tt.args.position); got != tt.want {
				t.Errorf("isAnotherFigurePresentBelow() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

func Test_placeAttackPlacesBelow(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Test vertically 1 below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 2,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', 'x', '_', '_', '_', '\n', '_', 'x', '_', 'x', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesBelow(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesBelow() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

func Test_placeAttackPlacesAbove(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Test above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 20,
		}, []rune{'_', 'x', '_', 'x', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', 'x', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},

		{"Test below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 18,
		}, []rune{'_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', 'x', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesAbove(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesAbove() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}
