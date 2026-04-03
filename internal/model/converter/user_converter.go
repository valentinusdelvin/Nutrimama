package converter

import (
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
)

func UserToResponse(user *entity.User, token string, name string) *model.UserResponse {
	return &model.UserResponse{
		UserId:      user.ID,
		Email:       user.Email,
		Name:        name,
		Role:        user.Role,
		Token:     	 token,
	}
}
