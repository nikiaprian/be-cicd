package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockProject struct {
	// Implement interface methods required for tests
}

func TestCreateBlog(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodPost, "/blogs", nil) // Tambahkan body jika perlu
	c.Request = req

	// Set up mockProject
	mockProject := &MockProject{}

	// Initialize handler
	blogHandler := &Handler{Project: mockProject}

	// Call the function
	blogHandler.CreateBlog(c)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)
}

// Tambahkan test lainnya dengan cara yang sama
