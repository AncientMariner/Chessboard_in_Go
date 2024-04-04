package figures

import "github.com/hashicorp/go-set/v2"

type Knight struct {
	Figure
	// add placement behaviour
}

func (knight *Knight) Handle(request string) *set.HashSet[*FigurePosition, string] {
	if request == "Hello" {
		return nil
	}
	if knight.next != nil {
		knight.next.Handle(request)
	}
	return set.NewHashSet[*FigurePosition, string](0)
}

func (*Knight) GetName() rune {
	return 'n'
}
