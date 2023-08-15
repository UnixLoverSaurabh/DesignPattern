package commands

import (
	"DesignPattern/ParkingLot/service"
	"fmt"
)

type Exit struct {
}

func (e *Exit) Validate(command *Command) bool {
	// exit

	if command == nil || len(command.Parameter) != 0 {
		return false
	}
	return true
}

func (e *Exit) Execute(command *Command, parkingLotService *service.ParkingLotService) {
	fmt.Println("execute method of Exit")
}
