package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type MessageRepository struct {
	Repository[entity.Message]
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		Repository: Repository[entity.Message]{},
	}
}

func (r *MessageRepository) FindByConsultationId(db *gorm.DB, consultationId int) ([]entity.Message, error) {
	var records []entity.Message
	err := db.Where("consultation_id = ?", consultationId).Order("created_at asc").Find(&records).Error
	return records, err
}
