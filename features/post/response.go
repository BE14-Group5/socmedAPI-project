package post

import (
	"simple-social-media-API/features/comment"
	"time"
)

type MyPostsResp struct {
	ID        uint           `json:"id" form:"id"`
	Content   string         `json:"content" form:"content"`
	Photo     string         `json:"photo" form:"photo"`
	UserID    uint           `json:"user_id" form:"user_id"`
	Writer    string         `json:"user_name" form:"user_name"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	Comments  []comment.Core `json:"comments" form:"comments"`
}

func ToMyPostResp(data Core) MyPostsResp {
	return MyPostsResp{
		ID:        data.ID,
		Content:   data.Content,
		Photo:     data.Photo,
		UserID:    data.UserID,
		Writer:    data.Writer,
		CreatedAt: data.CreatedAt,
	}
}
