package figures

type Queen struct {
	Figure
	// add placement behaviour
}

func (s *Queen) Handle(request string) string {
	if request == "Hello" {
		return "World"
	}
	return s.next.Handle(request)
}

func (*Queen) GetName() rune {
	return 'q'
}
