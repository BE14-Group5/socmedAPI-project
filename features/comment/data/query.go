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

func (cd *commentData) Add(userId uint, newComment comment.Core) (comment.Core, error) {
	cnv := CoreToData(newComment)
	cnv.UserId = userId
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
func (cd *commentData) Update(userId uint, commentId uint, postId uint, updComment comment.Core) (comment.Core, error) {
	updt := CoreToData(updComment)
	qry := cd.db.Where("id = ? AND user_id = ? AND post_id = ?", commentId, userId, postId).Updates(&updt)
	if qry.RowsAffected <= 0 {
		log.Println("update comment query error : data not found")
		return comment.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update comment query error :", err.Error())
		return comment.Core{}, errors.New("not found")
	}

	return ToCore(updt), nil
}
func (cd *commentData) Delete(userId, postId, commentId uint) error {
	qry := cd.db.Where("user_id = ? AND post_id = ?", userId, postId).Delete(&Comment{}, commentId)

	if qry.RowsAffected <= 0 {
		log.Println("no rows affected")
		return errors.New("data not found")
	}

	if err := qry.Error; err != nil {
		log.Println("delete comment query error")
		return errors.New("data not found")
	}
	return nil
}
