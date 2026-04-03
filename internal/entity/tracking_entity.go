package entity

import "time"

type Question struct {
	QuestionID   int              `gorm:"column:question_id;primaryKey;autoIncrement" json:"question_id"`
	ScheduleID   *int             `gorm:"column:schedule_id" json:"schedule_id"`
	QuestionText string           `gorm:"column:question_text" json:"question_text"`
	QuestionKey  string           `gorm:"column:question_key" json:"question_key"`
	Category     *string          `gorm:"column:category" json:"category"`
	InputType    *string          `gorm:"column:input_type" json:"input_type"`
	Unit         *string          `gorm:"column:unit" json:"unit"`
	IsRequired   *bool            `gorm:"column:is_required;default:true" json:"is_required"`
	DisplayOrder *int             `gorm:"column:display_order" json:"display_order"`
	CreatedAt    time.Time        `gorm:"column:created_at" json:"created_at"`
	QuestionOptions []QuestionOption `gorm:"foreignKey:QuestionID" json:"options"`
}

func (Question) TableName() string {
	return "questions"
}

type QuestionOption struct {
	OptionID     int     `gorm:"column:option_id;primaryKey;autoIncrement" json:"option_id"`
	QuestionID   int     `gorm:"column:question_id" json:"question_id"`
	OptionValue  string  `gorm:"column:option_value" json:"option_value"`
	OptionLabel  string  `gorm:"column:option_label" json:"option_label"`
	IconEmoji    *string `gorm:"column:icon_emoji" json:"icon_emoji"`
	DisplayOrder *int    `gorm:"column:display_order" json:"display_order"`
	IsDefault    *bool   `gorm:"column:is_default;default:false" json:"is_default"`
}


func (QuestionOption) TableName() string {
	return "question_options"
}

type UserTrackingLog struct {
	TrackingID       int       `gorm:"column:tracking_id;primaryKey;autoIncrement"`
	MotherID         *int      `gorm:"column:mother_id"`
	PregnancyID      *int      `gorm:"column:pregnancy_id"`
	ChildID          *int      `gorm:"column:child_id"`
	TrackingDate     time.Time `gorm:"column:tracking_date"`
	Frequency        string    `gorm:"column:frequency"`
	
	AnswersData      string    `gorm:"column:answers_data;type:json"`
	
	CompletionStatus *string   `gorm:"column:completion_status"`
	Notes            *string   `gorm:"column:notes"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}
