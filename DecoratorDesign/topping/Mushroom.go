package topping

import "DesignPattern/DecoratorDesign/pizza"

type Mushroom struct {
	pizza pizza.Pizza
}

func NewMushroom(pizza pizza.Pizza) *Mushroom {
	return &Mushroom{pizza: pizza}
}

func (m *Mushroom) GetPrice() int {
	return m.pizza.GetPrice() + 3
}
