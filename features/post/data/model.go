package data

import (
	"simple-social-media-API/features/post"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Body   string
	Image  string
	UserID uint
}

func DataToCore(data Post) post.Core {
	return post.Core{
		ID:    data.ID,
		Body:  data.Body,
		Image: data.Image,
	}
}

func CoreToData(data post.Core) Post {
	return Post{
		Model: gorm.Model{ID: data.ID},
		Body:  data.Body,
		Image: data.Image,
	}
}
