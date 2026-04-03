package entity

import "time"

type PushSubscription struct {
	SubscriptionID int       `gorm:"column:subscription_id;primaryKey;autoIncrement"`
	UserID         int       `gorm:"column:user_id"`
	Endpoint       string    `gorm:"column:endpoint"`
	P256dh         string    `gorm:"column:p256dh"`
	Auth           string    `gorm:"column:auth"`
	Platform       *string   `gorm:"column:platform"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}

type Notification struct {
	NotificationID int       `gorm:"column:notification_id;primaryKey;autoIncrement"`
	UserID         int       `gorm:"column:user_id"`
	Title          string    `gorm:"column:title"`
	Message        string    `gorm:"column:message"`
	IsRead         *bool     `gorm:"column:is_read;default:false"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}
