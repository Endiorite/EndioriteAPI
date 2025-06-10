package models

type UserLink struct {
	UserId   int    `json:"userId" binding:"required"`
	Username string `json:"username" binding:"required"`
	Code     int    `json:"code" binding:"required"`
}
