package handler

import "simple-social-media-API/features/post"

type AddPostRequest struct {
	Body  string `json:"body" form:"body"`
	Image string `json:"image" form:"image"`
}

func ToCore(data interface{}) *post.Core {
	res := post.Core{}

	switch data.(type) {
	case AddPostRequest:
		cnv := data.(AddPostRequest)
		res.Body = cnv.Body
		res.Image = cnv.Image
	default:
		return nil
	}

	return &res
}
