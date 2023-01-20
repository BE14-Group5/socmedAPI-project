package services

import (
	"errors"
	"simple-social-media-API/features/comment"
	"simple-social-media-API/helper"
	"simple-social-media-API/mocks"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewCommentData(t)

	inputData := comment.Core{
		PostId:  1,
		Content: "karya yang luar biasa",
	}
	resData := comment.Core{
		ID:        1,
		UserId:    2,
		UserName:  "",
		PostId:    1,
		CreatedAt: time.Now(),
		Content:   "karya yang luar biasa",
	}

	t.Run("success add comment", func(t *testing.T) {
		repo.On("Add", uint(2), inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		userID := uint(2)

		repo.On("Add", userID, inputData).Return(comment.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(2)

		repo.On("Add", userID, inputData).Return(comment.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.ID, uint(0))
		repo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.Add(pToken, inputData)
		assert.NotNil(t, err)
		// assert.Equal(t, res.UserID, uint(0))
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewCommentData(t)

	inputData := comment.Core{
		PostId:  1,
		Content: "karya yang luar biasa",
	}
	resData := comment.Core{
		ID:        1,
		UserId:    1,
		UserName:  "",
		PostId:    1,
		CreatedAt: time.Now(),
		Content:   "karya yang luar biasa",
	}

	t.Run("success update comment", func(t *testing.T) {
		userID := uint(1)
		postId := uint(1)
		commentId := uint(1)

		repo.On("Update", userID, commentId, postId, inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, commentId, postId, inputData)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, resData.ID)
		repo.AssertExpectations(t)
	})

	t.Run("comment not found", func(t *testing.T) {
		userID := uint(1)
		postId := uint(1)
		commentId := uint(1)
		repo.On("Update", userID, commentId, postId, inputData).Return(comment.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, commentId, postId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(1)
		postId := uint(1)
		commentId := uint(1)
		repo.On("Update", userID, commentId, postId, inputData).Return(comment.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, commentId, postId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("jwt not valid", func(t *testing.T) {
		postId := uint(1)
		commentId := uint(1)
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, commentId, postId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.UserId)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewCommentData(t)

	t.Run("success delete comment", func(t *testing.T) {
		userID := uint(2)
		commentID := uint(1)
		repo.On("Delete", userID, commentID).Return(nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, commentID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		userID := uint(2)
		commentID := uint(1)
		repo.On("Delete", userID, commentID).Return(errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, commentID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(2)
		commentID := uint(1)
		repo.On("Delete", userID, commentID).Return(errors.New("server problem")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, commentID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("jwt not valid", func(t *testing.T) {
		commentID := uint(1)

		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, commentID)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
