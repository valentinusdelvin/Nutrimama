package model

type ConsultantResponse struct {
	ConsultantID int      `json:"consultant_id"`
	UserID       int      `json:"user_id"`
	FullName     string   `json:"full_name"`
	FacilityName *string  `json:"facility_name,omitempty"`
	Rating       *float64 `json:"rating,omitempty"`
	HourlyRate   float64  `json:"hourly_rate"`
}

type PaginatedConsultantResponse struct {
	Data       []ConsultantResponse `json:"data"`
	TotalItems int64                  `json:"total_items"`
	TotalPages int                    `json:"total_pages"`
	Page       int                    `json:"page"`
	Limit      int                    `json:"limit"`
}
