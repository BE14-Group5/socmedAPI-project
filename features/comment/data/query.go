package data

import (
	"errors"
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
	// newComment.ID = cnv.ID
	// newComment.CreatedAt = cnv.CreatedAt

	retComm, _ := cd.getComment(cnv.ID)

	return retComm, nil
}
func (cd *commentData) getComment(commentId uint) (comment.Core, error) {
	comments := comment.Core{}
	qry := cd.db.Raw("SELECT c.id, user_id, name UserName, post_id, c.created_at, content FROM comments c JOIN users u ON c.user_id = u.id WHERE c.deleted_at IS NULL AND c.id = ?", commentId).Scan(&comments)
	if affrows := qry.RowsAffected; affrows <= 0 {
		return comment.Core{}, errors.New("empty comment")
	}
	if err := qry.Error; err != nil {
		log.Println("error query: ", err.Error())
		return comment.Core{}, err
	}
	return comments, nil
}
func (cd *commentData) GetComments(postId uint) ([]comment.Core, error) {
	comments := []comment.Core{}
	qry := cd.db.Raw("SELECT c.id, user_id, name UserName, post_id, c.created_at, content FROM comments c JOIN users u ON c.user_id = u.id WHERE c.deleted_at IS NULL AND post_id = ?", postId).Scan(&comments)
	if affrows := qry.RowsAffected; affrows <= 0 {
		return nil, errors.New("empty comment")
	}
	if err := qry.Error; err != nil {
		log.Println("error query: ", err.Error())
		return nil, err
	}
	return comments, nil
}
func (cd *commentData) Update(updComment comment.Core) (comment.Core, error) {
	return comment.Core{}, nil
}
func (cd *commentData) Delete(userId, postId, commentId uint) error {
	return nil
}
