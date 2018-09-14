package service

var _tweet string;

func PublishTweet(tweet string) {
	_tweet = tweet;
}

func GetTweet() string {
	return _tweet;
}
