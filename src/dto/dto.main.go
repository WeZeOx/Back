package dto

import "time"

type State struct {
	Message string `json:"message"`
	Auth    bool   `json:"auth"`
	Token   string `json:"token"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	VerifyPassword string    `json:"verify_password"`
	Email          string    `json:"email"`
	Post           []Post
}

type ResponseUser struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}

type ResponsePostUser struct {
	Username              string    `json:"username"`
	PasswordButINotWantIt string    `json:"password"`
	CreatedAt             time.Time `json:"created_at"`
	Email                 string    `json:"email"`
	PostButINotWantIt     string    `json:"post"`
	Content               string    `json:"content"`
	Like                  string    `json:"like"`
	Dislike               string    `json:"dislike"`
	PostID                string    `json:"post_id"`
}

type Post struct {
	UserID    string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	Like      string    `json:"like"`
	Dislike   string    `json:"dislike"`
	PostID    string    `json:"post_id"`
}

type Response struct {
	User
	State
}
