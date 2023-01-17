package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID              uint   `json:"id" form:"id"`
	Email           string `json:"email" form:"email"`
	Name            string `json:"name" form:"name"`
	PhoneNumber     string `json:"phone_number" form:"phone_number"`
	Password        string `json:"password" form:"password"`
	ProfilePhoto    string `json:"profile_photo" form:"profile_photo"`
	BackgroundPhoto string `json:"background_photo" form:"background_photo"`
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Deactive() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core, profilePhoto *multipart.FileHeader, backgroundPhoto *multipart.FileHeader) (Core, error)
	Login(email, password string) (string, Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, updateData Core) (Core, error)
	Deactive(token interface{}) error
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(email string) (Core, error)
	Profile(id uint) (Core, error)
	Update(id uint, updateData Core) (Core, error)
	Deactive(id uint) error
}
