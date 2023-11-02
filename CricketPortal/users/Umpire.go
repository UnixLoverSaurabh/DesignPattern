package users

import (
	"DesignPattern/CricketPortal/matchDetails"
	"DesignPattern/CricketPortal/models/dataTypes"
)

type Umpire struct {
	dataTypes.Person
}

func (u *Umpire) assignMatch(match matchDetails.Match) bool {
	// todo
	return false
}
