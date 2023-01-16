package handler

import (
	"simple-social-media-API/features/post"
)

type PostResponse struct {
	ID     uint   `json:"id"`
	Body   string `json:"body"`
	Image  string `json:"image"`
	Writer string `json:"writer"`
}

type AddPostResponse struct {
	ID     uint   `json:"id"`
	Body   string `json:"body"`
	Image  string `json:"image"`
	Writer string `json:"writer"`
}

func AddPostToResponse(dataCore post.Core) AddPostResponse {
	return AddPostResponse{
		ID:     dataCore.ID,
		Body:   dataCore.Body,
		Image:  dataCore.Image,
		Writer: dataCore.Writer,
	}
}
