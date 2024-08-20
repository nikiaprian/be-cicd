package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/models"
	"codein/repository"
)

type MockBlogRepository struct {
	mock.Mock
}

func (m *MockBlogRepository) CreateBlog(ctx *gin.Context, blogReq models.BlogRequest, photo string, userID int) (*models.Blog, error) {
	args := m.Called(ctx, blogReq, photo, userID)
	return args.Get(0).(*models.Blog), args.Error(1)
}

func (m *MockBlogRepository) GetAllBlog(ctx *gin.Context) (*[]models.Blog, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]models.Blog), args.Error(1)
}

// Tambahkan metode lainnya sesuai kebutuhan pengujian

func TestCreateBlog(t *testing.T) {
	mockRepo := new(MockBlogRepository)
	handler := &BlogHandler{repo: mockRepo}
	gin.SetMode(gin.TestMode)

	// Siapkan request dan recorder
	req := httptest.NewRequest(http.MethodPost, "/blogs", nil) // tambahkan payload jika perlu
	w := httptest.NewRecorder()

	// Siapkan ekspektasi pada mock
	mockRepo.On("CreateBlog", mock.Anything, mock.Anything, "", mock.Anything).Return(&models.Blog{ID: 1}, nil)

	// Panggil handler
	handler.CreateBlog(w, req)

	// Periksa hasil
	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetAllBlog(t *testing.T) {
	mockRepo := new(MockBlogRepository)
	handler := &BlogHandler{repo: mockRepo}
	gin.SetMode(gin.TestMode)

	// Siapkan request dan recorder
	req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
	w := httptest.NewRecorder()

	// Siapkan ekspektasi pada mock
	expectedBlogs := []models.Blog{{ID: 1}, {ID: 2}}
	mockRepo.On("GetAllBlog", mock.Anything).Return(&expectedBlogs, nil)

	// Panggil handler
	handler.GetAllBlog(w, req)

	// Periksa hasil
	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}
