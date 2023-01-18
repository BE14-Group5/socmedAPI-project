package handler

import (
	"log"
	"mime/multipart"
	"net/http"
	"simple-social-media-API/features/post"
	"simple-social-media-API/helper"
	"strconv"

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

		if file, err := c.FormFile("photo"); err != nil {
			log.Println("error read post photo")
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
	return func(c echo.Context) error {
		token := c.Get("user")
		var updatePhoto *multipart.FileHeader

		postID := c.Param("id")
		cnv, err := strconv.Atoi(postID)
		if err != nil {
			log.Println("update post param error")
			return c.JSON(http.StatusBadRequest, "id post salah")
		}

		input := UpdatePostRequest{}
		err2 := c.Bind(&input)
		if err2 != nil {
			log.Println("update post body scan error")
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		if file, err := c.FormFile("photo"); err != nil {
			log.Println("error read update post photo")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		} else {
			updatePhoto = file
		}

		res, err := ph.srvc.Update(token, uint(cnv), *ToCore(input), updatePhoto)
		if err != nil {
			log.Println("error running update post service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    UpdatePostToResponse(res),
			"message": "success update post",
		})
	}
}

func (ph *postHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := c.Param("id")
		cnv, err := strconv.Atoi(input)
		if err != nil {
			log.Println("delete post param error")
			return c.JSON(http.StatusBadRequest, "id post salah")
		}

		err2 := ph.srvc.Delete(token, uint(cnv))
		if err2 != nil {
			log.Println("error running update post service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusOK, "success delete post")
	}
}

func (ph *postHandle) MyPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		res, err := ph.srvc.MyPosts(token)
		if err != nil {
			log.Println("error running myposts service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ListMyPostsToResponse(res),
			"message": "success show all my posts post",
		})
	}
}

func (ph *postHandle) AllPosts() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := ph.srvc.AllPosts()
		if err != nil {
			log.Println("error running myposts service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ListAllPostsToResponse(res),
			"message": "success show all users posts",
		})
	}
}

func (ph *postHandle) GetPostById() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := c.Param("id")
		cnv, err := strconv.Atoi(input)
		if err != nil {
			log.Println("GetPostById param error")
			return c.JSON(http.StatusBadRequest, "id post salah")
		}

		res, err := ph.srvc.GetPostById(token, uint(cnv))
		if err != nil {
			log.Println("error running GetPostById service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    GetPostByIdToResponse(res),
			"message": "success get post",
		})
	}
}
