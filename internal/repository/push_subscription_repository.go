package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type PushSubscriptionRepository struct {
	Repository[entity.PushSubscription]
}

func NewPushSubscriptionRepository() *PushSubscriptionRepository {
	return &PushSubscriptionRepository{
		Repository: Repository[entity.PushSubscription]{},
	}
}

func (r *PushSubscriptionRepository) FindByUserId(db *gorm.DB, userId int) ([]entity.PushSubscription, error) {
	var subs []entity.PushSubscription
	err := db.Where("user_id = ?", userId).Find(&subs).Error
	return subs, err
}
