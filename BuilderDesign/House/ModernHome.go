package House

type modernHome struct {
	gate, window int
}

func (mh *modernHome) setNumOfWindow(window int) {
	//TODO implement me
	panic("implement me")
}

func (mh *modernHome) setNumOfGate(gate int) {
	//TODO implement me
	panic("implement me")
}

func newModernHome() *modernHome {
	return &modernHome{}
}

func (sh *modernHome) getHome() *home {
	return &home{
		gate:   sh.gate,
		window: sh.window,
	}
}
