package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CodeSave struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Name      string    `json:"user_name" gorm:"type:varchar(20);not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Code      string    `json:"code" gorm:"type:longtext;not null"`
	Title     string    `json:"title" gorm:"type:varchar(50);not null"`
	CreatedAt Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time      `json:"updated_at" gorm:"type:timestamp"`
}

//BeforeCreate b
func (code *CodeSave) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
