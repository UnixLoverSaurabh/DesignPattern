package model

type Slot struct {
	slotNumber int
	parkedCar  *Car
}

func NewSlot(slotNumber int) *Slot {
	return &Slot{slotNumber: slotNumber}
}

func (slot *Slot) isSlotFree() bool {
	return slot.parkedCar != nil
}

func (slot *Slot) assignCar(car *Car) {
	slot.parkedCar = car
}

func (slot *Slot) unAssignCar() {
	slot.parkedCar = nil
}
