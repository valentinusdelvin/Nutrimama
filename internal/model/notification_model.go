package model

type SubscribePushRequest struct {
	DeviceToken string `json:"device_token" binding:"required"`
	Platform    string `json:"platform"` // e.g. web, android, ios
}
