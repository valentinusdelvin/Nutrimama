package entity

import "time"

type EduTools struct {
	ID          int    `gorm:"column:edu_tools_id;primaryKey;autoIncrement"`
	Publisher   string `gorm:"column:publisher"`
	Title      string `gorm:"column:title"`
	Slug       string `gorm:"column:slug"`
	Category   string `gorm:"column:category"`
	Thumbnail  string `gorm:"column:thumbnail"`
	Content        string `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (EduTools) TableName() string {
	return "educational_tools"
}