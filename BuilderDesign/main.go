package main

import (
	"DesignPattern/BuilderDesign/House"
	"fmt"
)

func main() {
	builder := House.GetBuilder("simple")
	house := NewDirector().buildHouse(builder)
	fmt.Println(house)

	builder = House.GetBuilder("modern")
	house = NewDirector().buildHouse(builder)
	fmt.Println(house)
}
