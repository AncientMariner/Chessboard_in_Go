package figures

type King struct {
	Figure
	// add placement behaviour
	// add figure chain
}

func (s *King) Handle(request string) string {
	if request == "Hello" {
		return "World"
	}
	return s.next.Handle(request)
}

func (*King) GetName() rune {
	return 'k'
}
