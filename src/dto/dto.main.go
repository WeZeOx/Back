package dto

import (
	"Forum-Back-End/src/models"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type State struct {
	Message string `json:"message"`
	Auth    bool   `json:"auth"`
	Token   string `json:"token"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ContentCommentCreator struct {
	UserId         string    `json:"user_id"`
	PostId         string    `json:"post_id"`
	CreatedAt      time.Time `json:"created_at_comment"`
	ContentComment string    `json:"content_comment"`
	CommentId      string    `json:"comment_id"`
}

type Claims struct {
	ID      string `json:"id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
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
}

type ResponsePostUser struct {
	ID string `json:"id"`
	models.User
	models.Post
}

type PostUserResponseForFront struct {
	UserID       string    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	Username     string    `json:"username"`
	Content      string    `json:"content"`
	Like         string    `json:"like"`
	PostID       string    `json:"post_id"`
	Categories   string    `json:"categories"`
	Admin        bool      `json:"admin"`
	NumberOfPost int       `json:"number_of_post"`
}

type Post struct {
	UserID    string    `json:"id"`
	CreatedAt time.Time `json:"created_at_post"`
	Content   string    `json:"content"`
	Like      string    `json:"like"`
	Dislike   string    `json:"dislike"`
	Category  string    `json:"category"`
	PostID    string    `json:"post_id"`
}

type Response struct {
	User
	State
}

type PostWithCommentResponse struct {
	UserId         string    `json:"user_id"`
	ContentComment string    `json:"content_comment"`
	CreatedAt      time.Time `json:"created_at"`
	Content        string    `json:"content"`
	Like           string    `json:"like"`
	Category       string    `json:"category"`
	Username       string    `json:"username" gorm:"unique"`
}
