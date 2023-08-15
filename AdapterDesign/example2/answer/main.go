package main

import (
	"DesignPattern/AdapterDesign/example2/answer/client"
	"DesignPattern/AdapterDesign/example2/answer/model"
)

func main() {
	client.NewClient().InsertSquareUsbInComputer(model.NewMac())
	client.NewClient().InsertSquareUsbInComputer(model.NewWindowsAdapter())
}
