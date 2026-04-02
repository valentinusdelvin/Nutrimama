package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type ConsultationSessionRepository struct {
	Repository[entity.ConsultationSession]
}

func NewConsultationSessionRepository() *ConsultationSessionRepository {
	return &ConsultationSessionRepository{
		Repository: Repository[entity.ConsultationSession]{},
	}
}

func (r *ConsultationSessionRepository) FindById(db *gorm.DB, id int) (*entity.ConsultationSession, error) {
	var record entity.ConsultationSession
	err := db.Where("consultation_id = ?", id).First(&record).Error
	return &record, err
}

func (r *ConsultationSessionRepository) FindOverlapping(db *gorm.DB, consultantId int, date string, start string, end string) (int64, error) {
	var count int64
	// Mathematical bounding determining boolean logic checking explicitly for overlaps safely!
	err := db.Model(&entity.ConsultationSession{}).
		Where("consultant_id = ? AND session_date = ? AND time_start < ? AND hour_end > ? AND status != 'ended'", 
		consultantId, date, end, start).Count(&count).Error
	return count, err
}
