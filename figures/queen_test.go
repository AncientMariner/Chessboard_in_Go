package figures

import (
	"testing"
)

func TestQueen_GetName(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{"Test get name", fields{Figure{}}, 'q'},
		{"Test get name with empty figure", fields{}, 'q'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &Queen{
				Figure: tt.fields.Figure,
			}
			if got := qu.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueen_Handle(t *testing.T) {
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
		{"Test handle board with 1 king size 40", fields{Figure{next: nil}}, args{board: "xxx_____\nxkx_____\nxxx_____\n________\n________\n________\n________\n________\n"}, 40},
		{"Test handle board with 1 king size 40", fields{Figure{next: nil}}, args{board: "________\nxxx_____\nxkx_____\nxxx_____\n________\n________\n________\n________\n"}, 40},
		{"Test handle board with 2 king size 22", fields{Figure{next: nil}}, args{board: "________\nxxx_____\nxkx_____\nxxx_____\n________\n_____xxx\n_____xkx\n_____xxx\n"}, 22},
		{"Test handle board with 2 king size 13", fields{Figure{next: nil}}, args{board: "________\nxxx_____\nxkx_____\nxxx_____\nxkx_____\nxxx__xxx\n_____xkx\n_____xxx\n"}, 13},
		{"Test handle board with 2 king size 22", fields{Figure{next: nil}}, args{board: "xkx___xk\nxxx___xx\nxkx_____\nxxx_____\nxkx_____\nxxx__xxx\n_____xkx\n_____xxx\n"}, 6},
		{"Test handle board with 2 king size 4", fields{Figure{next: nil}}, args{board: "xkxxxxxk\nxxxxkxxx\nxkxxxx__\nxxx_____\nxkx_____\nxxx__xxx\n_____xkx\n_____xxx\n"}, 4},
		{"Test handle board with 2 king size 3", fields{Figure{next: nil}}, args{board: "xkxxxxxk\nxxxxkxxx\nxkxxxx__\nxxx_____\nxkx_____\nxxx__xxx\nxkx__xkx\nxxx__xxx\n"}, 3},
		{"Test handle board with 2 king size 2", fields{Figure{next: nil}}, args{board: "xkxxxxxk\nxxxxkxxx\nxkxxxxxx\nxxx__xkx\nxkx__xxx\nxxx__xxx\nxkx__xkx\nxxx__xxx\n"}, 2},
		{"Test handle board with 2 king size 0", fields{Figure{next: nil}}, args{board: "xkxxxxxk\nxxxxkxxx\nxkxxxxxx\nxxxxkxkx\nxkxxxxxx\nxxxxxxxx\nxkxxkxkx\nxxxxxxxx\n"}, 0},
		{"Test handle board with 1 king 1 rook 1 bishop size 0", fields{Figure{next: nil}}, args{board: "_x______\nxrxxxxxx\nxx___xxx\n_x___xkx\n_xx__xxx\n_x_x____\n_x__x__x\n_x____b_\n"}, 11},

		{"Test handle board with 7 queens", fields{Figure{next: nil}},
			args{board: "xxqxxxxx\n" +
				"_xxxxxxx\n" +
				"xxxxxxqx\n" +
				"xxxxqxxx\n" +
				"xxxxxxxq\n" +
				"xqxxxxxx\n" +
				"xxxqxxxx\n" +
				"xxxxxqxx\n"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queen := &Queen{
				Figure: tt.fields.Figure,
			}
			if got := queen.Handle(tt.args.board); len(got) != tt.want {
				t.Errorf("Handle() size= %v, want %v", len(got), tt.want)
			}
		})
	}
}

// qxxxxxxx
// xxxxxxqx
// xxxqxxxx
// xxxxxqxx
// xxxxxxxq
// xqxxxxxx
// xxxxqxxx
// xxqxxxxx
