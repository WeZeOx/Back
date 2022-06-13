package models

import "time"

type Comment struct {
	PostId         string    `json:"post_id"`
	UserId         string    `json:"user_id"`
	CommentId      string    `json:"comment_id"`
	CreatedAt      time.Time `json:"created_at_comment"`
	ContentComment string    `json:"content_comment"`
	Like           string    `json:"like"`
	User           User
}
