package main

import (
	"encoding/json"
	"fmt"
	"github.com/fngomez/go-twitter/domain"
	"github.com/fngomez/go-twitter/dto"
	"github.com/fngomez/go-twitter/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var manager = service.NewTweetManager()

func main() {
	// router := gin.Default() // Redirige todos los request al localhost:8080
	router := gin.New()

	router.POST("/tweet/add", addTweet)
	router.GET("/tweet/user/", getTweets)
	router.GET("external/consume/", getExternalData)
	router.POST("tweet/users/", getTweetsOfUsers)

	router.Run(":8081")
}

func addTweet(context *gin.Context) {

	var tweetDto dto.TweetDto
	context.BindJSON(&tweetDto)

	tweet := domain.NewTextTweet(tweetDto.User, tweetDto.Text)

	manager.PublishTweet(tweet)
}

func getTweets(context *gin.Context) {
	queryParam := context.Request.URL.Query()

	user := queryParam["user"][0]

	tweets := manager.GetTweetsByUser(user)

	fmt.Println(tweets)

	context.JSON(200, tweets)
}

func getTweetsOfUsers(context *gin.Context) {
	var users []dto.UserDto
	context.BindJSON(&users)

	fmt.Println(users)

	res := make(map[string][]domain.Tweet)

	for _, user := range users {
		res[user.Name] = manager.GetTweetsByUser(user.Name)
	}

	context.JSON(200, res)
}

func getExternalData(context *gin.Context) {

	response, error := http.Get("https://jsonplaceholder.typicode.com/posts")

	if error != nil {
		context.JSON(501, error.Error())
		return
	}

	defer response.Body.Close()

	var content []dto.Post
	json.NewDecoder(response.Body).Decode(&content)

	fmt.Println(content)

	context.JSON(response.StatusCode, content)
}