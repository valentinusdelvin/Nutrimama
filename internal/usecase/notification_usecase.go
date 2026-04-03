package usecase

import (
	"errors"
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

// Subscribe upserts a VAPID push subscription for the given user.
func (u *NotificationUseCase) Subscribe(userId int, req model.SubscribePushRequest) error {
	sub := entity.PushSubscription{
		UserID:   userId,
		Endpoint: req.Endpoint,
		P256dh:   req.P256dh,
		Auth:     req.Auth,
	}

	// Upsert by endpoint — same browser re-subscribing just updates user association
	return u.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "endpoint"}},
		DoUpdates: clause.AssignmentColumns([]string{"user_id", "p256dh", "auth"}),
	}).Create(&sub).Error
}

// GetMyNotifications fetches the localized bell inbox messages for the web app UI
func (u *NotificationUseCase) GetMyNotifications(userId int) ([]entity.Notification, error) {
	return u.NotificationRepo.FindByUserId(u.DB, userId)
}

// SendTestNotification fires a dummy push to a specific subscription ID for testing.
func (u *NotificationUseCase) SendTestNotification(userId int, subscriptionId int, pushService interface{ SendNotification(endpoint, p256dh, auth, title, message string) }) error {
	var sub entity.PushSubscription
	err := u.DB.Where("subscription_id = ? ", subscriptionId).First(&sub).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("subscription not found for this user")
		}
		return err
	}

	pushService.SendNotification(
		sub.Endpoint,
		sub.P256dh,
		sub.Auth,
		"🔔 Nutrimama Test Notification",
		"This is a dummy notification to confirm your push subscription is working!",
	)
	return nil
}
