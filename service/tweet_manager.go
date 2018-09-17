package service

import (
	"errors"
	"github.com/fngomez/go-twitter/domain"
	"time"
)

type TweetManager struct {
	Tweets []*domain.Tweet
	TweetsByUser map[string][]*domain.Tweet
}

func NewTweetManager() TweetManager {
	var res = TweetManager{}
	res.initializeService()
	return res
}

func (tweetManager *TweetManager) initializeService() {
	tweetManager.Tweets = make([]*domain.Tweet, 0)
	tweetManager.TweetsByUser = make(map[string][]*domain.Tweet)
}

func (tweetManager *TweetManager) PublishTweet(tweet *domain.Tweet) (id int, error error) {

	if tweet.User == "" {
		return -1, errors.New("user is required")
	}

	if len(tweet.Text) == 0 {
		return -1, errors.New("text is required")
	}

	if len(tweet.Text) > 140 {
		return -1, errors.New("text exceeds 140 characters")
	}

	tweetManager.addUserTweet(tweet, tweet.User)

	return tweet.Id, nil
}

func (tweetManager *TweetManager) addUserTweet(tweet *domain.Tweet, user string) {

	tweetManager.Tweets = append(tweetManager.Tweets, tweet)
	aux := time.Now()
	tweet.Date = &aux
	tweet.Id = len(tweetManager.Tweets) -1

	tweetManager.TweetsByUser[user] = append(tweetManager.TweetsByUser[user], tweet)

}

func (tweetManager *TweetManager) GetTweets() []*domain.Tweet {
	return tweetManager.Tweets
}

func (tweetManager *TweetManager) GetTweet() *domain.Tweet {
	return tweetManager.Tweets[len(tweetManager.Tweets) - 1]
}

func (tweetManager *TweetManager) GetTweetById(id int) *domain.Tweet {
	return tweetManager.Tweets[id]
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {
	var tweets = 0

	for _, tweet := range tweetManager.Tweets {
		if tweet.User == user {
			tweets++
		}
	}

	return tweets
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []*domain.Tweet{
	return tweetManager.TweetsByUser[user]
}

func (tweetManager *TweetManager) DeleteTweet (user string, idTweet int) {
	for  index ,tweet := range tweetManager.TweetsByUser[user] {
		if tweet.Id == idTweet {
			tweetManager.TweetsByUser[user] = deleteTweetFromSlice(tweetManager.TweetsByUser[user], index)
		}
	}

	//TODO: terminar
}

func deleteTweetFromSlice(arr []*domain.Tweet, index int) []*domain.Tweet{
	return append(arr[:index], arr[index+1:]...)
}