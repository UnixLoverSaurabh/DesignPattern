package printer

import "fmt"

type Hp struct {
	name string
}

func NewHp() *Hp {
	return &Hp{
		name: "Hp",
	}
}

func (h *Hp) PrintFile() {
	fmt.Println("Printing with", h.name)
}
