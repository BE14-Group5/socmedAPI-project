package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID              uint
	Email           string
	Name            string
	PhoneNumber     string
	Password        string
	ProfilePhoto    string
	BackgroundPhoto string
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
