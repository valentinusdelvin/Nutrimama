package model

import "time"

type EduToolsResponse struct {
	ID        int       `json:"id"`
	Publisher string    `json:"publisher"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Category  string    `json:"category"`
	Thumbnail string    `json:"thumbnail"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateEduToolsRequest struct {
	Publisher string `json:"publisher" form:"publisher" binding:"required"`
	Title     string `json:"title" form:"title" binding:"required"`	
	Slug      string `json:"slug" form:"slug" binding:"required"`
	Category  string `json:"category" form:"category" binding:"required"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
	Content   string `json:"content" form:"content" binding:"required"`
}

type UpdateEduToolsRequest struct {
	ID        int    `json:"id" form:"id" binding:"required"`
	Publisher string `json:"publisher" form:"publisher"`
	Title     string `json:"title" form:"title"`
	Category  string `json:"category" form:"category"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
	Content   string `json:"content" form:"content"`
}
