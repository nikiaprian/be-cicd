package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBlog(t *testing.T) {
	// Setup gin and handler
	r := gin.Default()
	handler := &Handler{} // Assumes Handler is properly initialized
	r.GET("/blogs", handler.GetAllBlog)

	// Test GetAllBlog (should return 200 OK)
	req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
