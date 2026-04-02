package usecase

import (
	"encoding/json"
	"errors"
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
	"nutrimama/internal/repository"
	"nutrimama/internal/utils"

	"gorm.io/gorm"
)

type TrackingUseCase struct {
	DB                    *gorm.DB
	TrackingLogRepo       *repository.TrackingLogRepository
	TrackingQuestionRepo  *repository.TrackingQuestionRepository
	NutritionTrackingRepo *repository.NutritionTrackingRepository
	MotherRepo            *repository.MotherRepository
}

func NewTrackingUseCase(
	db *gorm.DB,
	tlRepo *repository.TrackingLogRepository,
	tqRepo *repository.TrackingQuestionRepository,
	ntRepo *repository.NutritionTrackingRepository,
	mRepo *repository.MotherRepository,
) *TrackingUseCase {
	return &TrackingUseCase{
		DB:                    db,
		TrackingLogRepo:       tlRepo,
		TrackingQuestionRepo:  tqRepo,
		NutritionTrackingRepo: ntRepo,
		MotherRepo:            mRepo,
	}
}


func (u *TrackingUseCase) SubmitTracking(userId int, req model.SubmitTrackingRequest) error {
	c := u.DB.Begin()

	mother, err := u.MotherRepo.FindByUserId(c, userId)
	if err != nil {
		c.Rollback()
		return errors.New("only registered mothers can submit tracking logs")
	}

	log := entity.UserTrackingLog{
		MotherID:      &mother.MotherID,
		TrackingDate:  req.TrackingDate,
		Frequency:     req.Frequency,
		AnswersData:   req.AnswersData,
		Notes:         req.Notes,
	}
	
	if err := u.TrackingLogRepo.Create(c, &log); err != nil {
		c.Rollback()
		return err
	}

	finalScores := utils.CalculateNutrition(req.AnswersData, req.Frequency)
	finalScores.MotherID = &mother.MotherID

	if err := u.NutritionTrackingRepo.Create(c, finalScores); err != nil {
		c.Rollback()
		return err
	}

	return c.Commit().Error
}

func (u *TrackingUseCase) GetScores(userId int, date string) (map[string]float64, error) {
	c := u.DB

	mother, err := u.MotherRepo.FindByUserId(c, userId)
	if err != nil {
		return nil, errors.New("only registered mothers can view tracking scores")
	}

	record, err := u.NutritionTrackingRepo.FindByMotherIdDate(c, mother.MotherID, date)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return make(map[string]float64), nil
		}
		return nil, err
	}

	var scoresMap map[string]float64
	json.Unmarshal([]byte(record.ScoresData), &scoresMap)
	return scoresMap, nil
}
