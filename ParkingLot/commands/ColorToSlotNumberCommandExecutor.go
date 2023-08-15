package commands

import (
	"DesignPattern/ParkingLot/service"
	"fmt"
)

type ColorToSlotNumberCommandExecutor struct {
}

func (c *ColorToSlotNumberCommandExecutor) Validate(command *Command) bool {
	// slot_numbers_for_cars_with_colour White

	if command == nil || len(command.Parameter) != 1 {
		return false
	}
	return true
}

func (c *ColorToSlotNumberCommandExecutor) Execute(command *Command, parkingLotService *service.ParkingLotService) {
	for slotNumber, car := range parkingLotService.GetOccupiedSlots() {
		if car.Color == command.GetParameter()[0] {
			fmt.Printf("%d\t", slotNumber)
		}
	}
}
