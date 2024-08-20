package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/models"
	"codein/repository"
)

// Mock untuk Project yang menyertakan Usecase
type MockProject struct {
	Usecase *MockUsecase
	Storage *MockStorage // Jika Anda menggunakan Storage
}

// Mock untuk Usecase
type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) CreateBlog(c *gin.Context) (*models.Blog, error) {
	args := m.Called(c)
	return args.Get(0).(*models.Blog), args.Error(1)
}

func (m *MockUsecase) GetAllBlog(c *gin.Context) (*[]models.Blog, error) {
	args := m.Called(c)
	return args.Get(0).(*[]models.Blog), args.Error(1)
}

// Anda bisa menambahkan lebih banyak metode mock sesuai kebutuhan

func TestCreateBlog(t *testing.T) {
	mockUsecase := new(MockUsecase)
	mockProject := &MockProject{Usecase: mockUsecase}
	handler := &Handler{Project: mockProject}
	gin.SetMode(gin.TestMode)

	// Siapkan request dan recorder
	req := httptest.NewRequest(http.MethodPost, "/blogs", nil) // tambahkan payload yang diperlukan
	w := httptest.NewRecorder()

	// Siapkan ekspektasi pada mock
	expectedBlog := &models.Blog{ID: 1}
	mockUsecase.On("CreateBlog", mock.Anything).Return(expectedBlog, nil)

	// Panggil handler
	handler.CreateBlog(w, req)

	// Periksa hasil
	assert.Equal(t, http.StatusCreated, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestGetAllBlog(t *testing.T) {
	mockUsecase := new(MockUsecase)
	mockProject := &MockProject{Usecase: mockUsecase}
	handler := &Handler{Project: mockProject}
	gin.SetMode(gin.TestMode)

	// Siapkan request dan recorder
	req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
	w := httptest.NewRecorder()

	// Siapkan ekspektasi pada mock
	expectedBlogs := []models.Blog{{ID: 1}, {ID: 2}}
	mockUsecase.On("GetAllBlog", mock.Anything).Return(&expectedBlogs, nil)

	// Panggil handler
	handler.GetAllBlog(w, req)

	// Periksa hasil
	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}
