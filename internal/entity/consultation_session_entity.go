package entity

import "time"

type ConsultationSession struct {
	ConsultationID int       `gorm:"column:consultation_id;primaryKey;autoIncrement"`
	MotherID       int       `gorm:"column:mother_id"`
	ConsultantID   int       `gorm:"column:consultant_id"`
	SessionDate    time.Time `gorm:"column:session_date"` // MySQL DATE translates flawlessly to Time
	TimeStart      string    `gorm:"column:time_start"`   // HH:MM:SS format 
	HourEnd        string    `gorm:"column:hour_end"`     // HH:MM:SS format
	TotalPrice     float64   `gorm:"column:total_price"`
	PaymentStatus  string    `gorm:"column:payment_status"`
	Status         string    `gorm:"column:status"`

	Mother     *Mother     `gorm:"foreignKey:MotherID;references:MotherID"`
	Consultant *Consultant `gorm:"foreignKey:ConsultantID;references:ConsultantID"`
}

func (ConsultationSession) TableName() string {
	return "consultation_sessions"
}
