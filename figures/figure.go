package figures

type Figure struct {
	next FigureBehaviour
}

type FigureBehaviour interface {
	SetNext(FigureBehaviour) FigureBehaviour
	Handle(string) map[uint32]string
	GetName() rune
	GetNext() FigureBehaviour
}

func (f *Figure) SetNext(next FigureBehaviour) FigureBehaviour {
	f.next = next
	return next
}

// func (f *Figure) Handle(request string) *set.HashSet[*BoardWithFigurePosition, string] {
// 	if f.next != nil {
// 		return f.next.Handle(request)
// 	}
// 	return nil
// }

func (f *Figure) GetNext() FigureBehaviour {
	return f.next
}
