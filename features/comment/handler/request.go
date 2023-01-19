package handler

import (
	"simple-social-media-API/features/comment"
)

type AddCommentReq struct {
	PostId  uint   `json:"post_id" form:"post_id"`
	Content string `json:"content" form:"content"`
}

type DeleteCommentReq struct {
	PostId uint `json:"post_id" form:"post_id"`
}

func ToCore(data interface{}) *comment.Core {
	res := comment.Core{}

	switch data.(type) {
	case AddCommentReq:
		cnv := data.(AddCommentReq)
		res.PostId = cnv.PostId
		res.Content = cnv.Content
	case DeleteCommentReq:
		cnv := data.(DeleteCommentReq)
		res.PostId = cnv.PostId
	default:
		return nil
	}
	return &res
}
