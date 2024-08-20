package handler

import (
    "github.com/gin-gonic/gin"
)

// MockUsecase untuk menggantikan Usecase asli
type MockUsecase struct{}

func (m *MockUsecase) GetAllBlog(c *gin.Context) ([]models.Blog, error) {
    return []models.Blog{
        {ID: 1, Title: "Test Blog 1"},
        {ID: 2, Title: "Test Blog 2"},
    }, nil
}

// MockStorage untuk menggantikan Storage asli
type MockStorage struct{}

// Implementasikan metode yang diperlukan untuk StorageInterface
func (m *MockStorage) SomeMethod() {
    // Implementasi metode sesuai kebutuhan
}

// MockProject untuk menggantikan Project asli
type MockProject struct {
    Usecase *MockUsecase
    Storage *MockStorage // Menggunakan MockStorage yang sesuai
}

// Pastikan MockProject memiliki semua metode yang diperlukan
