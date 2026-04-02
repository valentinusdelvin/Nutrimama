package entity

import "time"

type User struct {
	ID     int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Email      string `gorm:"unique;column:email"`
	PhoneNumber string   `gorm:"column:phone_number"`
	Role        string    `gorm:"column:role"`
	Timezone    *string   `gorm:"column:timezone;default:'UTC'"`
	Password    string    `gorm:"column:password"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}