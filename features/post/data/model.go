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
		UserID:  data.UserID,
	}
}

func CoreToData(data post.Core) Post {
	return Post{
		Model:   gorm.Model{ID: data.ID},
		Content: data.Content,
		Photo:   data.Photo,
	}
}

// For MyPosts
func (dataModel *Post) ModelsToCore() post.Core {
	return post.Core{
		ID:      dataModel.ID,
		Content: dataModel.Content,
		Photo:   dataModel.Photo,
	}
}

func ListModelsToCore(dataModels []Post) []post.Core {
	var dataCore []post.Core
	for _, val := range dataModels {
		dataCore = append(dataCore, val.ModelsToCore())
	}
	return dataCore
}

// For All Posts
type PostWriter struct {
	ID      uint
	Content string
	Photo   string
	Writer  string
}

func (dataModel *PostWriter) AllModelsToCore() post.Core {
	return post.Core{
		ID:      dataModel.ID,
		Content: dataModel.Content,
		Photo:   dataModel.Photo,
		Writer:  dataModel.Writer,
	}
}

func ListAllModelsToCore(dataModels []PostWriter) []post.Core {
	var dataCore []post.Core
	for _, value := range dataModels {
		dataCore = append(dataCore, value.AllModelsToCore())
	}
	return dataCore
}
