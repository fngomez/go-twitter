package service

import (
	"github.com/fngomez/go-twitter/domain"
	"time"
	"errors"
)

var _tweets []*domain.Tweet;

func InitializeService() {
	_tweets = make([]*domain.Tweet, 0)
}

func isAuth(user domain.User) bool {
	return true
}

func PublishTweet(tweet *domain.Tweet) (id int, error error) {

	if tweet.User == "" {
		return -1, errors.New("user is required")
	}

	if len(tweet.Text) == 0 {
		return -1, errors.New("text is required")
	}

	if len(tweet.Text) > 140 {
		return -1, errors.New("text exceeds 140 characters")
	}

	_tweets = append(_tweets, tweet)
	aux := time.Now()
	tweet.Date = &aux
	tweet.Id = len(_tweets) -1

	return tweet.Id, nil
}

func GetTweets() []*domain.Tweet {
	return _tweets
}

func GetTweet() *domain.Tweet {
	return _tweets[len(_tweets) - 1]
}

func GetTweetById(id int) *domain.Tweet {
	return _tweets[id]
}

func CountTweetsByUser(user string) int {
	var tweets = 0

	for _, tweet := range _tweets {
		if tweet.User == user {
			tweets++
		}
	}

	return tweets
}
