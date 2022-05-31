package models

import "time"

type Post struct {
	UserID    string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	Like      string    `json:"like"`
	Dislike   string    `json:"dislike"`
	PostID    string    `json:"post_id"`
}
