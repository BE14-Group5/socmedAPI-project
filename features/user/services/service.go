package services

import (
	"simple-social-media-API/features/user"

	"github.com/go-playground/validator/v10"
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

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {

}
func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {

}
func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {

}
func (uuc *userUseCase) Update(token interface{}, updatedData user.Core) (user.Core, error) {

}
func (uuc *userUseCase) Deactive(token interface{}) error {

}
