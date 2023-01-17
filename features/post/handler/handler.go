package handler

import (
	"log"
	"mime/multipart"
	"net/http"
	"simple-social-media-API/features/post"
	"simple-social-media-API/helper"

	"github.com/labstack/echo/v4"
)

type postHandle struct {
	srvc post.PostService
}

func Isolation(ps post.PostService) post.PostHandler {
	return &postHandle{
		srvc: ps,
	}
}

func (ph *postHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddPostRequest{}
		var postPhoto *multipart.FileHeader
		if err := c.Bind(&input); err != nil {
			log.Println("add post body scan error")
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		if file, err := c.FormFile("image"); err != nil {
			log.Println("error read image")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		} else {
			postPhoto = file
		}

		res, err := ph.srvc.Add(token, *ToCore(input), postPhoto)
		if err != nil {
			log.Println("error running add post service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    AddPostToResponse(res),
			"message": "success posting",
		})
	}
}

func (ph *postHandle) Update() echo.HandlerFunc {
	return nil
}

func (ph *postHandle) Delete() echo.HandlerFunc {
	return nil
}

func (ph *postHandle) MyPosts() echo.HandlerFunc {
	return nil
}

func (ph *postHandle) AllPosts() echo.HandlerFunc {
	return nil
}
