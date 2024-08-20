package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/models"
)

// MockRepository is a mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// Mock method for CreateBlog
func (m *MockRepository) CreateBlog(ctx *gin.Context, blogReq models.BlogRequest, photo string, userID int) (*models.Blog, error) {
	args := m.Called(ctx, blogReq, photo, userID)
	return args.Get(0).(*models.Blog), args.Error(1)
}

func TestCreateBlog(t *testing.T) {
	// Inisialisasi gin context
	ctx, _ := gin.CreateTestContext(nil)

	// Inisialisasi MockRepository
	mockRepo := new(MockRepository)

	// Setup ekspektasi
	blogReq := models.BlogRequest{
		Title:   "Test Blog",
		Content: "This is a test blog content.",
	}
	userID := 1
	expectedBlog := &models.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog content.",
	}

	// Set ekspektasi return value dari method yang di-mock
	mockRepo.On("CreateBlog", ctx, blogReq, "Test Photo", userID).Return(expectedBlog, nil)

	// Panggil method yang di-mock
	blog, err := mockRepo.CreateBlog(ctx, blogReq, "Test Photo", userID)

	// Verifikasi hasil
	assert.Nil(t, err)
	assert.NotNil(t, blog)
	assert.Equal(t, expectedBlog.Title, blog.Title)
	assert.Equal(t, expectedBlog.Content, blog.Content)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}
