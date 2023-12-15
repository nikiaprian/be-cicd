package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) CreateLikeByBlogCommentId(c *gin.Context) {
	data, err := handler.Project.Usecase.CreateLikeByBlogCommentId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}

func (handler *Handler) DeleteLikeByBlogCommentId(c *gin.Context) {
	data, err := handler.Project.Usecase.DeleteLikeByBlogCommentId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data, Message: "Data berhasil dihapus"})
	return
}
