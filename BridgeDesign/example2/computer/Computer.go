package computer

import "DesignPattern/BridgeDesign/example2/printer"

type Computer interface {
	Print()
	SetPrinter(printer.Printer)
}
