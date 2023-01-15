package data

import (
	"simple-social-media-API/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {

}
func (uq *userQuery) Login(email string) (user.Core, error) {

}
func (uq *userQuery) Profile(id uint) (user.Core, error) {

}
func (uq *userQuery) Update(id uint, updatedData user.Core) (user.Core, error) {

}
func (uq *userQuery) Deactive(id uint) error {

}
