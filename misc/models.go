package misc

import (
	"time"
)

//ChannelData ...
type ChannelData struct {
	ID         string
	Ch         chan string
	IsActive   *bool
	Status     *string
	StepTime   int
	Start      int
	Counter    *int
	CreatedAt  time.Time
	ModifiedAt time.Time
}

//CheckData ...
type CheckData struct {
	ChannelID      string
	StepTime       int
	CurrentCounter int
	Status         string
	CreatedAt      time.Time
}
