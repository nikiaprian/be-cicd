package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) CreateLikeByForumCommentId(c *gin.Context) {
	data, err := handler.Project.Usecase.CreateLikeByForumCommentId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}

func (handler *Handler) DeleteLikeByForumCommentId(c *gin.Context) {
	data, err := handler.Project.Usecase.DeleteLikeByForumCommentId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data, Message: "Data berhasil dihapus"})
	return
}
