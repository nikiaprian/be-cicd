package repository

import (
	"codein/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateLikeByBlogCommentId(like *models.BlogCommentLikes) error {
	args := m.Called(like)
	return args.Error(0)
}

func (m *MockRepository) GetAllLikeByBlogCommentID(blogCommentID int) ([]models.BlogCommentLikesResponse, error) {
	args := m.Called(blogCommentID)
	return args.Get(0).([]models.BlogCommentLikesResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByBlogCommentId(blogCommentID int, userID int) error {
	args := m.Called(blogCommentID, userID)
	return args.Error(0)
}

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
