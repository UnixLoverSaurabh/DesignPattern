package concreteHandler

import (
	"DesignPattern/ChainOfResponsibilityDesign/handler"
	"DesignPattern/ChainOfResponsibilityDesign/model"
	"fmt"
)

type Cashier struct {
	next handler.Department
}

func NewCashier() *Cashier {
	return &Cashier{}
}

func (c *Cashier) Execute(patient *model.Patient) {
	if patient.PaymentDone {
		fmt.Println("Payment was done by patient")
	} else {
		fmt.Println("Patient is initializing payment")
	}
	// No need to execute next handler as payment is the last step in the process
}

func (c *Cashier) SetNext(department handler.Department) {
	c.next.SetNext(department)
}
