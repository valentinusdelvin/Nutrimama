package entity

type FoodLog struct {
	FoodLogID  int    `gorm:"column:food_log_id;primaryKey;autoIncrement"`
	MealPlanID int    `gorm:"column:meal_plan_id"`
	FoodID     int    `gorm:"column:food_id"`
	FoodName   string `gorm:"-" json:"food_name"`
	FoodImage  string `gorm:"-" json:"food_image"`
	LogDate    string `gorm:"column:log_date"` // YYYY-MM-DD
	DayName    string `gorm:"-" json:"day_name"`
	Eaten      bool   `gorm:"column:eaten"`
	MealTime   string `gorm:"column:meal_time"`
}

func (FoodLog) TableName() string {
	return "food_logs"
}
