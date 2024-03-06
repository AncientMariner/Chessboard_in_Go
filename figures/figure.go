package figures

type Figure struct {
	next FigureBehaviour
}

type FigureBehaviour interface {
	SetNext(FigureBehaviour) FigureBehaviour
	Handle(string) string
	GetName() string
}

func (f *Figure) SetNext(next FigureBehaviour) FigureBehaviour {
	f.next = next
	return next
}

func (f *Figure) Handle(request string) string {
	if f.next != nil {
		return f.next.Handle(request)
	}
	return ""
}
