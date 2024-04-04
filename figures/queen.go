package figures

import "github.com/hashicorp/go-set/v2"

type Queen struct {
	Figure
	// add placement behaviour
}

func (queen *Queen) Handle(request string) *set.HashSet[*FigurePosition, string] {
	if request == "Hello" {
		return nil
	}
	if queen.next != nil {
		queen.next.Handle(request)
	}
	return set.NewHashSet[*FigurePosition, string](0)
}

func (*Queen) GetName() rune {
	return 'q'
}
