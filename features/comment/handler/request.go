package handler

import (
	"simple-social-media-API/features/comment"
)

type AddCommentReq struct {
	PostId  uint   `json:"post_id" form:"post_id"`
	Content string `json:"content" form:"content"`
}

func ToCore(data interface{}) *comment.Core {
	res := comment.Core{}

	switch data.(type) {
	case AddCommentReq:
		cnv := data.(AddCommentReq)
		res.PostId = cnv.PostId
		res.Content = cnv.Content
	default:
		return nil
	}
	return &res
}
