package usecase

import (
	"nutrimama/internal/entity"
	"nutrimama/internal/repository"

	"gorm.io/gorm"
)

type BookmarkUseCase struct {
	DB               *gorm.DB
	UserBookmarkRepo *repository.UserBookmarkRepository
}

func NewBookmarkUseCase(db *gorm.DB, ubRepo *repository.UserBookmarkRepository) *BookmarkUseCase {
	return &BookmarkUseCase{
		DB:               db,
		UserBookmarkRepo: ubRepo,
	}
}

func (u *BookmarkUseCase) ToggleBookmark(userId int, eduToolsId int) (string, error) {
	c := u.DB.Begin()

	var existing entity.UserBookmark
	err := c.Where("user_id = ? AND edu_tools_id = ?", userId, eduToolsId).First(&existing).Error
	
	if err == nil {
		// ALREADY EXISTS -> DELETE IT
		if err := u.UserBookmarkRepo.DeleteBookmark(c, userId, eduToolsId); err != nil {
			c.Rollback()
			return "", err
		}
		c.Commit()
		return "deleted", nil
	}

	// DOES NOT EXIST -> CREATE IT
	bookmark := entity.UserBookmark{
		UserID:     userId,
		EduToolsID: eduToolsId,
	}
	if err := u.UserBookmarkRepo.Create(c, &bookmark); err != nil {
		c.Rollback()
		return "", err
	}

	c.Commit()
	return "created", nil
}

func (u *BookmarkUseCase) GetUserBookmarks(userId int) ([]entity.UserBookmark, error) {
	return u.UserBookmarkRepo.FindByUserId(u.DB, userId)
}
