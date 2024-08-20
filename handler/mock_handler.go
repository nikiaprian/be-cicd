package handler

import (
    "github.com/gin-gonic/gin"
    "codein/models"
    "codein/project"
)

// MockUsecase untuk menggantikan Usecase asli
type MockUsecase struct{}

func (m *MockUsecase) GetAllBlog(c *gin.Context) ([]models.Blog, error) {
    return []models.Blog{
        {ID: 1, Title: "Test Blog 1"},
        {ID: 2, Title: "Test Blog 2"},
    }, nil
}

// MockProject untuk menggantikan Project asli
type MockProject struct {
    Usecase *MockUsecase
    Storage StorageInterface // Pastikan ini sesuai dengan definisi StorageInterface Anda
}

// Jika Anda belum mendefinisikan StorageInterface, lakukan seperti ini:
type StorageInterface interface {
    // Tambahkan metode yang diperlukan untuk storage
}
