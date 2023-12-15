package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) CreateCommentBlog(c *gin.Context) {
	data, err := handler.Project.Usecase.CreateCommentBlog(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}

func (handler *Handler) GetAllCommentByBlogID(c *gin.Context) {
	data, err := handler.Project.Usecase.GetAllCommentByBlogID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}

func (handler *Handler) DeleteCommentByID(c *gin.Context) {
	err := handler.Project.Usecase.DeleteCommentByID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: nil})
	return
}