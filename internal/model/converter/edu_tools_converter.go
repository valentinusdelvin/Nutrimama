package converter

import (
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
)

func EduToolsToResponse(eduTools *entity.EduTools) *model.EduToolsResponse {
	return &model.EduToolsResponse{
		ID:        eduTools.ID,
		Publisher: eduTools.Publisher,
		Title:     eduTools.Title,
		Slug:      eduTools.Slug,
		Category:  eduTools.Category,
		Thumbnail: eduTools.Thumbnail,
		Content:   eduTools.Content,
		CreatedAt: eduTools.CreatedAt,
	}
}

func EduToolsListToResponse(eduToolsList []entity.EduTools) []model.EduToolsResponse {
	var responses []model.EduToolsResponse
	for _, eduTools := range eduToolsList {
		responses = append(responses, *EduToolsToResponse(&eduTools))
	}
	return responses
}
