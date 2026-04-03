package model

import "time"

type InboxItemResponse struct {
	ConsultationID int       `json:"consultation_id"`
	PartnerName    string    `json:"partner_name"`
	PartnerRole    string    `json:"partner_role"`
	SessionDate    time.Time `json:"session_date"`
	TimeStart      string    `json:"time_start"`
	HourEnd        string    `json:"hour_end"`
	Status         string    `json:"status"`
}
