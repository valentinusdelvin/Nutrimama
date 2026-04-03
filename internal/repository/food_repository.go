package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type FoodRepository struct {
	Repository[entity.Food]
}

func NewFoodRepository() *FoodRepository {
	return &FoodRepository{
		Repository: Repository[entity.Food]{},
	}
}

// GetAllFoods efficiently grabs all system references natively
func (r *FoodRepository) GetAllFoods(db *gorm.DB) ([]entity.Food, error) {
	var foods []entity.Food
	err := db.Find(&foods).Error
	return foods, err
}
