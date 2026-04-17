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
		want   byte
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
		board []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test handle board size 8 is not possible", fields{Figure{next: nil}}, args{board: []byte("________")}, 0},
		{"Test handle empty board size 64", fields{Figure{next: nil}}, args{board: []byte("________________________________________________________________")}, 64},
		{"Test handle board with 1 king size 40", fields{Figure{next: nil}}, args{board: []byte("xxx_____xkx_____xxx_____________________________________________")}, 40},
		{"Test handle board with 1 king size 40", fields{Figure{next: nil}}, args{board: []byte("________xxx_____xkx_____xxx_____________________________________")}, 40},
		{"Test handle board with 2 king size 22", fields{Figure{next: nil}}, args{board: []byte("________xxx_____xkx_____xxx__________________xxx_____xkx_____xxx")}, 22},
		{"Test handle board with 2 king size 13", fields{Figure{next: nil}}, args{board: []byte("________xxx_____xkx_____xxx_____xkx_____xxx__xxx_____xkx_____xxx")}, 13},
		{"Test handle board with 2 king size 22", fields{Figure{next: nil}}, args{board: []byte("xkx___xkxxx___xxxkx_____xxx_____xkx_____xxx__xxx_____xkx_____xxx")}, 6},
		{"Test handle board with 2 king size 4", fields{Figure{next: nil}}, args{board: []byte("xkxxxxxkxxxxkxxxxkxxxx__xxx_____xkx_____xxx__xxx_____xkx_____xxx")}, 4},
		{"Test handle board with 2 king size 3", fields{Figure{next: nil}}, args{board: []byte("xkxxxxxkxxxxkxxxxkxxxx__xxx_____xkx_____xxx__xxxxkx__xkxxxx__xxx")}, 3},
		{"Test handle board with 2 king size 2", fields{Figure{next: nil}}, args{board: []byte("xkxxxxxkxxxxkxxxxkxxxxxxxxx__xkxxkx__xxxxxx__xxxxkx__xkxxxx__xxx")}, 2},
		{"Test handle board with 2 king size 0", fields{Figure{next: nil}}, args{board: []byte("xkxxxxxkxxxxkxxxxkxxxxxxxxxxkxkxxkxxxxxxxxxxxxxxxkxxkxkxxxxxxxxx")}, 0},
		{"Test handle board with 1 king 1 rook 1 bishop size 0", fields{Figure{next: nil}}, args{board: []byte("_x______xrxxxxxxxx___xxx_x___xkx_xx__xxx_x_x_____x__x__x_x____b_")}, 11},

		{"Test handle board with 7 queens", fields{Figure{next: nil}},
			args{board: []byte("xxqxxxxx" +
				"_xxxxxxx" +
				"xxxxxxqx" +
				"xxxxqxxx" +
				"xxxxxxxq" +
				"xqxxxxxx" +
				"xxxqxxxx" +
				"xxxxxqxx")}, 1},
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
