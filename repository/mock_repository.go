package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"codein/models"
)

// MockRepository adalah struct untuk mock objek repository yang digunakan dalam pengujian
type MockRepository struct {
	mock.Mock
}

// Metode untuk mock pada repository/bloglike.go

func (m *MockRepository) CreateLikeByBlogId(ctx *gin.Context, user_id, blog_id int) (*models.BlogsLikesResponse, error) {
	args := m.Called(ctx, user_id, blog_id)
	return args.Get(0).(*models.BlogsLikesResponse), args.Error(1)
}

// Tambahkan method lainnya sesuai kebutuhan pengujian untuk bloglike.go

// Metode untuk mock pada repository/blogcommentlike.go

func (m *MockRepository) CreateLikeByBlogCommentId(ctx *gin.Context, user_id, blog_comment_id int) (*models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, user_id, blog_comment_id)
	return args.Get(0).(*models.BlogCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByBlogCommentID(ctx *gin.Context, blog_comment_id int) (*[]models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, blog_comment_id)
	return args.Get(0).(*[]models.BlogCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByBlogCommentId(ctx *gin.Context, user_id, blog_comment_id int) (*models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, user_id, blog_comment_id)
	return args.Get(0).(*models.BlogCommentLikeResponse), args.Error(1)
}

// Metode untuk mock pada repository/forumcommentlike.go

func (m *MockRepository) CreateLikeByForumCommentId(ctx *gin.Context, user_id, forum_comment_id int) (*models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, user_id, forum_comment_id)
	return args.Get(0).(*models.ForumCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByForumCommentID(ctx *gin.Context, forum_comment_id int) (*[]models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, forum_comment_id)
	return args.Get(0).(*[]models.ForumCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByForumCommentId(ctx *gin.Context, user_id, forum_comment_id int) (*models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, user_id, forum_comment_id)
	return args.Get(0).(*models.ForumCommentLikeResponse), args.Error(1)
}

// Metode untuk mock pada repository/forumlikes.go

func (m *MockRepository) CreateLikeByForumId(ctx *gin.Context, user_id, forum_id int) (*models.ForumLikesResponse, error) {
	args := m.Called(ctx, user_id, forum_id)
	return args.Get(0).(*models.ForumLikesResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByForumID(ctx *gin.Context, forum_id int) (*[]models.ForumLikesResponse, error) {
	args := m.Called(ctx, forum_id)
	return args.Get(0).(*[]models.ForumLikesResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByForumId(ctx *gin.Context, user_id, forum_id int) (*models.ForumLikesResponse, error) {
	args := m.Called(ctx, user_id, forum_id)
	return args.Get(0).(*models.ForumLikesResponse), args.Error(1)
}
