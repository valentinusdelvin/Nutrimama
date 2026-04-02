package usecase

import (
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
	"nutrimama/internal/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NotificationUseCase struct {
	DB               *gorm.DB
	NotificationRepo *repository.NotificationRepository
	SubscriptionRepo *repository.PushSubscriptionRepository
}

func NewNotificationUseCase(db *gorm.DB, notifRepo *repository.NotificationRepository, subRepo *repository.PushSubscriptionRepository) *NotificationUseCase {
	return &NotificationUseCase{
		DB:               db,
		NotificationRepo: notifRepo,
		SubscriptionRepo: subRepo,
	}
}

// Subscribe upserts a device push token for the given user, allowing the CRON job to target their device.
func (u *NotificationUseCase) Subscribe(userId int, req model.SubscribePushRequest) error {
	sub := entity.PushSubscription{
		UserID:      userId,
		DeviceToken: req.DeviceToken,
		Platform:    &req.Platform,
	}

	// We use an Upsert logic (OnConflict) because if the user logs out and logs in 5 times,
	// the device token is the same, so we simply update the associated user_id for that token!
	return u.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "device_token"}},
		DoUpdates: clause.AssignmentColumns([]string{"user_id", "platform"}),
	}).Create(&sub).Error
}

// GetMyNotifications fetches the localized bell inbox messages for the web app UI
func (u *NotificationUseCase) GetMyNotifications(userId int) ([]entity.Notification, error) {
	return u.NotificationRepo.FindByUserId(u.DB, userId)
}
