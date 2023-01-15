package handler

import (
	"simple-social-media-API/features/user"

	"github.com/labstack/echo/v4"
)

type userControl struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControl{
		srv: srv,
	}
}

func (uc *userControl) Register() echo.HandlerFunc {

}
func (uc *userControl) Login() echo.HandlerFunc {

}
func (uc *userControl) Profile() echo.HandlerFunc {

}
func (uc *userControl) Deactive() echo.HandlerFunc {

}
func (uc *userControl) Update() echo.HandlerFunc {

}
