package domain

type ImageTweet struct {
	TextTweet
	ImageUrl string
}

func NewImageTweet(_user, _text, imageUrl string) *ImageTweet {
	return &ImageTweet{TextTweet{-1,_user, _text, nil}, imageUrl };
}

func (tweet *ImageTweet) PrintableTweet() string {
	return "@" + tweet.User + ": " + tweet.Text + " " + tweet.ImageUrl
}