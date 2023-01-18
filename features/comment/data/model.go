package data

import (
	"simple-social-media-API/features/comment"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId  uint
	PostId  uint
	Content string
}

func ToCore(data Comment) comment.Core {
	return comment.Core{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UserId:    data.UserId,
		PostId:    data.PostId,
		Content:   data.Content,
	}
}

func CoreToData(data comment.Core) Comment {
	return Comment{
		Model: gorm.Model{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
		},
		UserId:  data.UserId,
		PostId:  data.PostId,
		Content: data.Content,
	}
}
