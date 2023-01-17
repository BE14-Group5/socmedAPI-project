package handler

import (
	"simple-social-media-API/features/post"
)

type PostResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
	Writer  string `json:"writer"`
}

type AddPostResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
	Writer  string `json:"writer"`
}

func AddPostToResponse(dataCore post.Core) AddPostResponse {
	return AddPostResponse{
		ID:      dataCore.ID,
		Content: dataCore.Content,
		Photo:   dataCore.Photo,
		Writer:  dataCore.Writer,
	}
}

type UpdatePostResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
}

func UpdatePostToResponse(dataCore post.Core) UpdatePostResponse {
	return UpdatePostResponse{
		ID:      dataCore.ID,
		Content: dataCore.Content,
		Photo:   dataCore.Photo,
	}
}
