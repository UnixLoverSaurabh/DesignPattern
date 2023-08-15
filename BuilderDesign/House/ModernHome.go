package House

type modernHome struct {
	Home
}

func (mh *modernHome) SetNumOfWindow() {
	mh.window = 4
}

func (mh *modernHome) SetNumOfGate() {
	mh.gate = 4
}

func newModernHome() *modernHome {
	return &modernHome{}
}

func (sh *modernHome) GetHome() *Home {
	return &Home{
		gate:   sh.gate,
		window: sh.window,
	}
}
