package model

import (
	uuid "github.com/satori/go.uuid"
)

//Post p
type Post struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Title     string    `json:"title" gorm:"type:varchar(50);not null"`
	Content   string    `json:"content" gorm:"tyep:text;not null"`
	CreatedAt Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time      `json:"updated_at" gorm:"type:timestamp"`
}
