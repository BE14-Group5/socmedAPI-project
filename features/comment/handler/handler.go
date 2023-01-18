package handler

import (
	"net/http"
	"simple-social-media-API/features/comment"
	"simple-social-media-API/helper"
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
func (ch *commentHandle) GetComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddCommentReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}
		res, err := ch.srv.GetComments(input.PostId)
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("no comment found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success get comments",
		})
	}
}
func (ch *commentHandle) Update() echo.HandlerFunc {
	return nil
}
func (ch *commentHandle) Delete() echo.HandlerFunc {
	return nil
}
