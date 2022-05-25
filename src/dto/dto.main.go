package dto

import "time"

type State struct {
	Message string
	Auth    bool
	Token   string
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID             string `json:"id"`
	CreatedAt      time.Time
	Username       string `json:"username"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
	Email          string `json:"email"`
	Post           []Post
}

type Post struct {
	UserID    string `json:"id"`
	CreatedAt time.Time
	Content   string `json:"content"`
	Like      string
	PostID    string
}

type Response struct {
	User
	State
}
