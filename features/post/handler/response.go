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

type MyPostsResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
}

func MyPostsToResponse(dataCore post.Core) MyPostsResponse {
	return MyPostsResponse{
		ID:      dataCore.ID,
		Content: dataCore.Content,
		Photo:   dataCore.Photo,
	}
}

func ListMyPostsToResponse(dataCore []post.Core) []MyPostsResponse {
	var DataResponse []MyPostsResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, MyPostsToResponse(value))
	}
	return DataResponse
}

type AllPostResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
	Writer  string `json:"writer"`
}

// For AllPosts
func AllPostsToResponse(dataCore post.Core) AllPostResponse {
	return AllPostResponse{
		ID:      dataCore.ID,
		Content: dataCore.Content,
		Photo:   dataCore.Photo,
		Writer:  dataCore.Writer,
	}
}
func ListAllPostsToResponse(dataCore []post.Core) []AllPostResponse {
	var DataResponse []AllPostResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, AllPostsToResponse(value))
	}
	return DataResponse
}
