package post

import (
	"simple-social-media-API/features/comment"
	"time"
)

type MyPostsResp struct {
	ID        uint
	Content   string
	Photo     string
	UserID    uint
	Writer    string
	CreatedAt time.Time
	Comments  []comment.Core
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
