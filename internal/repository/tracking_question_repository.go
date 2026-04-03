package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type TrackingQuestionRepository struct {
	Repository[entity.Question]
}

func NewTrackingQuestionRepository() *TrackingQuestionRepository {
	return &TrackingQuestionRepository{
		Repository: Repository[entity.Question]{},
	}
}

// FindByCategory fetches all questions for a specific target group (e.g., 'ibu_hamil' or 'balita')
func (r *TrackingQuestionRepository) FindByCategory(db *gorm.DB, category string) ([]entity.Question, error) {
	var questions []entity.Question
	err := db.Preload("Options").Where("category = ?", category).Order("display_order ASC").Find(&questions).Error
	return questions, err
}
