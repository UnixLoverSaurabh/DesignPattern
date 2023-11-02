package matchDetails

import (
	"DesignPattern/CricketPortal/models/enums"
	"DesignPattern/CricketPortal/teamDetails"
	"DesignPattern/CricketPortal/users"
	"time"
)

type Match struct {
	MatchFormat  enums.MatchFormat
	Number       int
	StartTime    time.Time
	Innings      []*Inning
	Commentators []*users.Commentator
	Umpires      []*users.Umpire
	Playing11    []*teamDetails.Playing11
	MatchResult  enums.MatchResult
}
