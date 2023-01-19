package services

import (
	"errors"
	"mime/multipart"
	"simple-social-media-API/features/post"
	"simple-social-media-API/helper"
	"simple-social-media-API/mocks"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewPostData(t)

	inputData := post.Core{
		Content: "pemandangan yang indah",
		Photo:   "mountain.jpg",
	}
	resData := post.Core{
		ID:      1,
		Content: "pemandangan yang indah",
		Photo:   "https://socmedapibucket.s3.ap-southeast-1.amazonaws.com/files/post/1/post-photo.jpeg",
	}

	var postPhoto *multipart.FileHeader

	t.Run("success add post", func(t *testing.T) {

		repo.On("Add", uint(1), mock.Anything).Return(resData, nil).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, postPhoto)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("error upload post photo", func(t *testing.T) {
		postPhoto := &multipart.FileHeader{
			Filename: "a",
			Size:     10,
		}
		srv := Isolation(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, postPhoto)
		assert.NotNil(t, err)
		assert.Equal(t, res.Photo, "")
	})

	t.Run("post not found", func(t *testing.T) {
		userID := uint(2)

		repo.On("Add", userID, inputData).Return(post.Core{}, errors.New("data not found")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, postPhoto)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(2)

		repo.On("Add", userID, inputData).Return(post.Core{}, errors.New("server problem")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, postPhoto)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.UserID, uint(0))
		repo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		srv := Isolation(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, postPhoto)
		assert.NotNil(t, err)
		assert.Equal(t, res.UserID, uint(0))
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewPostData(t)

	inputData := post.Core{
		Content: "pemandangan yang indah",
		Photo:   "mountain.jpg",
	}
	resData := post.Core{
		ID:      1,
		Content: "pemandangan yang indah",
		Photo:   "https://socmedapibucket.s3.ap-southeast-1.amazonaws.com/files/post/1/post-photo.jpeg",
	}

	var postPhoto *multipart.FileHeader

	t.Run("success update post", func(t *testing.T) {
		userID := uint(1)
		postId := uint(1)

		repo.On("Update", postId, userID, inputData).Return(resData, nil).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, postId, inputData, postPhoto)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, resData.ID)
		repo.AssertExpectations(t)
	})

	// t.Run("error upload post photo", func(t *testing.T) {
	// 	postID := uint(1)
	// 	userID := uint(1)
	// 	postPhoto := &multipart.FileHeader{
	// 		Filename: "a",
	// 		Size:     10,
	// 	}
	// 	srv := Isolation(repo)

	// 	_, token := helper.GenerateJWT(1)

	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true

	// 	res, err := srv.Update(pToken, postID, inputData, postPhoto)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, res.UserID, userID)
	// })

	t.Run("post not found", func(t *testing.T) {
		userID := uint(1)
		postId := uint(1)

		repo.On("Update", postId, userID, inputData).Return(post.Core{}, errors.New("data not found")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, postId, inputData, postPhoto)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(2)
		postId := uint(3)

		repo.On("Update", postId, userID, inputData).Return(post.Core{}, errors.New("server problem")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, postId, inputData, postPhoto)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("jwt not valid", func(t *testing.T) {
		postId := uint(2)

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, postId, inputData, postPhoto)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.UserID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewPostData(t)

	t.Run("success delete post", func(t *testing.T) {
		postID := uint(3)
		userID := uint(2)
		repo.On("Delete", postID, userID).Return(nil).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, postID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		postID := uint(3)
		userID := uint(2)
		repo.On("Delete", postID, userID).Return(errors.New("data not found")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, postID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		postID := uint(3)
		userID := uint(2)
		repo.On("Delete", postID, userID).Return(errors.New("server problem")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, postID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("jwt not valid", func(t *testing.T) {
		postID := uint(3)

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, postID)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestMyPosts(t *testing.T) {
	repo := mocks.NewPostData(t)

	resData := []post.MyPostsResp{
		{
			ID:        1,
			Content:   "what a wonderful world",
			Photo:     "https://socmedapibucket.s3.ap-southeast-1.amazonaws.com/files/post/1/post-photo.jpeg",
			UserID:    1,
			Writer:    "eka cahya",
			CreatedAt: time.Now(),
		},
	}
	t.Run("success access all my posts", func(t *testing.T) {
		userID := uint(2)
		repo.On("MyPosts", userID).Return(resData, nil).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.MyPosts(pToken)
		assert.Nil(t, err)
		assert.Equal(t, len(res), len(resData))
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		userID := uint(2)
		repo.On("MyPosts", userID).Return([]post.MyPostsResp{}, errors.New("data not found")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.MyPosts(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		// assert.Equal(t, len(res), int(0))
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(2)
		repo.On("MyPosts", userID).Return([]post.MyPostsResp{}, errors.New("server problem")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.MyPosts(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		// assert.Equal(t, len(res), int(0))
		repo.AssertExpectations(t)
	})
}

func TestAllPosts(t *testing.T) {
	repo := mocks.NewPostData(t)

	resData := []post.MyPostsResp{
		{
			ID:        1,
			Content:   "what a wonderful world",
			Photo:     "https://socmedapibucket.s3.ap-southeast-1.amazonaws.com/files/post/1/post-photo.jpeg",
			UserID:    1,
			Writer:    "eka cahya",
			CreatedAt: time.Now(),
		}, {
			ID:        2,
			Content:   "what a beautiful flower",
			Photo:     "https://socmedapibucket.s3.ap-southeast-1.amazonaws.com/files/post/1/post-photo2.jpeg",
			UserID:    2,
			Writer:    "Tony",
			CreatedAt: time.Now(),
		},
	}
	t.Run("success see all posts", func(t *testing.T) {
		repo.On("AllPosts").Return(resData, nil).Once()

		srv := Isolation(repo)

		res, err := srv.AllPosts()
		assert.Nil(t, err)
		assert.Equal(t, len(res), len(resData))
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("AllPosts").Return([]post.MyPostsResp{}, errors.New("data not found")).Once()

		srv := Isolation(repo)

		res, err := srv.AllPosts()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("AllPosts").Return([]post.MyPostsResp{}, errors.New("server problem")).Once()

		srv := Isolation(repo)

		res, err := srv.AllPosts()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})
}

func TestGetPostById(t *testing.T) {
	repo := mocks.NewPostData(t)

	resData := post.MyPostsResp{
		ID:        1,
		Content:   "what a wonderful world",
		Photo:     "https://socmedapibucket.s3.ap-southeast-1.amazonaws.com/files/post/1/post-photo.jpeg",
		UserID:    1,
		Writer:    "eka cahya",
		CreatedAt: time.Now(),
	}

	t.Run("success get post by ID", func(t *testing.T) {
		postID := uint(1)
		userID := uint(1)
		repo.On("GetPostById", postID, userID).Return(resData, nil).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.GetPostById(pToken, postID)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		postID := uint(1)
		userID := uint(1)
		repo.On("GetPostById", postID, userID).Return(post.MyPostsResp{}, errors.New("data not found")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.GetPostById(pToken, postID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		postID := uint(1)
		userID := uint(1)
		repo.On("GetPostById", postID, userID).Return(post.MyPostsResp{}, errors.New("server problem")).Once()

		srv := Isolation(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.GetPostById(pToken, postID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}
