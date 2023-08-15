package commands

import (
	"DesignPattern/ParkingLot/service"
	"strconv"
)

type LeaveCommandExecutor struct {
}

func (l *LeaveCommandExecutor) Validate(command *Command) bool {
	// leave 4

	if command == nil || len(command.Parameter) != 1 {
		return false
	}
	if _, err := strconv.Atoi(command.Parameter[0]); err != nil {
		return false
	}
	return true
}

func (l *LeaveCommandExecutor) Execute(command *Command, parkingLotService *service.ParkingLotService) {
	slotNumber, _ := strconv.Atoi(command.GetParameter()[0])
	parkingLotService.MakeSlotFree(slotNumber)
}
