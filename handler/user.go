package handler

import (
	"fmt"
	"kel15/models"
	"kel15/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *Handler) UserList(c *gin.Context) {
	utils.SetPaginationDefault(c)

	data, err := handler.Project.Usecase.UserList(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	pagination := utils.GetPagination(c)

	c.JSON(200, sendResponseSuccess{Success: true, Code: 200, Message: "", Data: data, Pagination: &pagination})
	return
}

func (handler *Handler) GetUserProfile(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	c.JSON(http.StatusCreated, sendResponseSuccess{Success: true, Code: http.StatusCreated, Message: "Success created Account", Data: user})
	return

}

func (handler *Handler) UserLogin(c *gin.Context) {
	data, err := handler.Project.Usecase.UserLogin(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(200, sendResponseSuccess{Success: true, Code: 200, Message: "", Data: data})
	return

}

func (handler *Handler) UserRegister(c *gin.Context) {
	data, err := handler.Project.Usecase.UserRegister(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sendResponseSuccess{Success: true, Code: http.StatusCreated, Message: data.Message})
	return
}

func (handler *Handler) UserLoginByProvider(c *gin.Context) {
	url, err := handler.Project.Usecase.UserLoginByProvider(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
	return
}

func (handler *Handler) UserLoginByProviderCallback(c *gin.Context) {
	data, err := handler.Project.Usecase.UserLoginByProviderCallback(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sendResponseSuccess{Success: true, Code: http.StatusCreated, Message: "Success created Account", Data: data})
	return

}

func (handler *Handler) UserProfileUpdate(c *gin.Context) {
	userLogin := c.MustGet("user").(*models.User)
	userByUsername, err := handler.Project.Usecase.GetUserByUsername(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	if userByUsername != nil && userLogin.ID != userByUsername.ID {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: "Username sudah digunakan"})
		return
	}

	var fileName string
	file, fileHeader, isRequired, err := utils.GetFileUpload(c, false)

	if err != nil && isRequired == nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	if err == nil {
		uuid_generate := uuid.New()
		tempFileName := fmt.Sprintf("user/user-%d/profile/%d-%s", userLogin.ID, userLogin.ID, uuid_generate)
		fileName, err = utils.UploadToS3(userLogin.ID, handler.Project.Storage, file, fileHeader, tempFileName)
		if err != nil {
			c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
			return
		}
	}

	data, err := handler.Project.Usecase.UserProfileUpdate(c, fileName)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: http.StatusOK, Message: "", Data: data})
}
