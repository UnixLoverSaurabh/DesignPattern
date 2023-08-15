package strategy

import (
	"github.com/emirpasic/gods/sets/treeset"
)

type NaturalOrderingParkingStrategy struct {
	slotTreeSet *treeset.Set
}

func NewNaturalOrderingParkingStrategy() *NaturalOrderingParkingStrategy {
	return &NaturalOrderingParkingStrategy{
		slotTreeSet: treeset.NewWithIntComparator(),
	}
}

func (n *NaturalOrderingParkingStrategy) GetNextSlot() int {
	availableSlots := n.slotTreeSet.Values()
	if len(availableSlots) == 0 {
		return 0
	}
	return availableSlots[0].(int)
}

func (n *NaturalOrderingParkingStrategy) AddSlot(slotNumber int) {
	n.slotTreeSet.Add(slotNumber)
}

func (n *NaturalOrderingParkingStrategy) RemoveSlot(slotNumber int) {
	n.slotTreeSet.Remove(slotNumber)
}
