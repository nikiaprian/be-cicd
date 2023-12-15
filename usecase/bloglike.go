package usecase

import (
	"codein/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) CreateLikeByBlogId(c *gin.Context) (*models.BlogsLikesResponse, error) {
	user := c.MustGet("user").(*models.User)
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	blogLike, err := usecase.repository.CreateLikeByBlogId(c, user.ID, convertId)

	if err != nil {
		return nil, err
	}

	blogLike.User = *user

	return blogLike, nil
}

func (usecase *Usecase) DeleteLikeByBlogId(c *gin.Context) (*models.BlogsLikesResponse, error) {
	user := c.MustGet("user").(*models.User)
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	blogLike, err := usecase.repository.DeleteLikeByBlogId(c, user.ID, convertId)

	if err != nil {
		return nil, err
	}

	blogLike.User = *user

	return blogLike, nil
}
