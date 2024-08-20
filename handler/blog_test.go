package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"codein/usecase"
	"codein/models"
)

// Mock Usecase
type MockBlogUsecase struct{}

func (m *MockBlogUsecase) GetAllBlog(c *gin.Context) ([]models.Blog, error) {
	return []models.Blog{
		{ID: 1, Title: "Test Blog 1"},
		{ID: 2, Title: "Test Blog 2"},
	}, nil
}

func TestGetAllBlog(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Initialize the handler with the mock usecase
	handler := &Handler{
		Project: &usecase.Project{
			Usecase: &MockBlogUsecase{},
		},
	}

	r.GET("/blogs", handler.GetAllBlog)

	req, _ := http.NewRequest(http.MethodGet, "/blogs", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Blog 1")
	assert.Contains(t, w.Body.String(), "Test Blog 2")
}
