package repository

import (
	"codein/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateLikeByForumCommentId(t *testing.T) {
	mockRepo := new(MockRepository)
	like := &models.ForumCommentLikes{
		ID: 1,
		User: models.User{
			ID: 1,
		},
		CommentForum: models.CommentForum{
			ID: 1,
		},
		CreatedAt: new(time.Time),
		UpdatedAt: new(time.Time),
	}

	mockRepo.On("CreateLikeByForumCommentId", like).Return(nil)

	err := mockRepo.CreateLikeByForumCommentId(like)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllLikeByForumCommentID(t *testing.T) {
	mockRepo := new(MockRepository)
	forumCommentID := 1
	likes := []models.ForumCommentLikesResponse{
		{
			ID:             1,
			ForumCommentID: 1,
			User: models.User{
				ID: 1,
			},
			CreatedAt: new(time.Time),
			UpdatedAt: new(time.Time),
		},
	}

	mockRepo.On("GetAllLikeByForumCommentID", forumCommentID).Return(likes, nil)

	result, err := mockRepo.GetAllLikeByForumCommentID(forumCommentID)
	assert.NoError(t, err)
	assert.Equal(t, likes, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteLikeByForumCommentId(t *testing.T) {
	mockRepo := new(MockRepository)
	forumCommentID := 1
	userID := 1

	mockRepo.On("DeleteLikeByForumCommentId", forumCommentID, userID).Return(nil)

	err := mockRepo.DeleteLikeByForumCommentId(forumCommentID, userID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
