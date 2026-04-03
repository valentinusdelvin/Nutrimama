package usecase

import (
	"encoding/json"
	"errors"
	"nutrimama/internal/entity"
	"nutrimama/internal/repository"
	"nutrimama/internal/utils"
	"time"

	"gorm.io/gorm"
)

type MealPlanUseCase struct {
	DB               *gorm.DB
	FoodRepo         *repository.FoodRepository
	TrackingRepo     *repository.NutritionTrackingRepository
	MotherRepo       *repository.MotherRepository
}

func NewMealPlanUseCase(db *gorm.DB, fRepo *repository.FoodRepository, tRepo *repository.NutritionTrackingRepository, mRepo *repository.MotherRepository) *MealPlanUseCase {
	return &MealPlanUseCase{
		DB:           db,
		FoodRepo:     fRepo,
		TrackingRepo: tRepo,
		MotherRepo:   mRepo,
	}
}

func (u *MealPlanUseCase) GenerateAIMealPlan(userId int) ([]entity.FoodLog, error) {
	// 1. Resolve exact physical mother ID cleanly
	mother, err := u.MotherRepo.FindByUserId(u.DB, userId)
	if err != nil {
		return nil, errors.New("unauthorized mapping mother structural access intelligently")
	}

	// 2. Extract mother's tracking conditions securely mapping JSON scores generically!
	latestTracking, err := u.TrackingRepo.GetLatestTracking(u.DB, mother.MotherID)
	if err != nil {
		return nil, errors.New("unable to discover tracking data. please submit a nutrition tracker first safely")
	}
	scoresContext := string(latestTracking.ScoresData)

	// 3. Extract the entire DB Food Context natively
	allFoods, err := u.FoodRepo.GetAllFoods(u.DB)
	if err != nil {
		return nil, errors.New("failed physically grabbing foods arrays safely")
	}
	foodsJSON, _ := json.Marshal(allFoods)

	// 4. Force AI to calculate exactly 7 days (21 meals) optimally mathematically!
	days := 7
	foodIDs, err := utils.GenerateMealPlanIDs(scoresContext, string(foodsJSON), days)
	if err != nil {
		return nil, err
	}

	// 5. Create quick lookup maps for name and image from the foods we already fetched
	foodNameMap := make(map[int]string)
	foodImageMap := make(map[int]string)
	for _, f := range allFoods {
		foodNameMap[f.FoodID] = f.Name
		if f.Image != nil && *f.Image != "" {
			foodImageMap[f.FoodID] = *f.Image
		} else {
			foodImageMap[f.FoodID] = "/uploads/food_placeholder.jpg"
		}
	}

	// 6. Build and isolate the exact Relational Logic smoothly securely via GORM Transactions
	tx := u.DB.Begin()

	mealPlan := entity.MealPlan{
		MotherID: mother.MotherID,
		Week:     1, // Optimally dynamically iterated normally
		Phase:    "AI-Generated Protocol",
	}

	if err := tx.Create(&mealPlan).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var savedLogs []entity.FoodLog

	startDay := time.Now()
	for i := 0; i < days; i++ {
		logDate := startDay.AddDate(0, 0, i).Format("2006-01-02")
		dayName := startDay.AddDate(0, 0, i).Format("Monday")
		
		// 3 Meals a day accurately parsing the strict JSON output mapped
		baseIndex := i * 3
		
		mealTimes := []string{"08:00:00", "12:30:00", "18:00:00"}
		for j := 0; j < 3; j++ {
			fID := foodIDs[baseIndex+j]
			log := entity.FoodLog{
				MealPlanID: mealPlan.MealPlanID,
				FoodID:     fID,
				FoodName:   foodNameMap[fID],
				FoodImage:  foodImageMap[fID],
				LogDate:    logDate,
				DayName:    dayName,
				Eaten:      false,
				MealTime:   mealTimes[j],
			}
			
			if err := tx.Create(&log).Error; err != nil {
				tx.Rollback()
				return nil, errors.New("crashed persisting physical logs uniquely")
			}
			
			savedLogs = append(savedLogs, log)
		}
	}

	tx.Commit()
	return savedLogs, nil
}

func (u *MealPlanUseCase) GetLatestMealPlan(userId int) ([]entity.FoodLog, error) {
	// 1. Resolve mother
	mother, err := u.MotherRepo.FindByUserId(u.DB, userId)
	if err != nil {
		return nil, errors.New("unauthorized mapping mother structural access intelligently")
	}

	// 2. Find the latest meal plan record
	var mealPlan entity.MealPlan
	if err := u.DB.Where("mother_id = ?", mother.MotherID).Order("meal_plan_id DESC").First(&mealPlan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no meal plan found. please generate one first")
		}
		return nil, err
	}

	// 3. Fetch logs and JOIN with foods to get names
	var logs []entity.FoodLog
	// We use strings for dates, so we can sort them to find the "max" date
	if err := u.DB.Where("meal_plan_id = ?", mealPlan.MealPlanID).Order("log_date ASC, meal_time ASC").Find(&logs).Error; err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, errors.New("meal plan contains no entries")
	}

	latestLogDate := logs[len(logs)-1].LogDate
	today := time.Now().Format("2006-01-02")
	if today > latestLogDate {
		return nil, errors.New("your previous meal plan has expired. please generate a new one for this week")
	}

	allFoods, _ := u.FoodRepo.GetAllFoods(u.DB)
	foodNameMap := make(map[int]string)
	foodImageMap := make(map[int]string)
	for _, f := range allFoods {
		foodNameMap[f.FoodID] = f.Name
		if f.Image != nil && *f.Image != "" {
			foodImageMap[f.FoodID] = *f.Image
		} else {
			foodImageMap[f.FoodID] = "/uploads/food_placeholder.jpg"
		}
	}

	for i := range logs {
		logs[i].FoodName = foodNameMap[logs[i].FoodID]
		logs[i].FoodImage = foodImageMap[logs[i].FoodID]

		dateStr := logs[i].LogDate
		if len(dateStr) >= 10 {
			dateStr = dateStr[:10]
		}
		t, _ := time.Parse("2006-01-02", dateStr)
		logs[i].DayName = t.Format("Monday")
	}

	return logs, nil
}
