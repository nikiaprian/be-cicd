package usecase

import (
	"kel15/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) CreateLikeByForumId(c *gin.Context) (*models.ForumsLikesResponse, error) {
	user := c.MustGet("user").(*models.User)
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	forumLike, err := usecase.repository.CreateLikeByForumId(c, user.ID, convertId)

	if err != nil {
		return nil, err
	}

	forumLike.User = *user

	return forumLike, nil
}

func (usecase *Usecase) DeleteLikeByForumId(c *gin.Context) (*models.ForumsLikesResponse, error) {
	user := c.MustGet("user").(*models.User)
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	forumLike, err := usecase.repository.DeleteLikeByForumId(c, user.ID, convertId)

	if err != nil {
		return nil, err
	}

	forumLike.User = *user

	return forumLike, nil
}
