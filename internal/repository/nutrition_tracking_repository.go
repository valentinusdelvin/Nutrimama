package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type NutritionTrackingRepository struct {
	Repository[entity.NutritionTracking]
}

func NewNutritionTrackingRepository() *NutritionTrackingRepository {
	return &NutritionTrackingRepository{
		Repository: Repository[entity.NutritionTracking]{},
	}
}

func (r *NutritionTrackingRepository) FindByMotherIdDate(db *gorm.DB, motherId int, date string) (*entity.NutritionTracking, error) {
	var record entity.NutritionTracking
	err := db.Where("mother_id = ? AND DATE(created_at) = ?", motherId, date).First(&record).Error
	return &record, err
}
