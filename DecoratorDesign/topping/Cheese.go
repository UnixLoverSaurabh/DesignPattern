package topping

import "DesignPattern/DecoratorDesign/pizza"

type Cheese struct {
	pizza pizza.Pizza
}

func NewCheese(pizza pizza.Pizza) *Cheese {
	return &Cheese{pizza: pizza}
}

func (c *Cheese) GetPrice() int {
	return c.pizza.GetPrice() + 2
}
