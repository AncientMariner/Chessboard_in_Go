package figures

import (
	"reflect"
	"testing"
)

func TestBishop_GetName(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{"Test get name", fields{Figure{}}, 'b'},
		{"Test get name with empty figure", fields{}, 'b'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bi := &Bishop{
				Figure: tt.fields.Figure,
			}
			if got := bi.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBishop_Handle(t *testing.T) {
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
		{"Test handle board with 1 king size 50", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxx_____\n________\n________\n________\n________\n________\n"}, 50},
		{"Test handle board with kings size 44", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxx_____\n________\n________\n_____xxx\n_____xkx\n_____xxx\n"}, 44},
		{"Test handle board with kings size 32", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxxxxx__\n___xkx__\n___xxx__\n_____xxx\n_____xkx\n_____xxx\n"}, 32},
		{"Test handle board with kings size 20", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxxxxx__\n___xkx__\n___xxx__\n__xxxxxx\n__xkxxkx\n__xxxxxx\n"}, 20},
		{"Test handle board with kings size 11", fields{Figure{next: nil}}, args{board: "xxxxxx__\nxkxxkx__\nxxxxxx__\n___xkx__\n___xxx__\n__xxxxxx\n__xkxxkx\n__xxxxxx\n"}, 11},
		{"Test handle board with kings size 7", fields{Figure{next: nil}}, args{board: "xxxxxx__\nxkxxkx__\nxxxxxx__\nxkxxkx__\nxxxxxx__\nxxxxxxxx\n__xkxxkx\n__xxxxxx\n"}, 7},
		{"Test handle board with kings size 5", fields{Figure{next: nil}}, args{board: "xxxxxx__\nxkxxkx__\nxxxxxx__\nxkxxkxxx\nxxxxxxxk\nxxxxxxxx\n__xkxxkx\n__xxxxxx\n"}, 5},
		{"Test handle board with kings size 1", fields{Figure{next: nil}}, args{board: "xxxxxxxx\nxkxxkxxk\nxxxxxxxx\nxkxxkxxx\nxxxxxxxk\nxxxxxxxx\n__xkxxkx\n__xxxxxx\n"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bishop := &Bishop{
				Figure: tt.fields.Figure,
			}
			if got := bishop.Handle(tt.args.board); len(got) != tt.want {
				t.Errorf("Handle() size= %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBishop_placeAttackPlacesDiagonallyBelow(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', 'x', '_', '_', '_', '_', '_', '\n', '_', '_', '_', 'x', '_', '_', '_', '_', '\n', '_', '_', '_', '_', 'x', '_', '_', '_', '\n', '_', '_', '_', '_', '_', 'x', '_', '_', '\n', '_', '_', '_', '_', '_', '_', 'x', '_', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n'}},
		{"Test position#", args{
			[]rune{
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			4}, []rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', 'x', '_', 'x', '_', '_', '\n',
			'_', '_', 'x', '_', '_', '_', 'x', '_', '\n',
			'_', 'x', '_', '_', '_', '_', '_', 'x', '\n',
			'x', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n',
				'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			3}, []rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', 'x', '_', 'x', '_', '_', '_', '\n',
			'_', 'x', '_', '_', '_', 'x', '_', '_', '\n',
			'x', '_', '_', '_', '_', '_', 'x', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'x', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			1}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', 'x', '_', '_', '_', '_', '_', '\n', '_', '_', '_', 'x', '_', '_', '_', '_', '\n', '_', '_', '_', '_', 'x', '_', '_', '_', '\n', '_', '_', '_', '_', '_', 'x', '_', '_', '\n', '_', '_', '_', '_', '_', '_', 'x', '_', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			8}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			17}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			100}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			5}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', 'x', '_', 'x', '_', '\n', '_', '_', '_', 'x', '_', '_', '_', 'x', '\n', '_', '_', 'x', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			9}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', 'x', '_', '_', '_', '_', '_', '\n', '_', '_', '_', 'x', '_', '_', '_', '_', '\n', '_', '_', '_', '_', 'x', '_', '_', '_', '\n', '_', '_', '_', '_', '_', 'x', '_', '_', '\n', '_', '_', '_', '_', '_', '_', 'x', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', 'x', '_', '_', '_', '_', '_', '\n', '_', '_', '_', 'x', '_', '_', '_', '_', '\n', '_', '_', '_', '_', 'x', '_', '_', '_', '\n', '_', '_', '_', '_', '_', 'x', '_', '_', '\n', '_', '_', '_', '_', '_', '_', 'x', '_', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			30}, []rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', 'x', '_', 'x', '_', '_', '_', '\n',
			'_', 'x', '_', '_', '_', 'x', '_', '_', '\n',
			'x', '_', '_', '_', '_', '_', 'x', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'x', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesDiagonallyBelow(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesDiagonallyBelow() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

func TestBishop_placeAttackPlacesDiagonallyAbove(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			1}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			5}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			9}, []rune{'_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, []rune{'x', '_', 'x', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			17}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			30}, []rune{
			'x', '_', '_', '_', '_', '_', 'x', '_', '\n',
			'_', 'x', '_', '_', '_', 'x', '_', '_', '\n',
			'_', '_', 'x', '_', 'x', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'}}, {"Test position#", args{
			[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			39}, []rune{
			'_', '_', '_', '_', '_', '_', '_', 'x', '\n',
			'x', '_', '_', '_', '_', '_', 'x', '_', '\n',
			'_', 'x', '_', '_', '_', 'x', '_', '_', '\n',
			'_', '_', 'x', '_', 'x', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesDiagonallyAbove(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesDiagonallyAbove() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

func Test_isAnotherFigurePresentDiag(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, false},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			8}, false},
		{"Test position#", args{[]rune{
			'k', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, true},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, true},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'k', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, true},
		{"Test position#", args{[]rune{
			'_', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			10}, true},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'k', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'k', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			7}, false},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'k', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'k', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'q', '\n'},
			0}, true},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'k', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'k', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'q', '\n'},
			50}, true},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'k', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'k', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'q', '\n'},
			1}, true},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'k', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'k', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'q', '\n'},
			51}, true},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'k', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'k', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'q', '\n'},
			55}, false},
		{"Test position#", args{[]rune{
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'k', '_', 'k', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', '_', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'k', '\n',
			'_', '_', '_', '_', '_', '_', '_', 'q', '\n'},
			63}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnotherFigurePresentDiag(tt.args.out, tt.args.position); got != tt.want {
				t.Errorf("isAnotherFigurePresenDiagBelow() = %v, want %v", got, tt.want)
			}
		})
	}
}
