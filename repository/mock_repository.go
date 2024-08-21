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

// Metode untuk BlogCommentLike
func (m *MockRepository) CreateLikeByBlogCommentId(ctx *gin.Context, blogCommentId int, userId int) (*models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, blogCommentId, userId)
	return args.Get(0).(*models.BlogCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByBlogCommentID(ctx *gin.Context, blogCommentId int) (*[]models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, blogCommentId)
	return args.Get(0).(*[]models.BlogCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByBlogCommentId(ctx *gin.Context, blogCommentId int, userId int) (*models.BlogCommentLikeResponse, error) {
	args := m.Called(ctx, blogCommentId, userId)
	return args.Get(0).(*models.BlogCommentLikeResponse), args.Error(1)
}

// Metode untuk ForumCommentLike
func (m *MockRepository) CreateLikeByForumCommentId(ctx *gin.Context, forumCommentId int, userId int) (*models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, forumCommentId, userId)
	return args.Get(0).(*models.ForumCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) GetAllLikeByForumCommentID(ctx *gin.Context, forumCommentId int) (*[]models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, forumCommentId)
	return args.Get(0).(*[]models.ForumCommentLikeResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByForumCommentId(ctx *gin.Context, forumCommentId int, userId int) (*models.ForumCommentLikeResponse, error) {
	args := m.Called(ctx, forumCommentId, userId)
	return args.Get(0).(*models.ForumCommentLikeResponse), args.Error(1)
}