package figures

type King struct {
	Figure
	// add placement behaviour
	// add figure chain
}

// func (k *King) NewKing(a map[string]int) FigureChain {
// 	k.fc.figureQuantityMap = a
// 	return k.fc
// }

func (s *King) Handle(request string) string {
	if request == "Hello" {
		return "World"
	}
	return s.next.Handle(request)
}

func (s *King) setNext(next FigureBehaviour) FigureBehaviour {
	s.next = next
	return next
}

func (*King) getName() string {
	return "king"
}
