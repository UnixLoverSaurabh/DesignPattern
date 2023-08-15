package commands

import "DesignPattern/ParkingLot/service"

type CommandExecutor interface {
	Validate(command *Command) bool
	Execute(command *Command, parkingLotService *service.ParkingLotService)
}
