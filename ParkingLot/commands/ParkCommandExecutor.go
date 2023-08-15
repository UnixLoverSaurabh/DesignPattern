package commands

import (
	"DesignPattern/ParkingLot/model"
	"DesignPattern/ParkingLot/service"
)

type ParkCommandExecutor struct {
}

func (p *ParkCommandExecutor) Validate(command *Command) bool {
	// park KA-01-HH-1234 White

	if command == nil || len(command.Parameter) != 2 {
		return false
	}

	return true
}

func (p *ParkCommandExecutor) Execute(command *Command, parkingLotService *service.ParkingLotService) {
	car := model.NewCar(command.GetParameter()[0], command.GetParameter()[1])
	parkingLotService.Park(car)
}
