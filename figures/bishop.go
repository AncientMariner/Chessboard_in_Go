package figures

type Bishop struct {
	Figure
}

func (s *Bishop) Handle(request string) string {
	if request == "Hello" {
		return "World"
	}
	return s.next.Handle(request)
}

//
// func (s *Bishop) setNext(next FigureBehaviour) FigureBehaviour {
// 	s.next = next
// 	return next
// }
func (*Bishop) GetName() string {
	return "bishop"
}
