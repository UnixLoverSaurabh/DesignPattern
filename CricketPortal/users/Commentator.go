package users

import (
	"DesignPattern/CricketPortal/matchDetails"
	"DesignPattern/CricketPortal/models/dataTypes"
)

type Commentator struct {
	dataTypes.Person
}

func (c *Commentator) assignMatch(match matchDetails.Match) bool {
	// todo
	return false
}
