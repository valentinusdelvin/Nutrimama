package usecase

import (
	"math"
	"nutrimama/internal/model"
	"nutrimama/internal/model/converter"
	"nutrimama/internal/repository"

	"gorm.io/gorm"
)

type ConsultantUseCase struct {
	DB                   *gorm.DB
	ConsultantRepository *repository.ConsultantRepository
}

func NewConsultantUseCase(db *gorm.DB, repo *repository.ConsultantRepository) *ConsultantUseCase {
	return &ConsultantUseCase{
		DB:                   db,
		ConsultantRepository: repo,
	}
}

func (u *ConsultantUseCase) ListConsultants(search string, page int, limit int) (*model.PaginatedConsultantResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	consultants, total, err := u.ConsultantRepository.SearchAndPaginate(u.DB, search, page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	response := &model.PaginatedConsultantResponse{
		Data:       converter.ConsultantListToResponse(consultants),
		TotalItems: total,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}

	return response, nil
}
