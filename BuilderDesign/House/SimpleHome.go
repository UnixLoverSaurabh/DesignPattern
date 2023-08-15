package House

type simpleHome struct {
	gate, window int
}

func (sh *simpleHome) setNumOfWindow(window int) {
	//TODO implement me
	panic("implement me")
}

func (sh *simpleHome) setNumOfGate(gate int) {
	//TODO implement me
	panic("implement me")
}

func newSimpleHome() *simpleHome {
	return &simpleHome{}
}

func (sh *simpleHome) getHome() *home {
	return &home{
		gate:   sh.gate,
		window: sh.window,
	}
}
