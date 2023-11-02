package enums

type MatchResult int

const (
	LIVE MatchResult = iota
	CANCELED
	FINISHED
	DRAWN
	YET_TO_START
)
