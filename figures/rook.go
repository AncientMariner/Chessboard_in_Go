package figures

import "github.com/hashicorp/go-set/v2"

type Rook struct {
	Figure
}

func (rook *Rook) Handle(request string) *set.HashSet[*FigurePosition, string] {
	if rook.next != nil {
		rook.next.Handle(request)
	}
	return set.NewHashSet[*FigurePosition, string](0)

}

func (*Rook) GetName() rune {
	return 'r'
}
