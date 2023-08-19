package pizza

type VeggieMania struct {
}

func (v *VeggieMania) GetPrice() int {
	return 10
}

func NewVeggieMania() *VeggieMania {
	return &VeggieMania{}
}
