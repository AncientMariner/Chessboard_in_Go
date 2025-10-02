package figures

import "testing"

func Test_getCountOfEmptyPlaces(t *testing.T) {
	tests := []struct {
		name  string
		board string
		want  int
	}{
		{name: "empty string", board: "", want: 0},
		{name: "no empty places", board: "rnbqkbnr\n", want: 0},
		{name: "one empty place", board: "rnbqkbn_\n", want: 1},
		{name: "empty board", board: "________\n________\n________\n________\n________\n________\n________\n________\n", want: 64},
		{name: "not empty board", board: "r_______\n________\n________\n________\n________\n________\n________\n________\n", want: 63},
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
