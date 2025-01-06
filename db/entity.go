package db

import "time"

type Book struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title,omitempty"`
	Author      string    `json:"author,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
