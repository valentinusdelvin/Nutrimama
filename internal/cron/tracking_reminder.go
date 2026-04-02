package cron

import (
	"log"
	"time"

	"nutrimama/internal/entity"
	"nutrimama/internal/repository"
	"nutrimama/internal/service"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type TrackingReminderCron struct {
	DB               *gorm.DB
	NotificationRepo *repository.NotificationRepository
	PushSubRepo      *repository.PushSubscriptionRepository
	PushService      *service.PushService
}

func NewTrackingReminderCron(
	db *gorm.DB,
	nRepo *repository.NotificationRepository,
	psRepo *repository.PushSubscriptionRepository,
	pushSvc *service.PushService,
) *TrackingReminderCron {
	return &TrackingReminderCron{
		DB:               db,
		NotificationRepo: nRepo,
		PushSubRepo:      psRepo,
		PushService:      pushSvc,
	}
}

func (c *TrackingReminderCron) Start() {
	scheduler := cron.New()
	_, err := scheduler.AddFunc("0 * * * *", c.processDailyReminders)
	if err != nil {
		log.Fatalf("failed to initialize tracking cron: %v", err)
	}
	scheduler.Start()
	log.Println("[CRON] TrackingReminder started successfully. Will run hourly.")
}

func (c *TrackingReminderCron) processDailyReminders() {
	log.Println("[CRON] Running hourly 7 PM tracking reminder check...")

	var mothers []struct {
		UserID   int
		MotherID int
		Timezone string
	}

	err := c.DB.Table("users").
		Select("users.user_id, mothers.mother_id, users.timezone").
		Joins("JOIN mothers ON users.user_id = mothers.user_id").
		Scan(&mothers).Error

	if err != nil {
		log.Println("[CRON ERROR] Failed fetching mothers: ", err)
		return
	}

	for _, m := range mothers {
		tz := "UTC"
		if m.Timezone != "" {
			tz = m.Timezone
		}

		loc, err := time.LoadLocation(tz)
		if err != nil {
			loc = time.UTC
		}

		localTime := time.Now().In(loc)

		// 2. Check if it is currently exactly the 7 PM hour in their local time zone!
		if localTime.Hour() == 19 {
			c.checkAndNotifyMother(m.UserID, m.MotherID, localTime)
		}
	}
}

func (c *TrackingReminderCron) checkAndNotifyMother(userId, motherId int, localTime time.Time) {
	todayStr := localTime.Format("2006-01-02")

	var count int64
	c.DB.Table("user_tracking_logs").
		Where("mother_id = ? AND tracking_date = ? AND frequency = 'daily'", motherId, todayStr).
		Count(&count)

	if count == 0 {
		
		notif := entity.Notification{
			UserID:  userId,
			Title:   "Daily Tracking Reminder",
			Message: "It's 7 PM! Don't forget to submit your nutrition tracking for today.",
		}
		c.NotificationRepo.Create(c.DB, &notif)

		subs, _ := c.PushSubRepo.FindByUserId(c.DB, userId)

		for _, sub := range subs {
			c.PushService.SendNotification(sub.DeviceToken, notif.Title, notif.Message)
		}
	}
}
