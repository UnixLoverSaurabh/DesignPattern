package matchDetails

import (
	"DesignPattern/CricketPortal/models/enums"
	"DesignPattern/CricketPortal/users"
)

type Ball struct {
	BalledBy   users.Player
	PlayedBy   users.Player
	BallType   enums.BallType
	Runs       []*Run
	Wicket     *Wicket
	Commentary *Commentary
}
