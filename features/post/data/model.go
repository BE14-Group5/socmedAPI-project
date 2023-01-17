package data

import (
	"simple-social-media-API/features/post"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string
	Photo   string
	UserID  uint
}

func DataToCore(data Post) post.Core {
	return post.Core{
		ID:      data.ID,
		Content: data.Content,
		Photo:   data.Photo,
	}
}

func CoreToData(data post.Core) Post {
	return Post{
		Model:   gorm.Model{ID: data.ID},
		Content: data.Content,
		Photo:   data.Photo,
	}
}
