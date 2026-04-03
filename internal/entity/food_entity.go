package entity

import "encoding/json"

type Food struct {
	FoodID        int      `gorm:"column:food_id;primaryKey;autoIncrement" json:"food_id"`
	Name          string   `gorm:"column:name" json:"name"`
	Carbohydrates float64  `gorm:"column:carbohydrates" json:"carbohydrates"`
	Fat           float64  `gorm:"column:fat" json:"fat"`
	Category      string   `gorm:"column:category" json:"category"`
	Image         *string  `gorm:"column:image" json:"-"` // nullable in DB
}

// MarshalJSON overrides JSON serialization so Image is never null in the response
func (f Food) MarshalJSON() ([]byte, error) {
	image := "food_placeholder.jpg"
	if f.Image != nil && *f.Image != "" {
		image = *f.Image
	}

	type Alias Food // prevent infinite recursion
	return json.Marshal(&struct {
		Alias
		Image string `json:"image"`
	}{
		Alias: Alias(f),
		Image: image,
	})
}

func (Food) TableName() string {
	return "foods"
}
