package main

import (
	"DesignPattern/ParkingLot/commands"
	"DesignPattern/ParkingLot/service"
	"bufio"
	"fmt"
	"os"
)

func main() {

	parkingLotService := service.NewParkingLotService()

	scanner := bufio.NewScanner(os.Stdin)
	var exitStatus bool
	for !exitStatus && scanner.Scan() {
		command := commands.NewCommand(scanner.Text())
		if command.Name == "exit" {
			exitStatus = true
		}
		commandExecFactoryObj := commands.NewCommandExecutorFactory().GetCommandExecutor(command.Name)
		if commandExecFactoryObj != nil && commandExecFactoryObj.Validate(command) {
			commandExecFactoryObj.Execute(command, parkingLotService)
		} else {
			fmt.Println("Invalid command")
		}

	}

}
