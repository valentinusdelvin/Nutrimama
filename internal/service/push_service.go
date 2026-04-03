package service

import (
	"encoding/json"
	"log"
	"os"

	webpush "github.com/SherClockHolmes/webpush-go"
)

type PushService struct{}

func NewPushService() *PushService {
	return &PushService{}
}

// SendNotification sends a real Web Push (VAPID) notification to a browser subscription.
func (s *PushService) SendNotification(endpoint, p256dh, auth, title, message string) {
	payload, _ := json.Marshal(map[string]string{
		"title": title,
		"body":  message,
	})

	sub := &webpush.Subscription{
		Endpoint: endpoint,
		Keys: webpush.Keys{
			P256dh: p256dh,
			Auth:   auth,
		},
	}

	resp, err := webpush.SendNotification(payload, sub, &webpush.Options{
		VAPIDPrivateKey: os.Getenv("VAPID_PRIVATE_KEY"),
		VAPIDPublicKey:  os.Getenv("VAPID_PUBLIC_KEY"),
		Subscriber:      os.Getenv("VAPID_EMAIL"),
		TTL:             30,
	})

	if err != nil {
		log.Printf("[PUSH ERROR] Failed to send to %s: %v\n", endpoint, err)
		return
	}
	defer resp.Body.Close()
	log.Printf("[PUSH SENT] Status %d → %s\n", resp.StatusCode, endpoint)
}
