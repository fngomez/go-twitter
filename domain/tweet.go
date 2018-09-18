package domain

import "time"

type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetId(int)
	SetDate(*time.Time)
	SetText(string)
	PrintableTweet() string
}