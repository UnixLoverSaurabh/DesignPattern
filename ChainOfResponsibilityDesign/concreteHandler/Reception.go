package concreteHandler

import (
	"DesignPattern/ChainOfResponsibilityDesign/handler"
	"DesignPattern/ChainOfResponsibilityDesign/model"
	"fmt"
)

type Reception struct {
	next handler.Department
}

func NewReception() *Reception {
	return &Reception{}
}

func (r *Reception) Execute(patient *model.Patient) {
	if patient.RegistrationDone {
		fmt.Println("Registration is already done")
	} else {
		fmt.Println("Executing registration for patient")
	}
	r.next.Execute(patient)
}

func (r *Reception) SetNext(department handler.Department) {
	r.next = department
}
