package matchDetails

import (
	"DesignPattern/CricketPortal/models/enums"
	"DesignPattern/CricketPortal/users"
)

type Wicket struct {
	WicketType enums.WicketType
	PlayerOut  users.Player
	CatchBy    users.Player
	RunOutBy   users.Player
	StumpedBy  users.Player
}
