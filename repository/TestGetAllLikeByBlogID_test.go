package repository

import (
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "codein/models"
)

func TestGetAllLikeByBlogID(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLikes := []models.BlogsLikesResponse{
		{ID: 1, BlogID: 1},
		{ID: 2, BlogID: 1},
	}
	mockRepo.On("GetAllLikeByBlogID", &ctx, 1).Return(&expectedLikes, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.GetAllLikeByBlogID(&ctx, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(*result))

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}