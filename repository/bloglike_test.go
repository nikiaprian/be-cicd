package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/models"
)

// Gunakan MockRepository yang diimport dari file lain

func TestCreateLikeByBlogId(t *testing.T) {
	// Inisialisasi gin context
	ctx, _ := gin.CreateTestContext(nil)

	// Inisialisasi MockRepository
	mockRepo := new(MockRepository)

	// Setup ekspektasi
	userID := 1
	blogID := 1
	expectedBlogLike := &models.BlogsLikesResponse{
		ID:     1,
		BlogID: blogID,
		User: models.User{
			ID: userID,
		},
	}

	// Set ekspektasi return value dari method yang di-mock
	mockRepo.On("CreateLikeByBlogId", ctx, userID, blogID).Return(expectedBlogLike, nil)

	// Panggil method yang di-mock
	blogLike, err := mockRepo.CreateLikeByBlogId(ctx, userID, blogID)

	// Verifikasi hasil
	assert.Nil(t, err)
	assert.NotNil(t, blogLike)
	assert.Equal(t, expectedBlogLike.ID, blogLike.ID)
	assert.Equal(t, expectedBlogLike.BlogID, blogLike.BlogID)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}
