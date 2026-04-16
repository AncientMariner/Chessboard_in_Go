package figures

import (
	"hash"
	"hash/fnv"
	"runtime"
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

// parallelThreshold defines minimum number of boards to process in parallel
// Below this threshold, sequential processing is more efficient due to goroutine overhead
const parallelThreshold = 10

// numWorkers determines the worker pool size for parallel processing
// Defaults to number of CPU cores for optimal performance
var numWorkers = runtime.GOMAXPROCS(0)

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
	// Use parallel processing if we have enough boards to justify the overhead
	if len(boards) >= parallelThreshold {
		return p.placeFigureParallel(boards, behaviour)
	}

	// Sequential processing for small workloads
	return p.placeFigureSequential(boards, behaviour)
}

// placeFigureSequential processes boards sequentially (original implementation)
func (p *Placement) placeFigureSequential(boards map[uint64][]byte, behaviour FigureBehaviour) map[uint64][]byte {
	resultingMap := getMapFromPool(len(boards) * 10) // Rough estimate of result size

	for _, board := range boards {

		boardsWithPlacement := p.placeFigureOnBoard(board, behaviour)

		for hash, element := range boardsWithPlacement {
			resultingMap[hash] = element
		}

		// Return the temporary map from Handle() to the pool
		putMapToPool(boardsWithPlacement)
	}
	return resultingMap
}

// placeFigureParallel processes boards in parallel using a worker pool pattern
func (p *Placement) placeFigureParallel(boards map[uint64][]byte, behaviour FigureBehaviour) map[uint64][]byte {
	// Create channels for work distribution
	type job struct {
		board []byte
	}

	type result struct {
		boards map[uint64][]byte
	}

	jobs := make(chan job, len(boards))
	results := make(chan result, len(boards))

	// Start worker pool
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				// Process the board
				boardsWithPlacement := p.placeFigureOnBoard(j.board, behaviour)
				results <- result{boards: boardsWithPlacement}
			}
		}()
	}

	// Send jobs to workers
	for _, board := range boards {
		jobs <- job{board: board}
	}
	close(jobs)

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results - use mutex to protect the shared map
	resultingMap := getMapFromPool(len(boards) * 10)
	var mu sync.Mutex

	for res := range results {
		mu.Lock()
		for hash, element := range res.boards {
			resultingMap[hash] = element
		}
		mu.Unlock()

		// Return the temporary map from Handle() to the pool
		putMapToPool(res.boards)
	}

	return resultingMap
}

func (p *Placement) placeFigureOnBoard(board []byte, behaviour FigureBehaviour) map[uint64][]byte {
	return behaviour.Handle(board)
}
