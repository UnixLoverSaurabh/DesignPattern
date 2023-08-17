package computer

import (
	"DesignPattern/BridgeDesign/example2/printer"
)

type Mac struct {
	printer printer.Printer
}

func NewMac() *Mac {
	return &Mac{}
}

func (w *Mac) Print() {
	w.printer.PrintFile()
}

func (w *Mac) SetPrinter(printer printer.Printer) {
	w.printer = printer
}
