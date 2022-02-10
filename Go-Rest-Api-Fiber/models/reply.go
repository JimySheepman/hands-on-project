package models

import "gorm.io/gorm"

// Book model
type Reply struct {
	gorm.Model
	ID     string `gorm:"primaryKey"`
	Author string `json:"author"`
	Body   string `json:"body"`
	PostID string `json:"postId"`
	Post   Post   `json:"post" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
