package entity

type Consultant struct {
	ConsultantID int      `gorm:"column:consultant_id;primaryKey;autoIncrement"`
	UserID       int      `gorm:"column:user_id"`
	FullName     string   `gorm:"column:full_name"`
	FacilityName *string  `gorm:"column:facility_name"`
	Rating       *float64 `gorm:"column:rating"`
	HourlyRate   float64  `gorm:"column:hourly_rate"`
}
