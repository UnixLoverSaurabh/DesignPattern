package matchDetails

import (
	"DesignPattern/CricketPortal/users"
	"time"
)

type Commentary struct {
	Comment     string
	CreatedAt   time.Time
	CommentedBy users.Commentator
}
