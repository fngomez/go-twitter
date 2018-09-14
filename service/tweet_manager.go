package service

import (
	"github.com/fngomez/go-twitter/domain"
	"time"
)

var _tweet *domain.Tweet;

func PublishTweet(tweet *domain.Tweet) {
	_tweet = tweet
	aux := time.Now()
	_tweet.Date = &aux
}

func GetTweet() *domain.Tweet {
	return _tweet;
}
