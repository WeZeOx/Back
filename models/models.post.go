package models

import "time"

type Post struct {
	UserID    string `json:"id"`
	CreatedAt time.Time
	Content   string `json:"content"`
	Like      string
	PostID    string
}
