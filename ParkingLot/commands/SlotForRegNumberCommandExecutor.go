package commands

import (
	"DesignPattern/ParkingLot/service"
	"fmt"
)

type SlotForRegNumberCommandExecutor struct {
}

func (s *SlotForRegNumberCommandExecutor) Validate(command *Command) bool {
	// slot_number_for_registration_number KA-01-HH-3141

	if command == nil || len(command.Parameter) != 1 {
		return false
	}
	return true
}

func (s *SlotForRegNumberCommandExecutor) Execute(command *Command, parkingLotService *service.ParkingLotService) {
	var requiredSlot int
	for slotNumber, car := range parkingLotService.GetOccupiedSlots() {
		if car.RegistrationNo == command.GetParameter()[0] {
			requiredSlot = slotNumber
		}
	}
	if requiredSlot > 0 {
		fmt.Println(requiredSlot)
	} else {
		fmt.Println("Not found")
	}
}
