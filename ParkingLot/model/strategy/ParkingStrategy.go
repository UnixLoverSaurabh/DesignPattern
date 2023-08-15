package strategy

type ParkingStrategy interface {
	GetNextSlot() int
	AddSlot(slotNumber int)
	RemoveSlot(slotNumber int)
}
