package handler

import (
	"fmt"
	"codein/models"
	"codein/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (blogHandler *Handler) GetAllBlog(c *gin.Context) {
	data, err := blogHandler.Project.Usecase.GetAllBlog(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}
func (blogHandler *Handler) CreateBlog(c *gin.Context) {
	user, _ := c.Get("user")
	userData := user.(*models.User)
	file, fileHeader, _, err := utils.GetFileUpload(c, true)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
	}

	uuid_generate := uuid.New()
	tempFileName := fmt.Sprintf("blog/user-%d/%d-%s", userData.ID, userData.ID, uuid_generate)
	fileName, err := utils.UploadToS3(userData.ID, blogHandler.Project.Storage, file, fileHeader, tempFileName)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
	}
	c.Set("file", fileName)
	data, err := blogHandler.Project.Usecase.CreateBlog(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sendResponseSuccess{Success: true, Code: 201, Data: data})
	return
}

// func (blogHandler *Handler) UpdateBlog(c *gin.Context) {
// 	user, _ := c.Get("user")
// 	userData := user.(*models.User)
// 	file, fileHeader, err := utils.GetFileUpload(c)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
// 	}

// 	fileName, err := utils.UploadToS3(userData.ID, blogHandler.Project.Storage, file, fileHeader)
// 	fmt.Println("filename", fileName)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
// 	}
// 	c.Set("file", fileName)

// 	data, err := blogHandler.Project.Usecase.UpdateBlog(c)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
// 	return
// }

func (blogHandler *Handler) DeleteBlog(c *gin.Context) {
	err := blogHandler.Project.Usecase.DeleteBlog(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: nil})
	return
}

func (blogHandler *Handler) GetBlogByID(c *gin.Context) {
	data, err := blogHandler.Project.Usecase.GetBlogByID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}
