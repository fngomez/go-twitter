package service

import (
	"github.com/fngomez/go-twitter/domain"
	"time"
	"errors"
)

var _tweet *domain.Tweet;

func PublishTweet(tweet *domain.Tweet) error {

	if tweet.User == "" {
		return errors.New("user is required")
	}

	if len(tweet.Text) == 0 {
		return errors.New("text is required")
	}

	if len(tweet.Text) > 140 {
		return errors.New("text exceeds 140 characters")
	}

	_tweet = tweet
	aux := time.Now()
	_tweet.Date = &aux

	return nil
}

func GetTweet() *domain.Tweet {
	return _tweet;
}
