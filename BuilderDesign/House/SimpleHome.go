package House

type simpleHome struct {
	gate, window int
}

func (sh *simpleHome) SetNumOfWindow() {
	sh.window = 2
}

func (sh *simpleHome) SetNumOfGate() {
	sh.gate = 1
}

func newSimpleHome() *simpleHome {
	return &simpleHome{}
}

func (sh *simpleHome) GetHome() *Home {
	return &Home{
		gate:   sh.gate,
		window: sh.window,
	}
}
