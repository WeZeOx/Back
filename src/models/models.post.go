package models

import "time"

type Post struct {
	UserID    string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Like      string    `json:"like"`
	Dislike   string    `json:"dislike"`
	PostID    string    `json:"post_id" gorm:"primaryKey"`
	Category  string    `json:"category"`
	Comment   []Comment
}
