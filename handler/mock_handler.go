package handler

import (
    "codein/models"
    "codein/project"
)

type MockUsecase struct{}

func (m *MockUsecase) GetAllBlog(c *gin.Context) ([]models.Blog, error) {
    return []models.Blog{
        {ID: 1, Title: "Test Blog 1"},
        {ID: 2, Title: "Test Blog 2"},
    }, nil
}

// Implementasikan semua metode lain yang dibutuhkan
type MockProject struct {
    Usecase *MockUsecase
    Storage StorageInterface // Ganti ini jika Anda memiliki StorageInterface
}
