package models

type User struct {
	Id int `json:"id"`
}

type UserResponse struct {
	Status string `json:"status"`
	Id     int    `json:"id"`
}