package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/models"
)

func TestCreateLikeByBlogCommentId(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	mockRepo.On("CreateLikeByBlogCommentId", &ctx, 1, 1).Return(&models.BlogCommentLikeResponse{ID: 1, BlogCommentID: 1}, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.CreateLikeByBlogCommentId(&ctx, 1, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestGetAllLikeByBlogCommentID(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLikes := []models.BlogCommentLikeResponse{
		{ID: 1, BlogCommentID: 1},
		{ID: 2, BlogCommentID: 1},
	}
	mockRepo.On("GetAllLikeByBlogCommentID", &ctx, 1).Return(&expectedLikes, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.GetAllLikeByBlogCommentID(&ctx, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(*result))

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestDeleteLikeByBlogCommentId(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLike := &models.BlogCommentLikeResponse{ID: 1, BlogCommentID: 1}
	mockRepo.On("DeleteLikeByBlogCommentId", &ctx, 1, 1).Return(expectedLike, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.DeleteLikeByBlogCommentId(&ctx, 1, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}
