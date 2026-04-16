package figures

var _ FigureBehaviour = (*King)(nil)

type King struct {
	Figure
}

func (king *King) Handle(board []byte) map[uint64][]byte {
	boards := getMapFromPool(getCountOfEmptyPlaces(board))

	for i := 0; i < len(board) && len(board) == ((defaultDimension+1)*defaultDimension); i++ {
		if board[i] == emptyField {
			// Get from pool
			outPtr := boardPool.Get().(*[]byte)
			out := *outPtr
			copy(out, board)

			if !isAnotherFigurePresent(out, i) {
				king.placeAttackPlacesHorizontally(out, i)
				king.placeAttackPlacesVertically(out, i)
				king.placeDiagonallyAbove(out, i)
				king.placeDiagonallyBelow(out, i)
				out[i] = king.GetName()

				// Make permanent copy for the map
				permanent := make([]byte, len(out))
				copy(permanent, out)
				boards[GenerateHash(permanent)] = permanent
			}

			// Always return to pool
			boardPool.Put(outPtr)
		}
	}
	return boards
}

func isAnotherFigurePresent(out []byte, position int) bool {

	positionOneLineAbove := position - defaultDimension - 1
	var positionsAround []int

	diagAboveRight := positionOneLineAbove + 1
	previousLineExists := position >= defaultDimension+1

	if previousLineExists && out[diagAboveRight] != '\n' {
		positionsAround = append(positionsAround, diagAboveRight)
	}
	diagAboveLeft := positionOneLineAbove - 1
	if previousLineExists && (position-1)%defaultDimension != 0 && diagAboveLeft >= 0 && out[diagAboveLeft] != '\n' {
		positionsAround = append(positionsAround, diagAboveLeft)
	}

	diagBelowRight := position + defaultDimension + 1 + 1
	diagBelowLeft := position + defaultDimension + 1 - 1
	isNotLastLine := position < len(out)-defaultDimension-1

	if isNotLastLine && diagBelowRight < len(out) && out[diagBelowRight] != '\n' {
		positionsAround = append(positionsAround, diagBelowRight)
	}
	if isNotLastLine && position%defaultDimension != 0 && diagBelowLeft < len(out) && out[diagBelowLeft] != '\n' {
		positionsAround = append(positionsAround, diagBelowLeft)
	}

	previousPosition := position - 1
	if previousPosition >= 0 && out[previousPosition] != '\n' {
		positionsAround = append(positionsAround, previousPosition)
	}
	nextPosition := position + 1
	if nextPosition < len(out) && out[nextPosition] != '\n' {
		positionsAround = append(positionsAround, nextPosition)
	}

	positionAbove := position - defaultDimension - 1
	if previousLineExists && out[positionAbove] != '\n' {
		positionsAround = append(positionsAround, positionAbove)
	}
	positionBelow := position + defaultDimension + 1
	if position < len(out)-defaultDimension-1 && out[positionBelow] != '\n' {
		positionsAround = append(positionsAround, positionBelow)
	}

	for _, number := range positionsAround {
		if out[number] != emptyField && out[number] != attackPlace && out[number] != '\n' {
			return true
		}
	}
	return false
}

func (king *King) placeDiagonallyAbove(out []byte, position int) {
	if position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	positionOneLineAbove := position - defaultDimension - 1

	diagAboveRight := positionOneLineAbove + 1
	previousLineExists := position >= defaultDimension+1

	if previousLineExists && out[diagAboveRight] == emptyField {
		out[diagAboveRight] = attackPlace
	}
	diagAboveLeft := positionOneLineAbove - 1
	if previousLineExists && (position-1)%defaultDimension != 0 && diagAboveLeft >= 0 && out[diagAboveLeft] == emptyField {
		out[diagAboveLeft] = attackPlace
	}
}

func (king *King) placeDiagonallyBelow(out []byte, position int) {
	if position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	diagBelowRight := position + defaultDimension + 1 + 1
	diagBelowLeft := position + defaultDimension + 1 - 1
	isNotLastLine := position < len(out)-defaultDimension-1

	if isNotLastLine && diagBelowRight < len(out) && out[diagBelowRight] == emptyField {
		out[diagBelowRight] = attackPlace
	}
	if isNotLastLine && position%defaultDimension != 0 && diagBelowLeft < len(out) && out[diagBelowLeft] == emptyField {
		out[diagBelowLeft] = attackPlace
	}
}

func (king *King) placeAttackPlacesVertically(out []byte, position int) {
	positionAbove := position - defaultDimension - 1
	if position >= defaultDimension+1 && out[positionAbove] == emptyField {
		out[positionAbove] = attackPlace
	}
	positionBelow := position + defaultDimension + 1
	if position < len(out)-defaultDimension-1 && out[positionBelow] == emptyField {
		out[positionBelow] = attackPlace
	}
}

func (king *King) placeAttackPlacesHorizontally(out []byte, position int) {
	if position == defaultDimension || position%(defaultDimension+1) == defaultDimension {
		return
	}
	previousPosition := position - 1
	if previousPosition >= 0 && out[previousPosition] == emptyField {
		out[previousPosition] = attackPlace
	}
	nextPosition := position + 1
	if nextPosition < len(out) && out[nextPosition] == emptyField {
		out[nextPosition] = attackPlace
	}
}

func (*King) GetName() byte {
	return 'k'
}
