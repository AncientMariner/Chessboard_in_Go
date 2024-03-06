package figures

type Rook struct {
	Figure
}

func (f *Rook) Handle(request string) string {
	// this one will not handle any requests
	return f.next.Handle(request)
}

// func (f *Rook) setNext(next FigureBehaviour) FigureBehaviour {
// 	f.next = next
// 	return next
// }

func (*Rook) GetName() string {
	return "rook"
}
