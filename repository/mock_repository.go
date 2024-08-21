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

func (m *MockRepository) CreateLikeByBlogCommentId(ctx *gin.Context, userID, commentID int) (*models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, userID, commentID)
	return args.Get(0).(*models.BlogCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByBlogCommentID(ctx *gin.Context, commentID int) (*[]models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, commentID)
	return args.Get(0).(*[]models.BlogCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByBlogCommentId(ctx *gin.Context, userID, commentID int) (*models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, userID, commentID)
	return args.Get(0).(*models.BlogCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) CreateLikeByForumCommentId(ctx *gin.Context, userID, commentID int) (*models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, userID, commentID)
	return args.Get(0).(*models.ForumCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByForumCommentID(ctx *gin.Context, commentID int) (*[]models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, commentID)
	return args.Get(0).(*[]models.ForumCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByForumCommentId(ctx *gin.Context, userID, commentID int) (*models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, userID, commentID)
	return args.Get(0).(*models.ForumCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) CreateLikeByForumId(ctx *gin.Context, userID, forumID int) (*models.ForumLikesResponse, error) {
	args := m.Called(ctx, userID, forumID)
	return args.Get(0).(*models.ForumLikesResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByForumID(ctx *gin.Context, forumID int) (*[]models.ForumLikesResponse, error) {
	args := m.Called(ctx, forumID)
	return args.Get(0).(*[]models.ForumLikesResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByForumId(ctx *gin.Context, userID, forumID int) (*models.ForumLikesResponse, error) {
	args := m.Called(ctx, userID, forumID)
	return args.Get(0).(*models.ForumLikesResponse), args.Error(1)
}

// Tambahkan method lainnya sesuai kebutuhan pengujian
