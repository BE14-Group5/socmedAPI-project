package data

import (
	"errors"
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
	convert := CoreToData(updatedPost)
	qry := pd.db.Where("id = ? AND user_id = ?", postID, userID).Updates(&convert)
	if qry.RowsAffected <= 0 {
		log.Println("update post query error : data not found")
		return post.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update post query error :", err.Error())
		return post.Core{}, errors.New("not found")
	}

	return DataToCore(convert), nil
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
