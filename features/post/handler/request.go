package handler

import "simple-social-media-API/features/post"

type AddPostRequest struct {
	Content string `json:"content" form:"content"`
	Photo   string `json:"photo" form:"photo"`
}

type UpdatePostRequest struct {
	Content string `json:"content" form:"content"`
	Photo   string `json:"photo" form:"photo"`
}

func ToCore(data interface{}) *post.Core {
	res := post.Core{}

	switch data.(type) {
	case AddPostRequest:
		cnv := data.(AddPostRequest)
		res.Content = cnv.Content
		res.Photo = cnv.Photo
	case UpdatePostRequest:
		cnv := data.(UpdatePostRequest)
		res.Content = cnv.Content
		res.Photo = cnv.Photo
	default:
		return nil
	}

	return &res
}
