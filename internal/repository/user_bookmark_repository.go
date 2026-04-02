package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type UserBookmarkRepository struct {
	Repository[entity.UserBookmark]
}

func NewUserBookmarkRepository() *UserBookmarkRepository {
	return &UserBookmarkRepository{
		Repository: Repository[entity.UserBookmark]{},
	}
}

func (r *UserBookmarkRepository) FindByUserId(db *gorm.DB, userId int) ([]entity.UserBookmark, error) {
	var records []entity.UserBookmark
	// Preload EduTools to easily fetch the mapped properties (title, content, publisher) alongside the bookmark ID
	err := db.Preload("EduTools").Where("user_id = ?", userId).Order("bookmark_id desc").Find(&records).Error
	return records, err
}

func (r *UserBookmarkRepository) DeleteBookmark(db *gorm.DB, userId int, eduToolsId int) error {
	return db.Where("user_id = ? AND edu_tools_id = ?", userId, eduToolsId).Delete(&entity.UserBookmark{}).Error
}
