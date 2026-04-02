package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	Repository[entity.Notification]
}

func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{
		Repository: Repository[entity.Notification]{},
	}
}

func (r *NotificationRepository) FindByUserId(db *gorm.DB, userId int) ([]entity.Notification, error) {
	var notifications []entity.Notification
	err := db.Where("user_id = ?", userId).Order("created_at desc").Find(&notifications).Error
	return notifications, err
}

func (r *NotificationRepository) MarkAsRead(db *gorm.DB, notificationId int, userId int) error {
	return db.Model(&entity.Notification{}).
		Where("notification_id = ? AND user_id = ?", notificationId, userId).
		Update("is_read", true).Error
}
