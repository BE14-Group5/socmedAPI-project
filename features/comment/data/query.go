package data

import (
	"log"
	"simple-social-media-API/features/comment"

	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.CommentData {
	return &commentData{
		db: db,
	}
}

func (cd *commentData) Add(newComment comment.Core) (comment.Core, error) {
	cnv := CoreToData(newComment)

	err := cd.db.Create(&cnv).Error
	if err != nil {
		log.Println("error add comment query: ", err)
		return comment.Core{}, err
	}
	newComment.ID = cnv.ID
	newComment.CreatedAt = cnv.CreatedAt
	return newComment, nil
}
func (cd *commentData) GetComments() ([]comment.Core, error) {
	return []comment.Core{}, nil
}
func (cd *commentData) Update(updComment comment.Core) (comment.Core, error) {
	return comment.Core{}, nil
}
func (cd *commentData) Delete(userId, postId, commentId uint) error {
	return nil
}
