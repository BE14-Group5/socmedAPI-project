package data

import (
	"errors"
	"log"
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
		log.Println("user already exist (duplicated)")
		return user.Core{}, errors.New("user already exist (duplicated)")
	}
	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("error create query: ", err.Error())
		return user.Core{}, err
	}

	newUser.ID = cnv.ID

	return newUser, nil
}
func (uq *userQuery) Login(email string) (user.Core, error) {
	res := User{}

	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error: ", err.Error())
		return user.Core{}, errors.New("user not found")
	}
	return ToCore(res), nil
}
func (uq *userQuery) Profile(id uint) (user.Core, error) {
	res := User{}
	if err := uq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return user.Core{}, err
	}

	return ToCore(res), nil
}
func (uq *userQuery) Update(id uint, updatedData user.Core) (user.Core, error) {
	cnvUpd := CoreToData(updatedData)
	qry := uq.db.Model(&User{}).Where("id = ?", id).Updates(cnvUpd)
	if err := qry.Error; err != nil {
		log.Println("error update user query : ", err)
		return updatedData, err
	}
	return updatedData, nil
}
func (uq *userQuery) Deactive(id uint) error {
	user := User{}
	qry := uq.db.Where("id = ?", id).Delete(&user)
	if affrows := qry.RowsAffected; affrows <= 0 {
		return errors.New("user doesn't exist")
	}
	if err := qry.Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
