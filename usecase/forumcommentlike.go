package usecase

import (
	"codein/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) CreateLikeByForumCommentId(c *gin.Context) (*models.ForumCommentLikesResponse, error) {
	user := c.MustGet("user").(*models.User)
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	forumCommentLike, err := usecase.repository.CreateLikeByForumCommentId(c, user.ID, convertId)

	if err != nil {
		return nil, err
	}

	forumCommentLike.User = *user

	return forumCommentLike, nil
}

func (usecase *Usecase) DeleteLikeByForumCommentId(c *gin.Context) (*models.ForumCommentLikesResponse, error) {
	user := c.MustGet("user").(*models.User)
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	forumCommentLike, err := usecase.repository.DeleteLikeByForumCommentId(c, user.ID, convertId)

	if err != nil {
		return nil, err
	}

	forumCommentLike.User = *user

	return forumCommentLike, nil
}
