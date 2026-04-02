package entity

type UserBookmark struct {
	BookmarkID int       `gorm:"column:bookmark_id;primaryKey;autoIncrement"`
	UserID     int       `gorm:"column:user_id"`
	EduToolsID int       `gorm:"column:edu_tools_id"`
	EduTools   *EduTools `gorm:"foreignKey:EduToolsID;references:ID"`
}

func (UserBookmark) TableName() string {
	return "user_bookmarks"
}
