package model

import "fmt"

type Mac struct {
}

func NewMac() *Mac {
	return &Mac{}
}

func (m *Mac) InsertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}
