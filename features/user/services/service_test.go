package services

import (
	"errors"
	"mime/multipart"
	"simple-social-media-API/features/user"
	"simple-social-media-API/helper"
	"simple-social-media-API/mocks"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)

	newUser := user.Core{
		Email:           "mfauzanptra@gmail.com",
		Name:            "Muhamad Fauzan Putra",
		PhoneNumber:     "085659171799",
		Password:        "paupau99",
		ProfilePhoto:    "1.jpg",
		BackgroundPhoto: "2.jpg",
	}
	expectedData := user.Core{
		Email:           "mfauzanptra@gmail.com",
		Name:            "Muhamad Fauzan Putra",
		PhoneNumber:     "085659171799",
		Password:        "paupau99",
		ProfilePhoto:    "files/1.jpg",
		BackgroundPhoto: "files/2.jpg",
	}

	var a, b *multipart.FileHeader
	t.Run("success register", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(expectedData, nil).Once()
		srv := New(data)
		res, err := srv.Register(newUser, a, b)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.ID, res.ID)
		assert.Equal(t, expectedData.Name, res.Name)
		data.AssertExpectations(t)
	})

	t.Run("duplicate", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()
		srv := New(data)
		res, err := srv.Register(newUser, a, b)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "user already exist")
		assert.Equal(t, res.Name, "")
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("server error")).Once()
		srv := New(data)
		res, err := srv.Register(newUser, a, b)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.Name, "")
	})

	t.Run("error upload profile photo", func(t *testing.T) {
		a := &multipart.FileHeader{
			Filename: "a",
			Size:     10,
		}
		srv := New(data)
		res, err := srv.Register(newUser, a, b)
		assert.NotNil(t, err)
		assert.Equal(t, res.Name, "")
	})

	t.Run("error upload background photo", func(t *testing.T) {
		b := &multipart.FileHeader{
			Filename: "a",
			Size:     10,
		}
		srv := New(data)
		res, err := srv.Register(newUser, a, b)
		assert.NotNil(t, err)
		assert.Equal(t, res.Name, "")
	})
}

func TestLogn(t *testing.T) {
	data := mocks.NewUserData(t)
	t.Run("succcess login", func(t *testing.T) {
		email := "mfauzanptra@gmail.com"
		password := "paupau99"
		hashed, _ := bcrypt.GenerateFromPassword([]byte("paupau99"), bcrypt.DefaultCost)

		expectedData := user.Core{
			Email:           "mfauzanptra@gmail.com",
			Name:            "Muhamad Fauzan Putra",
			PhoneNumber:     "085659171799",
			Password:        string(hashed),
			ProfilePhoto:    "files/1.jpg",
			BackgroundPhoto: "files/2.jpg",
		}
		data.On("Login", email).Return(expectedData, nil)
		srv := New(data)
		token, res, err := srv.Login(email, password)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.Email, res.Email)
		assert.NotNil(t, token)
		data.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		inputEmail := "mfp@gmail.com"
		data.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found"))

		srv := New(data)
		token, res, err := srv.Login(inputEmail, "be1422")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "mfauzanptra@gmail.com"
		hashed, _ := bcrypt.GenerateFromPassword([]byte("paupau99"), bcrypt.DefaultCost)
		expData := user.Core{
			Email:           "mfauzanptra@gmail.com",
			Name:            "Muhamad Fauzan Putra",
			PhoneNumber:     "085659171799",
			Password:        string(hashed),
			ProfilePhoto:    "files/1.jpg",
			BackgroundPhoto: "files/2.jpg",
		}
		data.On("Login", inputEmail).Return(expData, nil)

		srv := New(data)
		token, res, err := srv.Login(inputEmail, "be1423")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wrong password")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	data := mocks.NewUserData(t)

	t.Run("Sukses lihat profile", func(t *testing.T) {
		resData := user.Core{ID: uint(1), Name: "jerry", Email: "jerry@alterra.id", PhoneNumber: "08123456"}

		data.On("Profile", uint(1)).Return(resData, nil).Once()

		srv := New(data)

		claims := jwt.MapClaims{}
		claims["authorized"] = true
		claims["userID"] = 1
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		token.Valid = true

		res, err := srv.Profile(token)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		res, err := srv.Profile(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Profile", uint(4)).Return(user.Core{}, errors.New("data not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		data.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}
