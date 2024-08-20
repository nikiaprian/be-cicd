package handler

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestGetAllBlog(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Inisialisasi Mock
    mockProject := &MockProject{
        Usecase: &MockUsecase{},
    }

    // Inisialisasi handler
    blogHandler := NewHandler(mockProject)

    // Panggil fungsi
    blogHandler.GetAllBlog(c)

    // Assert respons
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBlog(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Inisialisasi Mock
    mockProject := &MockProject{
        Usecase: &MockUsecase{},
    }

    // Inisialisasi handler
    blogHandler := NewHandler(mockProject)

    // Simulasi input
    c.Request, _ = http.NewRequest(http.MethodPost, "/blogs", nil)

    // Panggil fungsi
    blogHandler.CreateBlog(c)

    // Assert respons
    assert.Equal(t, http.StatusCreated, w.Code)
}
