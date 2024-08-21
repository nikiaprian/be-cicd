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

// BlogCommentLike related methods
func (m *MockRepository) CreateLikeByBlogCommentId(like *models.BlogCommentLikes) error {
	args := m.Called(like)
	return args.Error(0)
}

func (m *MockRepository) GetAllLikeByBlogCommentID(blogCommentID int) ([]models.BlogCommentLikesResponse, error) {
	args := m.Called(blogCommentID)
	return args.Get(0).([]models.BlogCommentLikesResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByBlogCommentId(blogCommentID int, userID int) error {
	args := m.Called(blogCommentID, userID)
	return args.Error(0)
}

// ForumCommentLike related methods
func (m *MockRepository) CreateLikeByForumCommentId(like *models.ForumCommentLikes) error {
	args := m.Called(like)
	return args.Error(0)
}

func (m *MockRepository) GetAllLikeByForumCommentID(forumCommentID int) ([]models.ForumCommentLikesResponse, error) {
	args := m.Called(forumCommentID)
	return args.Get(0).([]models.ForumCommentLikesResponse), args.Error(1)
}

func (m *MockRepository) DeleteLikeByForumCommentId(forumCommentID int, userID int) error {
	args := m.Called(forumCommentID, userID)
	return args.Error(0)
}

// // untukforumlike
// func (m *MockRepository) CreateLikeByForumId(like models.ForumsLikes) error {
//     args := m.Called(like)
//     return args.Error(0)
// }

// func (m *MockRepository) GetAllLikeByForumID(forumID int64) ([]models.ForumsLikesResponse, error) {
//     args := m.Called(forumID)
//     return args.Get(0).([]models.ForumsLikesResponse), args.Error(1)
// }

// func (m *MockRepository) DeleteLikeByForumId(likeID int64) error {
//     args := m.Called(likeID)
//     return args.Error(0)
// }