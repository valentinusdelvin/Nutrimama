package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type TrackingLogRepository struct {
	Repository[entity.UserTrackingLog]
}

func NewTrackingLogRepository() *TrackingLogRepository {
	return &TrackingLogRepository{
		Repository: Repository[entity.UserTrackingLog]{},
	}
}

// FindFirstByMotherID returns the earliest tracking log for a mother to calculate interval
func (r *TrackingLogRepository) FindFirstByMotherID(db *gorm.DB, motherID int) (*entity.UserTrackingLog, error) {
	var log entity.UserTrackingLog
	err := db.Where("mother_id = ?", motherID).Order("tracking_date ASC").First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}
