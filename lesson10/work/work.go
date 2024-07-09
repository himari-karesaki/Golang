package main

import "time"

// Post構造体を定義
type Post struct {
	ID        int       `json:"id"gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}
