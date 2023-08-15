package handler

import "DesignPattern/ChainOfResponsibilityDesign/model"

type Department interface {
	Execute(patient *model.Patient)
	SetNext(department Department)
}
