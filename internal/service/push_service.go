package service

import (
	"log"
)

type PushService struct {
	// If you later use Firebase (FCM), you will store the FCM Client here.
	// If you use standard Web Push (VAPID), you store the VAPID keys here!
}

func NewPushService() *PushService {
	return &PushService{}
}

func (s *PushService) SendNotification(deviceToken string, title string, message string) {
	// =========================================================
	// PUSH NOTIFICATION SENDER
	// =========================================================
	// Depending on your UI preference, here you will implement:
	//   A. webpush.SendNotification() [If using VAPID]
	//   B. messaging.Send()           [If using Firebase FCM]
	
	log.Printf("[PUSH SENT] To token: %s | Title: %s | Msg: %s\n", deviceToken, title, message)
}
