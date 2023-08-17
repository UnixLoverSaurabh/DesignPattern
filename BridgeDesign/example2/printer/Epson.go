package printer

import "fmt"

type Epson struct {
	name string
}

func NewEpson() *Epson {
	return &Epson{
		name: "Epson",
	}
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing with", e.name)
}
