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

	res, err := cs.qry.Add(uint(userId), newComment)
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
func (cs *commentSrv) Update(token interface{}, commentId uint, postId uint, updComment comment.Core) (comment.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		return comment.Core{}, errors.New("user not found")
	}

	res, err := cs.qry.Update(uint(userId), commentId, postId, updComment)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "comment data not found"
		} else {
			msg = "server problem"
		}
		return comment.Core{}, errors.New(msg)
	}

	return res, nil
}
func (cs *commentSrv) Delete(token interface{}, postId, commentId uint) error {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		return errors.New("user not found")
	}

	if err := cs.qry.Delete(uint(userId), postId, commentId); err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "comment not found"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}
	return nil
}
