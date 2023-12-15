package usecase

import (
	"codein/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) CreateLikeByBlogCommentId(c *gin.Context) (*models.BlogCommentLikesResponse, error) {
	user := c.MustGet("user").(*models.User)
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	blogCommentLike, err := usecase.repository.CreateLikeByBlogCommentId(c, user.ID, convertId)

	if err != nil {
		return nil, err
	}

	blogCommentLike.User = *user

	return blogCommentLike, nil
}

func (usecase *Usecase) DeleteLikeByBlogCommentId(c *gin.Context) (*models.BlogCommentLikesResponse, error) {
	user := c.MustGet("user").(*models.User)
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	blogCommentLike, err := usecase.repository.DeleteLikeByBlogCommentId(c, user.ID, convertId)

	if err != nil {
		return nil, err
	}

	blogCommentLike.User = *user

	return blogCommentLike, nil
}
