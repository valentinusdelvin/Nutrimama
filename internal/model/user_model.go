package model

type RegisterRequest struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Role        string `json:"role" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	UserId      int    `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Token       string `json:"token"`
}
