package computer

import (
	"DesignPattern/BridgeDesign/example2/printer"
)

type Windows struct {
	printer printer.Printer
}

func NewWindows() *Windows {
	return &Windows{}
}

func (w *Windows) Print() {
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(printer printer.Printer) {
	w.printer = printer
}
