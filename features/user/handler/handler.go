package handler

import (
	"log"
	"mime/multipart"
	"net/http"
	"simple-social-media-API/features/user"
	"simple-social-media-API/helper"
	"strings"

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
	return func(c echo.Context) error {
		input := RegisterRequest{}
		var profilePhoto, backgroundPhoto *multipart.FileHeader

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		if file, err := c.FormFile("profile_photo"); err != nil {
			log.Println("error read profile_photo")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		} else {
			profilePhoto = file
		}

		if file, err := c.FormFile("background_photo"); err != nil {
			log.Println("error read background_photo")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		} else {
			backgroundPhoto = file
		}

		// if file, err := c.FormFile("profile_photo"); err != nil {
		// 	log.Println("error read profile_photo")
		// 	return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		// } else {
		// 	dir, err := helper.UploadProfilePhoto(*file, input.Email)
		// 	if err != nil {
		// 		log.Println("error running UploadProfilePhoto")
		// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		// 	}
		// 	input.ProfilePhoto = dir
		// }

		// if file, err := c.FormFile("background_photo"); err != nil {
		// 	log.Println("error read background_photo ")
		// 	return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		// } else {
		// 	dir, err := helper.UploadBackgroundPhoto(*file, input.Email)
		// 	if err != nil {
		// 		log.Println("error running UploadBackgroundPhoto")
		// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		// 	}
		// 	input.BackgroundPhoto = dir
		// }

		res, err := uc.srv.Register(*ToCore(input), profilePhoto, backgroundPhoto)
		if err != nil {
			if strings.Contains(err.Error(), "exist") {
				log.Println("error running register service: user already exist")
				return c.JSON(http.StatusConflict, helper.ErrorResponse("user or email already exist"))
			} else {
				log.Println("error running register service")
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success register",
		})
	}
}
func (uc *userControl) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginReqest{}
		if err := c.Bind(&input); err != nil {
			log.Println("error login request: ", err.Error())
			return c.JSON(http.StatusBadRequest, "wrong input")
		}

		token, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			if strings.Contains(err.Error(), "password") {
				log.Println("wrong password: ", err.Error())
				return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("wrong password"))
			} else if strings.Contains(err.Error(), "not found") {
				log.Println("user not found: ", err.Error())
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("user not found"))
			} else {
				log.Println("error login service: ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToLoginResp(res, token),
			"message": "login success",
		})
	}
}
func (uc *userControl) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("user not found: ", err.Error())
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("user not found"))
			} else {
				log.Println("error profile service: ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToResponse(res),
			"message": "get profile success",
		})
	}
}
func (uc *userControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		updatedData := RegisterRequest{}
		if err := c.Bind(&updatedData); err != nil {
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}
		res, err := uc.srv.Update(token, *ToCore(updatedData))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("user not found: ", err.Error())
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("user not found"))
			} else {
				log.Println("error update service: ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success update user's data",
		})
	}
}
func (uc *userControl) Deactive() echo.HandlerFunc {
	return nil
}
