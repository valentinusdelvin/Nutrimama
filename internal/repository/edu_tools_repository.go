package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type EduToolsRepository struct {
	Repository[entity.EduTools]
}

func NewEduToolsRepository() *EduToolsRepository {
	return &EduToolsRepository{
		Repository: Repository[entity.EduTools]{},
	}
}

func (r *EduToolsRepository) GetBySlug(db *gorm.DB, slug string) (*entity.EduTools, error) {
	var eduTools entity.EduTools
	if err := db.Where("slug = ?", slug).First(&eduTools).Error; err != nil {
		return nil, err
	}
	return &eduTools, nil
}
