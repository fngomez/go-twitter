package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(_user, _text string) *Tweet {
	return &Tweet{_user, _text, nil};
}