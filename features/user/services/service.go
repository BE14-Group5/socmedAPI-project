package services

import (
	"errors"
	"mime/multipart"
	"simple-social-media-API/features/user"
	"simple-social-media-API/helper"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	qry user.UserData
	vld *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
		vld: validator.New(),
	}
}

func (uuc *userUseCase) Register(newUser user.Core, profilePhoto *multipart.FileHeader, backgroundPhoto *multipart.FileHeader) (user.Core, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashed)

	if profilePhoto != nil {
		path, err := helper.UploadProfilePhotoS3(*profilePhoto, newUser.Email)
		if err != nil {
			return user.Core{}, err
		}
		newUser.ProfilePhoto = path
	}

	if backgroundPhoto != nil {
		path, err := helper.UploadBackgroundPhotoS3(*backgroundPhoto, newUser.Email)
		if err != nil {
			return user.Core{}, err
		}
		newUser.BackgroundPhoto = path
	}

	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "user already exist"
		} else {
			msg = "server problem"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}
func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {
	return "", user.Core{}, nil
}
func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	return user.Core{}, nil
}
func (uuc *userUseCase) Update(token interface{}, updatedData user.Core) (user.Core, error) {
	return user.Core{}, nil
}
func (uuc *userUseCase) Deactive(token interface{}) error {
	return nil
}
