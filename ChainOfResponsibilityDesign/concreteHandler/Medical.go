package concreteHandler

import (
	"DesignPattern/ChainOfResponsibilityDesign/handler"
	"DesignPattern/ChainOfResponsibilityDesign/model"
	"fmt"
)

type Medical struct {
	next handler.Department
}

func NewMedical() *Medical {
	return &Medical{}
}

func (m *Medical) Execute(patient *model.Patient) {
	if patient.MedicineDone {
		fmt.Println("Patient medication already done")
	} else {
		fmt.Println("Patient is taking medicine")
	}
	m.next.Execute(patient)
}

func (m *Medical) SetNext(department handler.Department) {
	m.next = department
}
