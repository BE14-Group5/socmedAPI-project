package handler

import "simple-social-media-API/features/user"

type UserResponse struct {
	ID              uint   `json:"id" form:"id"`
	Email           string `json:"email" form:"email"`
	Name            string `json:"name" form:"name"`
	PhoneNumber     string `json:"phone_number" form:"phone_number"`
	ProfilePhoto    string `json:"profile_photo" form:"profile_photo"`
	BackgroundPhoto string `json:"background_photo" form:"background_photo"`
}

func ToResponse(data user.Core) UserResponse {
	return UserResponse{
		ID:              data.ID,
		Email:           data.Email,
		Name:            data.Name,
		PhoneNumber:     data.PhoneNumber,
		ProfilePhoto:    data.ProfilePhoto,
		BackgroundPhoto: data.BackgroundPhoto,
	}
}
