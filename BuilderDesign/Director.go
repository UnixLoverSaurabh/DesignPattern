package main

import "DesignPattern/BuilderDesign/House"

type Director struct {
}

func NewDirector() *Director {
	return &Director{}
}

func (d *Director) buildHouse(builder House.IBuilder) *House.Home {
	builder.SetNumOfWindow()
	builder.SetNumOfGate()
	return builder.GetHome()
}
