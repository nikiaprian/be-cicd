package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"codein/models"
)

func TestCreateLikeByBlogId(t *testing.T) {
	ctx, _ := gin.CreateTestContext(nil)
	mockRepo := new(MockRepository)

	blogID := 1
	userID := 1
	expectedLike := &models.BlogLike{
		BlogID: blogID,
		UserID: userID,
	}

	mockRepo.On("CreateLikeByBlogId", ctx, blogID, userID).Return(expectedLike, nil)

	like, err := mockRepo.CreateLikeByBlogId(ctx, blogID, userID)

	assert.Nil(t, err)
	assert.NotNil(t, like)
	assert.Equal(t, expectedLike.BlogID, like.BlogID)
	assert.Equal(t, expectedLike.UserID, like.UserID)

	mockRepo.AssertExpectations(t)
}
