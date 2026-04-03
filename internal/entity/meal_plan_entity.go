package entity

type MealPlan struct {
	MealPlanID int    `gorm:"column:meal_plan_id;primaryKey;autoIncrement"`
	MotherID   int    `gorm:"column:mother_id"`
	Week       int    `gorm:"column:week"`
	Phase      string `gorm:"column:phase"`
}

func (MealPlan) TableName() string {
	return "meal_plans"
}
