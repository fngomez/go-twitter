package domain

import "time"

type TextTweet struct {
	Id int
	User string
	Text string
	Date *time.Time
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) SetDate(date *time.Time) {
	tweet.Date = date
}

func NewTextTweet(_user, _text string) *TextTweet {
	return &TextTweet{-1,_user, _text, nil};
}

func (tweet *TextTweet) PrintableTweet() string {
	return "@" + tweet.User + ": " + tweet.Text
}

func (tweet *TextTweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text
}


