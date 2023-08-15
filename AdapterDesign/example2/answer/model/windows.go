package model

import "fmt"

type Windows struct {
}

func NewWindows() *Windows {
	return &Windows{}
}

func (w *Windows) InsertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}
