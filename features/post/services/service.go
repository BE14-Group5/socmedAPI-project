package services

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
		return post.Core{}, errors.New("user not found")
	}

	if postPhoto != nil {
		path, err := helper.UploadPostPhotoS3(*postPhoto, helper.ExtractToken(token))
		if err != nil {
			return post.Core{}, err
		}
		newPost.Photo = path
	}

	res, err := ps.data.Add(uint(userID), newPost)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "server problem"
		}
		return post.Core{}, errors.New(msg)
	}

	return res, nil
}

func (ps *postSrvc) Update(token interface{}, postID uint, updatedPost post.Core, updatePhoto *multipart.FileHeader) (post.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return post.Core{}, errors.New("user not found")
	}

	// res, err := ps.data.GetPostById(postID, uint(userID))
	// if err != nil {
	// 	msg := ""
	// 	if strings.Contains(err.Error(), "not found") {
	// 		msg = "data not found"
	// 	} else {
	// 		msg = "server problem"
	// 	}
	// 	return post.Core{}, errors.New(msg)
	// }

	if updatePhoto != nil {
		path, _ := helper.UploadPostPhotoS3(*updatePhoto, helper.ExtractToken(token))
		updatedPost.Photo = path
	}

	res, err := ps.data.Update(postID, uint(userID), updatedPost)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "post data not found"
		} else {
			msg = "server problem"
		}
		return post.Core{}, errors.New(msg)
	}

	return res, nil
}

func (ps *postSrvc) Delete(token interface{}, postID uint) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user not found")
	}

	err := ps.data.Delete(postID, uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}
	return nil
}

func (ps *postSrvc) MyPosts(token interface{}) ([]post.MyPostsResp, error) {
	userID := helper.ExtractToken(token)

	res, err := ps.data.MyPosts(uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []post.MyPostsResp{}, errors.New(msg)
	}
	return res, nil
}

func (ps *postSrvc) AllPosts() ([]post.MyPostsResp, error) {
	res, err := ps.data.AllPosts()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []post.MyPostsResp{}, errors.New(msg)
	}
	return res, nil
}

func (ps *postSrvc) GetPostById(token interface{}, postID uint) (post.MyPostsResp, error) {
	userID := helper.ExtractToken(token)
	res, err := ps.data.GetPostById(postID, uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return post.MyPostsResp{}, errors.New(msg)
	}
	return res, nil
}
