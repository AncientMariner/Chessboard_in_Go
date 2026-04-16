package figures

type Figure struct {
	next FigureBehaviour
}

type FigureBehaviour interface {
	SetNext(FigureBehaviour) FigureBehaviour
	Handle([]byte) map[uint64][]byte
	GetName() byte
	GetNext() FigureBehaviour
}

func (f *Figure) SetNext(next FigureBehaviour) FigureBehaviour {
	f.next = next
	return next
}

func (f *Figure) GetNext() FigureBehaviour {
	return f.next
}

// getDimensionFromBoard calculates the board dimension from the board byte slice
// Board format: each row is dimension bytes + 1 newline, total dimension rows
// So total length = dimension * (dimension + 1)
func getDimensionFromBoard(board []byte) int {
	// Solve: len = d * (d + 1)
	// This gives us d^2 + d - len = 0
	// Using quadratic formula: d = (-1 + sqrt(1 + 4*len)) / 2
	length := len(board)
	// Simple approach: try values until we find the right one
	for d := 1; d <= 100; d++ {
		if d*(d+1) == length {
			return d
		}
	}
	return 8 // Default fallback
}
