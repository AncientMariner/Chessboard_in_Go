package figures

import (
	"hash"
	"hash/fnv"
	"strings"
	"sync"
)

type Placement struct {
}

// hashPool reuses FNV hash functions to avoid repeated allocations
var hashPool = sync.Pool{
	New: func() interface{} {
		return fnv.New64a()
	},
}

// GenerateHash returns a uint64 hash for efficient map keys (no allocations)
// Uses pooled hash functions to avoid creating new hashers on each call
func GenerateHash(s []byte) uint64 {
	h := hashPool.Get().(hash.Hash64)
	h.Reset() // Reset the hasher state before reuse
	h.Write(s)
	sum := h.Sum64()
	hashPool.Put(h) // Return hasher to pool
	return sum
}

var defaultDimension = 8

const emptyField = '_'
const attackPlace = 'x'

func drawEmptyBoard() []byte {

	var board strings.Builder
	board.Grow(defaultDimension*defaultDimension + defaultDimension)

	for x := 0; x < defaultDimension; x++ {
		for y := 0; y < defaultDimension; y++ {
			board.WriteByte(emptyField)
		}
		board.WriteByte('\n')
	}
	return []byte(board.String())
}

func (p *Placement) SetDimension(value int) {
	defaultDimension = value
}

func (p *Placement) PlaceFigures(numberOfFigures int, behaviour FigureBehaviour, boards map[uint64][]byte) map[uint64][]byte {
	for i := 0; i < numberOfFigures; i++ {
		if len(boards) == 0 {
			boards = p.placeFigureOnBoard(drawEmptyBoard(), behaviour)
		} else {
			boards = p.placeFigure(boards, behaviour)
		}
	}
	return boards
}

func (p *Placement) placeFigure(boards map[uint64][]byte, behaviour FigureBehaviour) map[uint64][]byte {
	var resultingMap = make(map[uint64][]byte)

	for _, board := range boards {

		boardsWithPlacement := p.placeFigureOnBoard(board, behaviour)

		for hash, element := range boardsWithPlacement {
			resultingMap[hash] = element
		}
	}
	return resultingMap
}

func (p *Placement) placeFigureOnBoard(board []byte, behaviour FigureBehaviour) map[uint64][]byte {
	return behaviour.Handle(board)
}
