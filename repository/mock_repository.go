package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/gin-gonic/gin"
	"codein/models"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateBlog(ctx *gin.Context, blog models.Blog) (*models.Blog, error) {
	args := m.Called(ctx, blog)
	return args.Get(0).(*models.Blog), args.Error(1)
}

// Tambahkan method lainnya sesuai kebutuhan pengujian
