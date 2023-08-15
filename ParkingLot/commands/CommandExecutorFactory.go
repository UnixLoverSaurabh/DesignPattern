package commands

type CommandExecutorFactory struct {
	commands map[string]CommandExecutor
}

func NewCommandExecutorFactory() *CommandExecutorFactory {

	commands := make(map[string]CommandExecutor)
	commands[createParkingLot] = &CreateParkingLotCommandExecutor{}
	commands[park] = &ParkCommandExecutor{}
	commands[leave] = &LeaveCommandExecutor{}
	commands[status] = &StatusCommandExecutor{}
	commands[registrationNumbersForCarsWithColour] = &ColorToRegNumberCommandExecutor{}
	commands[slotNumbersForCarsWithColour] = &ColorToSlotNumberCommandExecutor{}
	commands[slotNumberForRegistrationNumber] = &SlotForRegNumberCommandExecutor{}
	commands[exit] = &Exit{}

	return &CommandExecutorFactory{commands}
}

func (cf *CommandExecutorFactory) GetCommandExecutor(command string) CommandExecutor {
	return cf.commands[command]
}
