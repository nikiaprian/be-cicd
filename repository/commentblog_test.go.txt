package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/models"
)

// MockRepository for CommentBlog
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateCommentBlog(ctx *gin.Context, comment models.CommentBlogRequest, userID int) (*models.CommentBlog, error) {
	args := m.Called(ctx, comment, userID)
	return args.Get(0).(*models.CommentBlog), args.Error(1)
}

func TestCreateCommentBlog(t *testing.T) {
	ctx, _ := gin.CreateTestContext(nil)
	mockRepo := new(MockRepository)

	commentReq := models.CommentBlogRequest{
		BlogID:  1,
		Content: "This is a comment",
	}
	userID := 1
	expectedComment := &models.CommentBlog{
		BlogID:  1,
		UserID:  userID,
		Content: "This is a comment",
	}

	mockRepo.On("CreateCommentBlog", ctx, commentReq, userID).Return(expectedComment, nil)

	comment, err := mockRepo.CreateCommentBlog(ctx, commentReq, userID)

	assert.Nil(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, expectedComment.BlogID, comment.BlogID)
	assert.Equal(t, expectedComment.Content, comment.Content)

	mockRepo.AssertExpectations(t)
}
