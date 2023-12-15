package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) CreateCommentForum(c *gin.Context) {
	data, err := handler.Project.Usecase.CreateCommentForum(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}

func (handler *Handler) GetAllCommentByForumID(c *gin.Context) {
	data, err := handler.Project.Usecase.GetAllCommentByForumID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}

func (handler *Handler) DeleteCommentForum(c *gin.Context) {
	err := handler.Project.Usecase.DeleteCommentForum(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: nil})
	return
}

func (handler *Handler) SelectedCommentAnswer(c *gin.Context) {
	comment, err := handler.Project.Usecase.SelectedCommentAnswer(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: comment})
	return
}
