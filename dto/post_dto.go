package dto

type Post struct {
	UserId int `json:"userId" binding:"required"`
	Id int `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Body string `json:"body" binding:"required"`
}
