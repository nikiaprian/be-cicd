package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/models"
)

func TestCreateLikeByForumCommentId(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	mockRepo.On("CreateLikeByForumCommentId", &ctx, 1, 1).Return(&models.ForumCommentLikeResponse{ID: 1, ForumCommentID: 1}, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.CreateLikeByForumCommentId(&ctx, 1, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestGetAllLikeByForumCommentID(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLikes := []models.ForumCommentLikeResponse{
		{ID: 1, ForumCommentID: 1},
		{ID: 2, ForumCommentID: 1},
	}
	mockRepo.On("GetAllLikeByForumCommentID", &ctx, 1).Return(&expectedLikes, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.GetAllLikeByForumCommentID(&ctx, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(*result))

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestDeleteLikeByForumCommentId(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLike := &models.ForumCommentLikeResponse{ID: 1, ForumCommentID: 1}
	mockRepo.On("DeleteLikeByForumCommentId", &ctx, 1, 1).Return(expectedLike, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.DeleteLikeByForumCommentId(&ctx, 1, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}
