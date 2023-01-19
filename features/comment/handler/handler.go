package handler

import (
	"log"
	"net/http"
	"simple-social-media-API/features/comment"
	"simple-social-media-API/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type commentHandle struct {
	srv comment.CommentService
}

func New(cs comment.CommentService) comment.CommentHandler {
	return &commentHandle{
		srv: cs,
	}
}

func (ch *commentHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddCommentReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}

		cnv := ToCore(input)
		res, err := ch.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			if strings.Contains(err.Error(), "wrong input") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input format"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success comment",
		})
	}
}
func (ch *commentHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		commentId := c.Param("id")
		cnv, err := strconv.Atoi(commentId)
		if err != nil {
			log.Println("update commentID convert error")
			return c.JSON(http.StatusBadRequest, "wrong id comment")
		}

		input := AddCommentReq{}
		err2 := c.Bind(&input)
		if err2 != nil {
			log.Println("update comment body scan error")
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}
		postId := input.PostId
		res, err := ch.srv.Update(token, uint(cnv), postId, *ToCore(input))
		if err != nil {
			log.Println("error running update comment service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success update comment",
		})
	}
}
func (ch *commentHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		commentId := c.Param("id")
		cnv, err := strconv.Atoi(commentId)
		if err != nil {
			log.Println("delete commentId convert error")
			return c.JSON(http.StatusBadRequest, "wrong ID comment")
		}

		input := DeleteCommentReq{}
		err2 := c.Bind(&input)
		if err2 != nil {
			log.Println("delete comment body scan error")
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}

		postID := input.PostId
		error := ch.srv.Delete(token, postID, uint(cnv))
		if error != nil {
			log.Println("error running delete comment service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success delete comment",
		})
	}
}
