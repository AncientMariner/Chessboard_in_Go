package figures

import "testing"

func Test_getCountOfEmptyPlaces(t *testing.T) {
	tests := []struct {
		name  string
		board []byte
		want  int
	}{
		{name: "empty string", board: []byte(""), want: 0},
		{name: "no empty places", board: []byte("rnbqkbnr\n"), want: 0},
		{name: "one empty place", board:[]byte("rnbqkbn_\n"), want: 1},
		{name: "empty board", board:[]byte("________\n________\n________\n________\n________\n________\n________\n________\n"), want: 64},
		{name: "not empty board", board:[]byte("r_______\n________\n________\n________\n________\n________\n________\n________\n"), want: 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCountOfEmptyPlaces(tt.board)

			if got != tt.want {
				t.Errorf("getCountOfEmptyPlaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
