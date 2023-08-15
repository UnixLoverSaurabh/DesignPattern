package commands

import (
	"DesignPattern/ParkingLot/model"
	"DesignPattern/ParkingLot/model/strategy"
	"DesignPattern/ParkingLot/service"
	"fmt"
	"strconv"
)

type CreateParkingLotCommandExecutor struct {
}

func (c *CreateParkingLotCommandExecutor) Validate(command *Command) bool {
	// create_parking_lot 6

	if command == nil || len(command.Parameter) != 1 {
		return false
	}
	if _, err := strconv.Atoi(command.Parameter[0]); err != nil {
		return false
	}
	return true
}

func (c *CreateParkingLotCommandExecutor) Execute(command *Command, parkingLotService *service.ParkingLotService) {

	lotCapacity, _ := strconv.Atoi(command.GetParameter()[0])

	fmt.Printf("Created a parking lot with %d slots\n", lotCapacity)
	parkingLotService.CreateParkingLot(model.NewParkingLot(lotCapacity), strategy.NewNaturalOrderingParkingStrategy())

}
