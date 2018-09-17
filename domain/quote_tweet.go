package domain

type QuoteTweet struct {
	TextTweet
	quotedTweet Tweet
}

func NewQuoteTweet(_user, _text string, tweet Tweet) *QuoteTweet {
	return &QuoteTweet{TextTweet{-1, _user, _text, nil }, tweet }
}

func (tweet *QuoteTweet) PrintableTweet() string {
	return "@" + tweet.User + ": " + tweet.Text + " \"" + tweet.quotedTweet.PrintableTweet() + "\""
}
