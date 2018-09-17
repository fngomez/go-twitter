package domain

import "time"

type Tweet struct {
	Id int
	User string
	Text string
	Date *time.Time
}

func NewTweet(_user, _text string) *Tweet {
	return &Tweet{-1,_user, _text, nil};
}

func (tweet *Tweet) PrintableTweet() string {
	return  "@" + tweet.User + ": " + tweet.Text
}