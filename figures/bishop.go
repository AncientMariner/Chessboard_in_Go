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

func (*Bishop) GetName() string {
    return "bishop"
}
