package main

import (
	"DesignPattern/DecoratorDesign/pizza"
	"DesignPattern/DecoratorDesign/topping"
	"fmt"
)

func main() {
	vm := pizza.NewVeggieMania()
	cheese := topping.NewCheese(vm)
	cheeseAndMushroom := topping.NewMushroom(cheese)
	fmt.Println(cheeseAndMushroom.GetPrice())
}
