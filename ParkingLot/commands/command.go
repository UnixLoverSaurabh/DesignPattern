package commands

import "strings"

const (
	createParkingLot                     = "create_parking_lot"
	park                                 = "park"
	leave                                = "leave"
	status                               = "status"
	registrationNumbersForCarsWithColour = "registration_numbers_for_cars_with_colour"
	slotNumbersForCarsWithColour         = "slot_numbers_for_cars_with_colour"
	slotNumberForRegistrationNumber      = "slot_number_for_registration_number"
	exit                                 = "exit"
)

type Command struct {
	Name      string
	Parameter []string
}

func NewCommand(command string) *Command {
	args := strings.Split(command, " ")
	return &Command{Name: args[0], Parameter: args[1:]}
}

func (c *Command) GetParameter() []string {
	return c.Parameter
}
