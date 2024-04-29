package figures

type Figure struct {
	next FigureBehaviour
}

type FigureBehaviour interface {
	SetNext(FigureBehaviour) FigureBehaviour
	Handle(string) map[string]string
	GetName() rune
	GetNext() FigureBehaviour
}

func (f *Figure) SetNext(next FigureBehaviour) FigureBehaviour {
	f.next = next
	return next
}

func (f *Figure) GetNext() FigureBehaviour {
	return f.next
}
