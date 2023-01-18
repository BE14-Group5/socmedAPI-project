package services

import (
	"errors"
	"log"
	"simple-social-media-API/features/comment"
	"simple-social-media-API/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type commentSrv struct {
	qry comment.CommentData
	vld *validator.Validate
}

func New(cd comment.CommentData) comment.CommentService {
	return &commentSrv{
		qry: cd,
		vld: validator.New(),
	}
}

func (cs *commentSrv) Add(token interface{}, newComment comment.Core) (comment.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("error extract token add comment")
		return comment.Core{}, errors.New("user not found")
	}
	newComment.UserId = uint(userId)

	res, err := cs.qry.Add(newComment)
	if err != nil {
		errmsg := ""
		if strings.Contains(err.Error(), "not found") {
			errmsg = "comment not found"
		} else {
			errmsg = "server problem"
		}
		return comment.Core{}, errors.New(errmsg)
	}
	return res, nil
}
func (cs *commentSrv) GetComments() ([]comment.Core, error) {
	return []comment.Core{}, nil
}
func (cs *commentSrv) Update(token interface{}, updComment comment.Core, postId, commentId uint) (comment.Core, error) {
	return comment.Core{}, nil
}
func (cs *commentSrv) Delete(token interface{}, postId, commentId uint) error {
	return nil
}
