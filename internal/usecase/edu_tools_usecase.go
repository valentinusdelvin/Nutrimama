package usecase

import (
	"errors"
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
	"nutrimama/internal/model/converter"
	"nutrimama/internal/repository"

	"gorm.io/gorm"
)

type EduToolsUseCase struct {
	DB                 *gorm.DB
	EduToolsRepository *repository.EduToolsRepository
}

func NewEduToolsUseCase(db *gorm.DB, eduToolsRepo *repository.EduToolsRepository) *EduToolsUseCase {
	return &EduToolsUseCase{
		DB:                 db,
		EduToolsRepository: eduToolsRepo,
	}
}

func (u *EduToolsUseCase) Create(req model.CreateEduToolsRequest) (*model.EduToolsResponse, error) {
	c := u.DB.Begin()

	eduTools := entity.EduTools{
		Publisher: req.Publisher,
		Slug:      req.Slug,
		Title:     req.Title,
		Category:  req.Category,
		Thumbnail: req.Thumbnail,
		Content:   req.Content,
	}

	if err := u.EduToolsRepository.Create(c, &eduTools); err != nil {
		c.Rollback()
		return nil, err
	}

	if err := c.Commit().Error; err != nil {
		return nil, err
	}

	return converter.EduToolsToResponse(&eduTools), nil
}

func (u *EduToolsUseCase) Update(req model.UpdateEduToolsRequest) (*model.EduToolsResponse, error) {
	c := u.DB.Begin()

	eduTools, err := u.EduToolsRepository.Get(c, req.ID)
	if err != nil {
		c.Rollback()
		return nil, errors.New("edu tools not found")
	}

	if req.Publisher != "" {
		eduTools.Publisher = req.Publisher
	}
	if req.Title != "" {
		eduTools.Title = req.Title
	}
	if req.Category != "" {
		eduTools.Category = req.Category
	}
	if req.Thumbnail != "" {
		eduTools.Thumbnail = req.Thumbnail
	}
	if req.Content != "" {
		eduTools.Content = req.Content
	}

	if err := u.EduToolsRepository.Update(c, eduTools); err != nil {
		c.Rollback()
		return nil, err
	}

	if err := c.Commit().Error; err != nil {
		return nil, err
	}

	return converter.EduToolsToResponse(eduTools), nil
}

func (u *EduToolsUseCase) Get(slug string) (*model.EduToolsResponse, error) {
	eduTools, err := u.EduToolsRepository.GetBySlug(u.DB, slug)
	if err != nil {
		return nil, errors.New("edu tools not found organically by slug")
	}

	return converter.EduToolsToResponse(eduTools), nil
}

func (u *EduToolsUseCase) List() ([]model.EduToolsResponse, error) {
	eduToolsList, err := u.EduToolsRepository.List(u.DB)
	if err != nil {
		return nil, err
	}

	return converter.EduToolsListToResponse(eduToolsList), nil
}

func (u *EduToolsUseCase) Delete(id int) error {
	c := u.DB.Begin()

	if err := u.EduToolsRepository.Delete(c, id); err != nil {
		c.Rollback()
		return err
	}

	return c.Commit().Error
}
