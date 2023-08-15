package commands

import (
	"DesignPattern/ParkingLot/service"
	"fmt"
)

type ColorToRegNumberCommandExecutor struct {
}

func (c *ColorToRegNumberCommandExecutor) Validate(command *Command) bool {
	// registration_numbers_for_cars_with_colour White

	if command == nil || len(command.Parameter) != 1 {
		return false
	}
	return true
}

func (c *ColorToRegNumberCommandExecutor) Execute(command *Command, parkingLotService *service.ParkingLotService) {
	for _, car := range parkingLotService.GetOccupiedSlots() {
		if car.Color == command.GetParameter()[0] {
			fmt.Printf("%s\t", car.RegistrationNo)
		}
	}
}
