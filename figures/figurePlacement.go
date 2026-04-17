package figures

import (
	"runtime"
	"sync"
)

type Placement struct {
}

// GenerateHash returns a uint64 hash for efficient map keys (no allocations)
// Uses direct FNV-1a implementation for maximum performance
func GenerateHash(s []byte) uint64 {
	// FNV-1a hash algorithm constants
	const offset64 uint64 = 14695981039346656037
	const prime64 uint64 = 1099511628211

	h := offset64
	for _, b := range s {
		h ^= uint64(b)
		h *= prime64
	}
	return h
}

var defaultDimension = 8

// numWorkers determines the worker pool size for parallel processing
// Defaults to number of CPU cores for optimal performance
var numWorkers = runtime.GOMAXPROCS(0)

// getParallelThreshold calculates optimal threshold for parallel vs sequential processing
// Returns minimum number of boards needed to justify parallel processing overhead
func getParallelThreshold() int {
	// Dynamic threshold: at least 10 boards per worker to justify overhead
	// But never less than 10 total to avoid goroutine overhead for tiny workloads
	minPerWorker := 10
	threshold := numWorkers * minPerWorker
	if threshold < 10 {
		threshold = 10
	}
	return threshold
}

const emptyField = '_'
const attackPlace = 'x'

func drawEmptyBoard() []byte {
	board := make([]byte, defaultDimension*defaultDimension)
	for i := range board {
		board[i] = emptyField
	}
	return board
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
	threshold := getParallelThreshold()
	if len(boards) >= threshold {
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
