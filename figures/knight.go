package figures

type Knight struct {
	Figure
	// add placement behaviour
}

func (s *Knight) Handle(request string) string {
	if request == "Hello" {
		return "World"
	}
	return s.next.Handle(request)
}

func (*Knight) GetName() rune {
	return 'n'
}
