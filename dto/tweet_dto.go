package dto

type TweetDto struct {
	User string `json:"user" binding:"required"`
	Text string `json:"text" binding:"required"`
}