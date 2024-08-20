package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"codein/usecase" // import your usecase package
	"codein/project" // import your project package
)

func TestGetAllBlog(t *testing.T) {
	// Setup gin and handler
	r := gin.Default()
	usecase := &usecase.BlogUsecase{} // Assuming you have a valid BlogUsecase
	project := &project.Project{Usecase: usecase}
	handler := &Handler{Project: project} // Initialize Handler with Project

	r.GET("/blogs", handler.GetAllBlog)

	// Test GetAllBlog (should return 200 OK)
	req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
