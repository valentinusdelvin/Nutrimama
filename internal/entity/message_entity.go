package entity

import "time"

type Message struct {
	MessageID      int       `gorm:"column:message_id;primaryKey;autoIncrement"`
	ConsultationID *int      `gorm:"column:consultation_id"`
	ConsultantID   *int      `gorm:"column:consultant_id"`
	MotherID       *int      `gorm:"column:mother_id"`
	Message        string    `gorm:"column:message"`
	Status         string    `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}

func (Message) TableName() string {
	return "messages"
}
