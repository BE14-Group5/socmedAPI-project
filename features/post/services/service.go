package service

import (
	"errors"
	"mime/multipart"
	"simple-social-media-API/features/post"
	"simple-social-media-API/helper"
	"strings"
)

type postSrvc struct {
	data post.PostData
}

func Isolation(d post.PostData) post.PostService {
	return &postSrvc{
		data: d,
	}
}

func (ps *postSrvc) Add(token interface{}, newPost post.Core, postPhoto *multipart.FileHeader) (post.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return post.Core{}, errors.New("user tidak ditemukan")
	}

	if postPhoto != nil {
		path, err := helper.UploadPostPhotoS3(*postPhoto, helper.ExtractToken(token))
		if err != nil {
			return post.Core{}, err
		}
		newPost.Image = path
	}

	res, err := ps.data.Add(uint(userID), newPost)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "postingan tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return post.Core{}, errors.New(msg)
	}

	return res, nil
}

func (ps *postSrvc) Update(token interface{}, postID uint, updatedPost post.Core) (post.Core, error) {
	return post.Core{}, nil
}

func (ps *postSrvc) Delete(token interface{}, postID uint) error {
	return nil
}

func (ps *postSrvc) MyPosts(token interface{}) ([]post.Core, error) {
	return []post.Core{}, nil
}

func (ps *postSrvc) AllPosts() ([]post.Core, error) {
	return []post.Core{}, nil
}
