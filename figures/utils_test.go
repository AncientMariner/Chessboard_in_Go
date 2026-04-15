package figures

import "testing"

func Test_getCountOfEmptyPlaces(t *testing.T) {
	tests := []struct {
		name  string
		board []byte
		want  int
	}{
		{name: "empty string", board: make([]byte, 0), want: 0}, 
		{name: "empty string", board: make([]byte, 1), want: 0}, 
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

func Test_boardPool(t *testing.T) {
	board := boardPool.Get().(*[]byte)
	if board == nil {
		t.Error("Expected a board from the pool, got nil")
	}

	boardPool.Put(board)

	board2 := boardPool.Get().(*[]byte)
	if board2 == nil {
		t.Error("Expected a board from the pool, got nil")
	}

	if board != board2 {
		t.Error("Expected to get the same board from the pool, got a different one")
	}
}

