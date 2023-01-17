package services

import (
	"errors"
	"log"
	"mime/multipart"
	"simple-social-media-API/config"
	"simple-social-media-API/features/user"
	"simple-social-media-API/helper"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
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
	res, err := uuc.qry.Login(email)

	if err != nil {
		errmsg := ""
		if strings.Contains(err.Error(), "not found") {
			errmsg = err.Error()
		} else {
			errmsg = "server problem"
		}
		return "", user.Core{}, errors.New(errmsg)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil {
		log.Println("wrong password ", err.Error())
		return "", user.Core{}, errors.New("wrong password")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWT_KEY))

	return useToken, res, nil
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
