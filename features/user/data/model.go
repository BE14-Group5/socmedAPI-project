package data

import (
	"simple-social-media-API/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string
	Name            string
	PhoneNumber     string
	Password        string
	ProfilePhoto    string
	BackgroundPhoto string
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:              data.ID,
		Email:           data.Email,
		Name:            data.Name,
		PhoneNumber:     data.PhoneNumber,
		Password:        data.Password,
		ProfilePhoto:    data.ProfilePhoto,
		BackgroundPhoto: data.BackgroundPhoto,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model:           gorm.Model{ID: data.ID},
		Email:           data.Email,
		Name:            data.Name,
		PhoneNumber:     data.PhoneNumber,
		Password:        data.Password,
		ProfilePhoto:    data.ProfilePhoto,
		BackgroundPhoto: data.BackgroundPhoto,
	}
}
