package figures

type Knight struct {
	Figure
	// add placement behaviour
}

func (knight *Knight) Handle(request string) map[uint32]string {
	if request == "Hello" {
		return nil
	}
	if knight.next != nil {
		knight.next.Handle(request)
	}
	return nil
}

func (*Knight) GetName() rune {
	return 'n'
}
