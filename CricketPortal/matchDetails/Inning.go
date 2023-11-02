package matchDetails

import "time"

type Inning struct {
	Number    int
	StartTime time.Time
	Overs     []*Over
}

func (inning *Inning) addOver(over Over) bool {
	// todo
	return false
}
