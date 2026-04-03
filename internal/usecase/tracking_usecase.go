package usecase

import (
	"encoding/json"
	"errors"
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
	"nutrimama/internal/repository"
	"nutrimama/internal/utils"
	"strings"
	"time"

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

// inferFrequency checks the question_key suffix to determine daily/weekly/monthly
func inferFrequency(key string) string {
	switch {
	case strings.Contains(key, "mingguan"):
		return "weekly"
	case strings.Contains(key, "bulanan"):
		return "monthly"
	default:
		// "harian" or anything else defaults to daily
		return "daily"
	}
}

// GetQuestions returns the target questions for today based on target type and elapsed days
func (u *TrackingUseCase) GetQuestions(userId int) ([]entity.Question, error) {
	mother, err := u.MotherRepo.FindByUserId(u.DB, userId)
	if err != nil {
		return nil, errors.New("only registered mothers can view tracking questions")
	}

	// 1. Determine target group (pregnant mother vs toddler/balita)
	target := "ibu_hamil" // default
	var pregnancy struct {
		Status string `gorm:"column:status"`
	}
	// Fetch the most recent pregnancy; check if it is still active
	u.DB.Raw("SELECT status FROM pregnancies WHERE mother_id = ? ORDER BY pregnancy_id DESC LIMIT 1", mother.MotherID).Scan(&pregnancy)
	if pregnancy.Status != "" && pregnancy.Status != "active" {
		target = "balita"
	}

	// 2. Calculate interval based on elapsed days since first log
	allowedFreqs := map[string]bool{
		"daily": true,
	}

	firstLog, firstErr := u.TrackingLogRepo.FindFirstByMotherID(u.DB, mother.MotherID)
	if firstErr == nil {
		// Calculate days elapsed from first tracking date
		today := time.Now().Truncate(24 * time.Hour)
		firstDate := firstLog.TrackingDate.Truncate(24 * time.Hour)
		daysElapsed := int(today.Sub(firstDate).Hours() / 24)

		if daysElapsed > 0 && daysElapsed%7 == 0 {
			allowedFreqs["weekly"] = true
		}
		if daysElapsed > 0 && daysElapsed%30 == 0 {
			allowedFreqs["monthly"] = true
		}
	}

	// 3. Fetch all questions for this target group
	allQuestions, err := u.TrackingQuestionRepo.FindByCategory(u.DB, target)
	if err != nil {
		return nil, err
	}

	// 4. Filter by the keyword frequency we detect from the key
	var filtered []entity.Question
	for _, q := range allQuestions {
		freq := inferFrequency(q.QuestionKey)
		if allowedFreqs[freq] {
			filtered = append(filtered, q)
		}
	}

	return filtered, nil
}
