package concreteHandler

import (
	"DesignPattern/ChainOfResponsibilityDesign/handler"
	"DesignPattern/ChainOfResponsibilityDesign/model"
	"fmt"
)

type Doctor struct {
	next handler.Department
}

func NewDoctor() *Doctor {
	return &Doctor{}
}

func (d *Doctor) Execute(patient *model.Patient) {
	if patient.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
	} else {
		fmt.Println("Patient is visiting the doctor")
	}
	d.next.Execute(patient)
}

func (d *Doctor) SetNext(department handler.Department) {
	d.next = department
}
