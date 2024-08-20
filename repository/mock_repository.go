package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"codein/models"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateBlog(ctx *gin.Context, blogReq models.BlogRequest, photo string, userID int) (*models.Blog, error) {
	args := m.Called(ctx, blogReq, photo, userID)
	return args.Get(0).(*models.Blog), args.Error(1)
}

func (m *MockRepository) CreateLikeByBlogId(ctx *gin.Context, user_id, blog_id int) (*models.BlogsLikesResponse, error) {
	args := m.Called(ctx, user_id, blog_id)
	return args.Get(0).(*models.BlogsLikesResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByBlogID(ctx *gin.Context, id int) (*[]models.BlogsLikesResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*[]models.BlogsLikesResponse), args.Error(1)
}

func (m *MockRepository) GetLikeByUserIDAndBlogID(ctx *gin.Context, user_id, blog_id int) (*models.BlogsLikesResponse, error) {
	args := m.Called(ctx, user_id, blog_id)
	return args.Get(0).(*models.BlogsLikesResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByBlogId(ctx *gin.Context, user_id, blog_id int) (*models.BlogsLikesResponse, error) {
	args := m.Called(ctx, user_id, blog_id)
	return args.Get(0).(*models.BlogsLikesResponse), args.Error(1)
}

// Tambahkan method lainnya sesuai kebutuhan pengujian
