package model

type Patient struct {
	Name                                                           string
	RegistrationDone, DoctorCheckUpDone, MedicineDone, PaymentDone bool
}

func NewPatient(name string) *Patient {
	return &Patient{Name: name}
}
