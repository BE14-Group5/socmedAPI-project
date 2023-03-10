package post

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Content   string
	Photo     string
	UserID    uint
	Writer    string
	CreatedAt time.Time
}

type PostHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	MyPosts() echo.HandlerFunc
	AllPosts() echo.HandlerFunc
	GetPostById() echo.HandlerFunc
}

type PostService interface {
	Add(token interface{}, newPost Core, postPhoto *multipart.FileHeader) (Core, error)
	Update(token interface{}, postID uint, updatedPost Core, updatePhoto *multipart.FileHeader) (Core, error)
	Delete(token interface{}, postID uint) error
	MyPosts(token interface{}) ([]MyPostsResp, error)
	AllPosts() ([]MyPostsResp, error)
	GetPostById(token interface{}, postID uint) (MyPostsResp, error)
}

type PostData interface {
	Add(userID uint, newPost Core) (Core, error)
	Update(postID uint, userID uint, updatedPost Core) (Core, error)
	Delete(postID uint, userID uint) error
	MyPosts(userID uint) ([]MyPostsResp, error)
	AllPosts() ([]MyPostsResp, error)
	GetPostById(postID uint, userID uint) (MyPostsResp, error)
}
