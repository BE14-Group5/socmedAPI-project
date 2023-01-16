package data

import (
	"log"
	"simple-social-media-API/features/post"

	"gorm.io/gorm"
)

type postData struct {
	db *gorm.DB
}

func Isolation(db *gorm.DB) post.PostData {
	return &postData{
		db: db,
	}
}

func (pd *postData) Add(userID uint, newPost post.Core) (post.Core, error) {
	convert := CoreToData(newPost)
	convert.UserID = userID

	err := pd.db.Create(&convert).Error
	if err != nil {
		log.Println("add book query error")
		return post.Core{}, err
	}

	newPost.ID = convert.ID

	return newPost, nil
}

func (pd *postData) Update(postID uint, userID uint, updatedPost post.Core) (post.Core, error) {
	return post.Core{}, nil
}

func (pd *postData) Delete(postID uint, userID uint) error {
	return nil
}

func (pd *postData) MyPosts(userID uint) ([]post.Core, error) {
	return []post.Core{}, nil
}

func (pd *postData) AllPosts() ([]post.Core, error) {
	return []post.Core{}, nil
}
