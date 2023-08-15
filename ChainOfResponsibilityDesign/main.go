package main

import (
	"DesignPattern/ChainOfResponsibilityDesign/concreteHandler"
	"DesignPattern/ChainOfResponsibilityDesign/model"
)

func main() {

	// patient -> reception -> doctor -> medical -> cashier

	cashier := concreteHandler.NewCashier()

	medical := concreteHandler.NewMedical()
	medical.SetNext(cashier)

	doctor := concreteHandler.NewDoctor()
	doctor.SetNext(medical)

	reception := concreteHandler.NewReception()
	reception.SetNext(doctor)

	reception.Execute(model.NewPatient("xyz"))
}
