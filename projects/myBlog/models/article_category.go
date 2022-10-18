package models

import "github.com/jinzhu/gorm"

// 文章种类
type ArtcileCategory struct {
	gorm.Model
	CategoryID   uint   `gorm:"not null"`
	ArticleID    uint   `gorm:"not null"`
	CategoryName string `gorm:"type:varchar(20)"`
}
