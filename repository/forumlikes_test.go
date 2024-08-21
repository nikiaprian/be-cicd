package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"codein/models"
)

func TestCreateLikeByForumId(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	like := models.ForumsLikes{
		ID:    1,
		User:  models.User{ID: 1},
		Forum: models.Forum{ID: 1},
	}
	mockRepo.On("CreateLikeByForumId", like).Return(nil)

	// Panggil metode yang diuji
	err := mockRepo.CreateLikeByForumId(like)

	// Periksa hasil
	assert.Nil(t, err)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestGetAllLikeByForumID(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLikes := []models.ForumsLikesResponse{
		{ID: 1, ForumID: 1},
		{ID: 2, ForumID: 1},
	}
	mockRepo.On("GetAllLikeByForumID", int64(1)).Return(expectedLikes, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.GetAllLikeByForumID(1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestDeleteLikeByForumId(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	mockRepo.On("DeleteLikeByForumId", int64(1)).Return(nil)

	// Panggil metode yang diuji
	err := mockRepo.DeleteLikeByForumId(1)

	// Periksa hasil
	assert.Nil(t, err)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}
