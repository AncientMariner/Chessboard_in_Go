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
		want   byte
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
		board []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test handle board size 8 is not possible", fields{Figure{next: nil}}, args{board: []byte("________")}, 0},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: []byte(
			"________" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________")}, 64},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: []byte(
			"xxx_____" +
				"xkx_____" +
				"xxx_____" +
				"________" +
				"________" +
				"________" +
				"________" +
				"________")}, 51},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: []byte(
			"xxx_____" +
				"xkx_____" +
				"xxx_____" +
				"___xxx__" +
				"___xkx__" +
				"___xxx__" +
				"________" +
				"________")}, 36},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: []byte(
			"xxx_____" +
				"xkx_____" +
				"xxx_____" +
				"___xxx__" +
				"___xkx__" +
				"___xxx__" +
				"xxx__xxx" +
				"xkx__xkx")}, 22},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: []byte(
			"xxx__xxx" +
				"xkx__xkx" +
				"xxx__xxx" +
				"xxxxxx__" +
				"xkxxkx__" +
				"xxxxxx__" +
				"xxx__xxx" +
				"xkx__xkx")}, 6},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: []byte(
			"xxx__xxx" +
				"xkx__xkx" +
				"xxx__xxx" +
				"xxxxxx__" +
				"xkxxkxxx" +
				"xxxxxxxk" +
				"xxxxxxxx" +
				"xkxkxxkx")}, 2},
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
		out      []byte
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test vertically 1 above", args{
			out: []byte{
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 5,
		}, false},
		{"Test vertically 1 above", args{
			out: []byte{
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 10,
		}, false},
		{"Test vertically 1 above", args{
			out: []byte{
				'b', '_', '_', 'k', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 10,
		}, true},
		{"Test vertically 1 above", args{
			out: []byte{
				'_', 'b', '_', '_', '_', 'k', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 11,
		}, true},
		{"Test vertically 1 above", args{
			out: []byte{
				'b', '_', '_', '_', 'k', 'q', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 15,
		}, true},
		{"Test vertically 1 above", args{
			out: []byte{
				'_', 'b', '_', 'b', '_', 'q', '_', '_',
				'k', '_', '_', '_', 'k', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 20,
		}, true},
		{"Test vertically 1 above", args{
			out: []byte{
				'_', 'b', 'k', 'b', '_', 'q', '_', '_',
				'k', '_', 'r', '_', 'k', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 18,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnotherFigurePresentAbove(tt.args.out, tt.args.position, 8); got != tt.want {
				t.Errorf("isAnotherFigurePresentAbove() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

func Test_isAnotherFigurePresentBelow(t *testing.T) {
	type args struct {
		out      []byte
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test vertically 1 below", args{
			out: []byte{
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 5,
		}, false},
		{"Test vertically 1 below", args{
			out: []byte{
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 10,
		}, false},
		{"Test vertically 1 below", args{
			out: []byte{
				'b', '_', '_', 'k', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 10,
		}, false},
		{"Test vertically 1 below", args{
			out: []byte{
				'b', '_', '_', '_', 'k', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 11,
		}, false},
		{"Test vertically 1 below", args{
			out: []byte{
				'b', '_', '_', '_', 'k', 'q', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 16,
		}, false},
		{"Test vertically 1 below", args{
			out: []byte{
				'_', 'b', '_', 'b', '_', 'q', '_', '_',
				'k', '_', '_', '_', 'k', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 20,
		}, false},
		{"Test vertically 1 below", args{
			out: []byte{
				'_', 'b', 'k', 'b', '_', 'q', '_', '_',
				'k', '_', 'r', '_', 'k', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 18,
		}, false},
		{"Test vertically 1 below", args{
			out: []byte{
				'_', 'b', 'k', 'b', '_', '_', '_', '_',
				'_', '_', 'b', '_', '_', '_', 'b', '_',
				'_', '_', '_', 'b', '_', 'b', '_', '_',
			},
			position: 4,
		}, true},
		{"Test vertically 1 below", args{
			out: []byte{
				'_', 'b', 'k', 'b', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', 'b', '_',
				'_', '_', '_', 'b', '_', 'b', '_', '_',
			},
			position: 4,
		}, true},
		{"Test vertically 1 below", args{
			out: []byte{
				'_', 'b', 'k', 'b', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', 'b', '_', 'b', '_', '_',
			},
			position: 4,
		}, true},
		{"Test vertically 1 below", args{
			out: []byte{
				'_', 'b', 'k', 'b', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', 'b', '_', '_',
			},
			position: 4,
		}, true},
		{"Test vertically 1 below", args{
			out: []byte{
				'_', 'b', 'k', 'b', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 4,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnotherFigurePresentBelow(tt.args.out, tt.args.position, 8); got != tt.want {
				t.Errorf("isAnotherFigurePresentBelow() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

func Test_placeAttackPlacesBelow(t *testing.T) {
	type args struct {
		out      []byte
		position int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Test vertically 1 below", args{
			out: []byte{
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 2,
		}, []byte{
			'_', '_', '_', '_', '_', '_', '_', '_',
			'x', '_', '_', '_', 'x', '_', '_', '_',
			'_', 'x', '_', 'x', '_', '_', '_', '_',
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesBelow(tt.args.out, tt.args.position, 8)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesBelow() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

func Test_placeAttackPlacesAbove(t *testing.T) {
	type args struct {
		out      []byte
		position int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Test above", args{
			out: []byte{
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 18,
		}, []byte{
			'_', 'x', '_', 'x', '_', '_', '_', '_',
			'x', '_', '_', '_', 'x', '_', '_', '_',
			'_', '_', '_', '_', '_', '_', '_', '_',
		}},

		{"Test below", args{
			out: []byte{
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
				'_', '_', '_', '_', '_', '_', '_', '_',
			},
			position: 16,
		}, []byte{
			'_', 'x', '_', '_', '_', '_', '_', '_',
			'_', '_', 'x', '_', '_', '_', '_', '_',
			'_', '_', '_', '_', '_', '_', '_', '_',
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesAbove(tt.args.out, tt.args.position, 8)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesAbove() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}
