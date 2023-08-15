package service

import (
	"DesignPattern/ParkingLot/model"
	"DesignPattern/ParkingLot/model/strategy"
	"fmt"
)

type ParkingLotService struct {
	parkingLot      *model.ParkingLot
	parkingStrategy strategy.ParkingStrategy
}

func NewParkingLotService() *ParkingLotService {
	return &ParkingLotService{}
}

func (pls *ParkingLotService) CreateParkingLot(parkingLot *model.ParkingLot, parkingStrategy strategy.ParkingStrategy) {
	pls.parkingLot = parkingLot
	pls.parkingStrategy = parkingStrategy

	for i := 1; i <= pls.parkingLot.Capacity; i++ {
		pls.parkingStrategy.AddSlot(i)
	}
}

func (pls *ParkingLotService) Park(car *model.Car) {
	nextAvailableSlot := pls.parkingStrategy.GetNextSlot()
	if nextAvailableSlot <= 0 {
		fmt.Println("Sorry, parking lot is full")
		return
	}

	fmt.Println("Allocated slot number: ", nextAvailableSlot)
	pls.parkingStrategy.RemoveSlot(nextAvailableSlot)

	pls.parkingLot.Slots[nextAvailableSlot] = car
}

func (pls *ParkingLotService) MakeSlotFree(slotNumber int) {

	pls.parkingStrategy.AddSlot(slotNumber)

	delete(pls.parkingLot.Slots, slotNumber)

	fmt.Printf("Slot number %d is free", slotNumber)
}

func (pls *ParkingLotService) GetOccupiedSlots() map[int]*model.Car {
	return pls.parkingLot.Slots
}
