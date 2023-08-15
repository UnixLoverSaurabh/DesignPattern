package commands

import (
	"DesignPattern/ParkingLot/service"
	"fmt"
)

type StatusCommandExecutor struct {
}

func (s *StatusCommandExecutor) Validate(command *Command) bool {
	// status

	if command == nil || len(command.Parameter) != 0 {
		return false
	}
	return true
}

func (s *StatusCommandExecutor) Execute(command *Command, parkingLotService *service.ParkingLotService) {

	fmt.Println("Slot No.\t Registration No\t Colour")
	for slotNumber, car := range parkingLotService.GetOccupiedSlots() {
		fmt.Printf("%d\t %s\t %s\n", slotNumber, car.RegistrationNo, car.Color)
	}
}
