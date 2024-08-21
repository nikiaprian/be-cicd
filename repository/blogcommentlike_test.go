package repository

import (
	"codein/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateLikeByBlogCommentId(t *testing.T) {
	mockRepo := new(MockRepository)
	like := &models.BlogCommentLikes{
		ID: 1,
		User: models.User{
			ID: 1,
		},
		CommentBlog: models.CommentBlog{
			ID: 1,
		},
		CreatedAt: new(time.Time),
		UpdatedAt: new(time.Time),
	}

	mockRepo.On("CreateLikeByBlogCommentId", like).Return(nil)

	err := mockRepo.CreateLikeByBlogCommentId(like)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllLikeByBlogCommentID(t *testing.T) {
	mockRepo := new(MockRepository)
	blogCommentID := 1
	likes := []models.BlogCommentLikesResponse{
		{
			ID:            1,
			BlogCommentID: 1,
			User: models.User{
				ID: 1,
			},
			CreatedAt: new(time.Time),
			UpdatedAt: new(time.Time),
		},
	}

	mockRepo.On("GetAllLikeByBlogCommentID", blogCommentID).Return(likes, nil)

	result, err := mockRepo.GetAllLikeByBlogCommentID(blogCommentID)
	assert.NoError(t, err)
	assert.Equal(t, likes, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteLikeByBlogCommentId(t *testing.T) {
	mockRepo := new(MockRepository)
	blogCommentID := 1
	userID := 1

	mockRepo.On("DeleteLikeByBlogCommentId", blogCommentID, userID).Return(nil)

	err := mockRepo.DeleteLikeByBlogCommentId(blogCommentID, userID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
