package dto

import (
	"Forum-Back-End/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type ResponseState struct {
	Message string `json:"message"`
	Auth    bool   `json:"auth"`
	Token   string `json:"token"`
}

type BodyLoginRequest struct {
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

type JwtClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
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

type ResponseWithSafeField struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

type ResponsePostUser struct {
	ID string `json:"id"`
	models.User
	models.Post
}

type PostModel struct {
	UserID          string    `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
	Username        string    `json:"username"`
	Content         string    `json:"content"`
	Like            string    `json:"like"`
	PostID          string    `json:"post_id"`
	Categories      string    `json:"categories"`
	Admin           bool      `json:"admin"`
	NumberOfComment int       `json:"number_of_comment"`
}

type CommentsWithPost struct {
	Comments []fiber.Map
	Post     PostModel
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

type PostWithCommentResponse struct {
	UserId         string    `json:"user_id"`
	ContentComment string    `json:"content_comment"`
	CreatedAt      time.Time `json:"created_at"`
	Username       string    `json:"username"`
	Like           string    `json:"like"`
	CommentId      string    `json:"comment_id"`
}

type ResponseComment struct {
	UserId         string    `json:"user_id"`
	ContentComment string    `json:"content_comment"`
	CreatedAt      time.Time `json:"created_at"`
	Username       string    `json:"username"`
	Like           string    `json:"like"`
	CommentId      string    `json:"comment_id"`
}
