package models

import "time"

type Post struct {
	UserID    string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	Like      string    `json:"like"`
	Category  string    `json:"category"`
	PostID    string    `json:"post_id"`
}
