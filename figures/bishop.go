package figures

import "github.com/hashicorp/go-set/v2"

type Bishop struct {
	Figure
}

func (bishop *Bishop) Handle(request string) *set.HashSet[*FigurePosition, string] {
	if request == "Hello" {
		return nil
	}
	if bishop.next != nil {
		bishop.next.Handle(request)
	}
	return set.NewHashSet[*FigurePosition, string](0)

}

func (*Bishop) GetName() rune {
	return 'b'
}
