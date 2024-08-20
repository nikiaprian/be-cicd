package handler

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

// handler/blog_test.go

// MockProject telah didefinisikan
func TestGetAllBlog(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Inisialisasi Mock
    mockProject := &MockProject{
        Usecase: &MockUsecase{},
        Storage: &MockStorage{}, // Inisialisasi dengan MockStorage
    }

    // Inisialisasi handler
    blogHandler := NewHandler(mockProject) // Pastikan mockProject sesuai tipe yang diterima

    // Panggil fungsi
    blogHandler.GetAllBlog(c)

    // Assert respons
    assert.Equal(t, http.StatusOK, w.Code)
}

// Begitu juga untuk TestCreateBlog
