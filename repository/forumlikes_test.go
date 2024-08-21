package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/models"
)

func TestCreateLikeByForumId(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	mockRepo.On("CreateLikeByForumId", &ctx, 1, 1).Return(&models.ForumLikesResponse{ID: 1, ForumID: 1}, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.CreateLikeByForumId(&ctx, 1, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestGetAllLikeByForumID(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLikes := []models.ForumLikesResponse{
		{ID: 1, ForumID: 1},
		{ID: 2, ForumID: 1},
	}
	mockRepo.On("GetAllLikeByForumID", &ctx, 1).Return(&expectedLikes, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.GetAllLikeByForumID(&ctx, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(*result))

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestDeleteLikeByForumId(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLike := &models.ForumLikesResponse{ID: 1, ForumID: 1}
	mockRepo.On("DeleteLikeByForumId", &ctx, 1, 1).Return(expectedLike, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.DeleteLikeByForumId(&ctx, 1, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}
