package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type ConsultantRepository struct {
	Repository[entity.Consultant]
}

func NewConsultantRepository() *ConsultantRepository {
	return &ConsultantRepository{
		Repository: Repository[entity.Consultant]{},
	}
}

func (r *ConsultantRepository) FindByUserId(db *gorm.DB, userId int) (*entity.Consultant, error) {
	var consultant entity.Consultant
	err := db.Where("user_id = ?", userId).First(&consultant).Error
	if err != nil {
		return nil, err
	}
	return &consultant, nil
}

func (r *ConsultantRepository) SearchAndPaginate(db *gorm.DB, search string, page int, limit int) ([]entity.Consultant, int64, error) {
	var consultants []entity.Consultant
	var total int64

	query := db.Model(&entity.Consultant{})

	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("full_name LIKE ? OR facility_name LIKE ?", searchTerm, searchTerm)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&consultants).Error; err != nil {
		return nil, 0, err
	}

	return consultants, total, nil
}
