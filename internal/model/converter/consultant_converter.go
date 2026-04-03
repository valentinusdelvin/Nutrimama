package converter

import (
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
)

func ConsultantToResponse(consultant *entity.Consultant) *model.ConsultantResponse {
	return &model.ConsultantResponse{
		ConsultantID: consultant.ConsultantID,
		UserID:       consultant.UserID,
		FullName:     consultant.FullName,
		FacilityName: consultant.FacilityName,
		Rating:       consultant.Rating,
		HourlyRate:   consultant.HourlyRate,
	}
}

func ConsultantListToResponse(consultants []entity.Consultant) []model.ConsultantResponse {
	var responses []model.ConsultantResponse
	for _, c := range consultants {
		responses = append(responses, *ConsultantToResponse(&c))
	}
	return responses
}
