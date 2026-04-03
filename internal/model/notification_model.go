package model

type SubscribePushRequest struct {
	Endpoint string `json:"endpoint" binding:"required"`
	P256dh   string `json:"p256dh" binding:"required"`
	Auth     string `json:"auth" binding:"required"`
}

type SendTestPushRequest struct {
	SubscriptionID int `json:"subscription_id" binding:"required"`
}
