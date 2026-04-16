package figures

import (
	"reflect"
	"testing"
)

func TestFigure_GetNext(t *testing.T) {
	type fields struct {
		next FigureBehaviour
	}
	tests := []struct {
		name   string
		fields fields
		want   FigureBehaviour
	}{
		{"Test get next nil", fields{nil}, nil},
		{"Test get next nil", fields{&King{Figure{next: nil}}}, &King{Figure{next: nil}}},
		{"Test get next nil", fields{&Queen{Figure{next: nil}}}, &Queen{Figure{next: nil}}},
		{"Test get next nil", fields{&Bishop{Figure{next: nil}}}, &Bishop{Figure{next: nil}}},
		{"Test get next nil", fields{&Rook{Figure{next: nil}}}, &Rook{Figure{next: nil}}},
		{"Test get next nil", fields{&Knight{Figure{next: nil}}}, &Knight{Figure{next: nil}}},
		{"Test get next nil", fields{&Knight{Figure{next: &King{Figure{next: &Queen{}}}}}}, &Knight{Figure{next: &King{Figure{next: &Queen{}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Figure{
				next: tt.fields.next,
			}
			if got := f.GetNext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestFigure_GetNextGetNext(t *testing.T) {
	type fields struct {
		next FigureBehaviour
	}
	tests := []struct {
		name   string
		fields fields
		want   FigureBehaviour
	}{
		{"Test get next nil", fields{&Knight{Figure{next: &King{Figure{next: &Queen{}}}}}}, &King{Figure{next: &Queen{}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Figure{
				next: tt.fields.next,
			}
			if got := f.GetNext().GetNext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFigure_SetNext(t *testing.T) {
	type fields struct {
		next FigureBehaviour
	}
	type args struct {
		next FigureBehaviour
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   FigureBehaviour
	}{
		{"Test set next", fields{&Knight{Figure{next: nil}}}, args{&Queen{Figure{next: nil}}}, &Queen{Figure{next: nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Figure{
				next: tt.fields.next,
			}
			if got := f.SetNext(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDimensionFromBoard(t *testing.T) {
	tests := []struct {
		name  string
		board []byte
		want  int
	}{
		{
			name:  "empty board - invalid",
			board: []byte{},
			want:  8, // Default fallback
		},
		{
			name:  "1x1 board (2 bytes: 1 char + 1 newline)",
			board: []byte("_\n"),
			want:  1,
		},
		{
			name:  "2x2 board (6 bytes: 2*(2+1))",
			board: []byte("__\n__\n"),
			want:  2,
		},
		{
			name:  "3x3 board (12 bytes: 3*(3+1))",
			board: []byte("___\n___\n___\n"),
			want:  3,
		},
		{
			name:  "7x7 board (56 bytes: 7*(7+1))",
			board: []byte("_______\n_______\n_______\n_______\n_______\n_______\n_______\n"),
			want:  7,
		},
		{
			name:  "8x8 board (72 bytes: 8*(8+1))",
			board: []byte("________\n________\n________\n________\n________\n________\n________\n________\n"),
			want:  8,
		},
		{
			name:  "10x10 board (110 bytes: 10*(10+1))",
			board: make([]byte, 110), // 10 * 11 = 110
			want:  10,
		},
		{
			name:  "invalid size - fallback to default",
			board: []byte("invalid"),
			want:  8, // Default fallback
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDimensionFromBoard(tt.board); got != tt.want {
				t.Errorf("getDimensionFromBoard() = %v, want %v (board length = %d)", got, tt.want, len(tt.board))
			}
		})
	}
}
