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
		{name: "one empty place", board: []byte("rnbqkbn_\n"), want: 1},
		{name: "empty board", board: []byte("________\n________\n________\n________\n________\n________\n________\n________\n"), want: 64},
		{name: "not empty board", board: []byte("r_______\n________\n________\n________\n________\n________\n________\n________\n"), want: 63},
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

func Test_getBoardFromPool(t *testing.T) {
	tests := []struct {
		name      string
		dimension int
		wantSize  int
	}{
		{
			name:      "dimension 1",
			dimension: 1,
			wantSize:  2, // 1 * (1+1) = 2
		},
		{
			name:      "dimension 7",
			dimension: 7,
			wantSize:  56, // 7 * (7+1) = 56
		},
		{
			name:      "dimension 8",
			dimension: 8,
			wantSize:  72, // 8 * (8+1) = 72
		},
		{
			name:      "dimension 10",
			dimension: 10,
			wantSize:  110, // 10 * (10+1) = 110
		},
		{
			name:      "dimension 20",
			dimension: 20,
			wantSize:  420, // 20 * (20+1) = 420
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boardPtr := getBoardFromPool(tt.dimension)
			if boardPtr == nil {
				t.Errorf("getBoardFromPool() returned nil")
				return
			}
			board := *boardPtr
			if len(board) != tt.wantSize {
				t.Errorf("getBoardFromPool() board size = %v, want %v", len(board), tt.wantSize)
			}
			// Return to pool for reuse
			boardPool.Put(boardPtr)
		})
	}
}

func Test_getBoardFromPool_reusesPooledBoards(t *testing.T) {
	// Get a board for dimension 8
	boardPtr1 := getBoardFromPool(8)
	board1 := *boardPtr1
	expectedSize8 := 72 // 8 * 9 = 72

	if len(board1) != expectedSize8 {
		t.Errorf("First board size = %v, want %v", len(board1), expectedSize8)
	}

	// Return it to pool
	boardPool.Put(boardPtr1)

	// Get another board for same dimension - should reuse the same pointer
	boardPtr2 := getBoardFromPool(8)

	if boardPtr1 != boardPtr2 {
		t.Errorf("Expected to reuse the same board pointer from pool")
	}

	boardPool.Put(boardPtr2)
}

func Test_getBoardFromPool_resizesWhenNeeded(t *testing.T) {
	// Get a board for dimension 8
	boardPtr := getBoardFromPool(8)
	expectedSize8 := 72 // 8 * 9 = 72

	if len(*boardPtr) != expectedSize8 {
		t.Errorf("Board size = %v, want %v", len(*boardPtr), expectedSize8)
	}

	// Return it to pool
	boardPool.Put(boardPtr)

	// Get a board for different dimension 7 - should resize
	boardPtr7 := getBoardFromPool(7)
	expectedSize7 := 56 // 7 * 8 = 56
	board7 := *boardPtr7

	if len(board7) != expectedSize7 {
		t.Errorf("Resized board size = %v, want %v", len(board7), expectedSize7)
	}

	boardPool.Put(boardPtr7)
}
