package figures

import "github.com/hashicorp/go-set/v2"

type Figure struct {
	next FigureBehaviour
}

type FigureBehaviour interface {
	SetNext(FigureBehaviour) FigureBehaviour
	Handle(string) *set.HashSet[*FigurePosition, string]
	GetName() rune
	GetNext() FigureBehaviour
}

func (f *Figure) SetNext(next FigureBehaviour) FigureBehaviour {
	f.next = next
	return next
}

// func (f *Figure) Handle(request string) *set.HashSet[*FigurePosition, string] {
// 	if f.next != nil {
// 		return f.next.Handle(request)
// 	}
// 	return nil
// }

func (f *Figure) GetNext() FigureBehaviour {
	return f.next
}
