package main

import (
	"DesignPattern/BridgeDesign/example2/computer"
	"DesignPattern/BridgeDesign/example2/printer"
)

func main() {
	windows := computer.NewWindows()
	windows.SetPrinter(printer.NewEpson())
	windows.Print()

	windows.SetPrinter(printer.NewHp())
	windows.Print()

	// ----------------------------------
	mac := computer.NewMac()
	mac.SetPrinter(printer.NewEpson())
	mac.Print()

	mac.SetPrinter(printer.NewHp())
	mac.Print()

}
