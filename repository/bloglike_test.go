package repository

import (
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "codein/models"
)

func TestCreateLikeByBlogId(t *testing.T) {
    mockRepo := new(MockRepository)
    ctx := gin.Context{}

    // Siapkan ekspektasi pada mock
    mockRepo.On("CreateLikeByBlogId", &ctx, 1, 1).Return(&models.BlogsLikesResponse{ID: 1, BlogID: 1}, nil)

    // Panggil metode yang diuji
    result, err := mockRepo.CreateLikeByBlogId(&ctx, 1, 1)

    // Periksa hasil
    assert.Nil(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, 1, result.ID)

    // Verifikasi bahwa ekspektasi terpenuhi
    mockRepo.AssertExpectations(t)
}
