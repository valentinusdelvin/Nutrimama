package entity

import "time"

type NutritionTracking struct {
	TrackID      int       `gorm:"column:track_id;primaryKey;autoIncrement"`
	MotherID     *int      `gorm:"column:mother_id"`
	ChildID      *int      `gorm:"column:child_id"`
	TrackingType *string   `gorm:"column:tracking_type"`
	ScoresData   string    `gorm:"column:scores_data"` // Flat string mapping
	CreatedAt    time.Time `gorm:"column:created_at"`
}

func (NutritionTracking) TableName() string {
	return "nutrition_tracking"
}
