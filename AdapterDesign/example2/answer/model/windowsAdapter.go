package model

type WindowsAdapter struct {
	windows *Windows
}

func NewWindowsAdapter() *WindowsAdapter {
	return &WindowsAdapter{}
}

func (w *WindowsAdapter) InsertInSquarePort() {
	w.windows.InsertInCirclePort()
}
