package model

type Car struct {
	RegistrationNo, Color string
}

func NewCar(registrationNo, color string) *Car {
	return &Car{RegistrationNo: registrationNo, Color: color}
}
