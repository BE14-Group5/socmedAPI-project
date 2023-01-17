package handler

import (
	"simple-social-media-API/features/post"
)

type PostResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
	UserID  uint   `json:"user_id"`
}

type AddPostResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
	UserID  uint   `json:"user_id"`
}

func AddPostToResponse(dataCore post.Core) AddPostResponse {
	return AddPostResponse{
		ID:      dataCore.ID,
		Content: dataCore.Content,
		Photo:   dataCore.Photo,
		UserID:  dataCore.UserID,
	}
}

type UpdatePostResponse struct {
	// ID      uint   `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
}

func UpdatePostToResponse(dataCore post.Core) UpdatePostResponse {
	return UpdatePostResponse{
		// ID:      dataCore.ID,
		Content: dataCore.Content,
		Photo:   dataCore.Photo,
	}
}
