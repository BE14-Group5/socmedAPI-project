package data

import (
	"errors"
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
	existed := 0
	uq.db.Raw("SELECT COUNT(*) FROM users WHERE deleted_at IS NULL AND email = ?", newUser.Email).Scan(&existed)
	if existed >= 1 {
		return user.Core{}, errors.New("user already exist (duplicated)")
	}
	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		return user.Core{}, err
	}

	newUser.ID = cnv.ID

	return newUser, nil
}
func (uq *userQuery) Login(email string) (user.Core, error) {
	return user.Core{}, nil
}
func (uq *userQuery) Profile(id uint) (user.Core, error) {
	return user.Core{}, nil
}
func (uq *userQuery) Update(id uint, updatedData user.Core) (user.Core, error) {
	return user.Core{}, nil
}
func (uq *userQuery) Deactive(id uint) error {
	return nil
}
