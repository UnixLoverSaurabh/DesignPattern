package main

import (
	"DesignPattern/AdapterDesign/example2/question/client"
	"DesignPattern/AdapterDesign/example2/question/model"
)

func main() {
	client.NewClient().InsertSquareUsbInComputer(model.NewMac())
}
