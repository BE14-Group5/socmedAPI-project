package handler

import (
	"simple-social-media-API/features/user"
)

type RegisterRequest struct {
	Email           string `json:"email" form:"email"`
	Name            string `json:"name" form:"name"`
	PhoneNumber     string `json:"phone_number" form:"phone_number"`
	Password        string `json:"password" form:"password"`
	ProfilePhoto    string `json:"profile_photo" form:"profile_photo"`
	BackgroundPhoto string `json:"background_photo" form:"background_photo"`
}

type LoginReqest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Email = cnv.Email
		res.Name = cnv.Name
		res.PhoneNumber = cnv.PhoneNumber
		res.Password = cnv.Password
		res.ProfilePhoto = cnv.ProfilePhoto
		res.BackgroundPhoto = cnv.BackgroundPhoto
	default:
		return nil
	}
	return &res
}
