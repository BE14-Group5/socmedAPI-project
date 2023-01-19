package comment

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint      `json:"id" form:"id"`
	UserId    uint      `json:"user_id" form:"user_id"`
	UserName  string    `json:"name" form:"name"`
	PostId    uint      `json:"post_id" form:"post_id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	Content   string    `json:"content" form:"content"`
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	GetComments() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CommentService interface {
	Add(token interface{}, newComment Core) (Core, error)
	GetComments(postId uint) ([]Core, error)
	Update(token interface{}, commentId uint, postId uint, updComment Core) (Core, error)
	Delete(token interface{}, postId, commentId uint) error
}

type CommentData interface {
	Add(userId uint, newComment Core) (Core, error)
	GetComments(postId uint) ([]Core, error)
	Update(userId uint, commentId uint, postId uint, updComment Core) (Core, error)
	Delete(userId, postId, commentId uint) error
}
