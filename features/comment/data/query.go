package data

// import (
// 	"simple-social-media-API/features/comment"

// 	"gorm.io/gorm"
// )

// type commentQuery struct {
// 	db *gorm.DB
// }

// func New(db *gorm.DB) comment.CommentData {
// 	return comment.CommentData{
// 		db: db,
// 	}
// }

// func (cq *commentQuery) Add(newComment comment.Core) (comment.Core, error) {
// 	return comment.Core{}, nil
// }
// func (cq *commentQuery) GetComments() ([]comment.Core, error) {
// 	return []comment.Core{}, nil
// }
// func (cq *commentQuery) Update(updComment comment.Core) (comment.Core, error) {
// 	return comment.Core{}, nil
// }
// func (cq *commentQuery) Delete(userId, postId, commentId uint) error {
// 	return nil
// }
