package model

type ParkingLot struct {
	Capacity int
	Slots    map[int]*Car
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Capacity: capacity,
		Slots:    make(map[int]*Car),
	}
}
