package usecase

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    "codein/models"
    "codein/repository"
)

type MockRepository struct {
    mock.Mock
}

func (m *MockRepository) CreateBlog(ctx context.Context, blog models.Blog) (*models.Blog, error) {
    args := m.Called(ctx, blog)
    return args.Get(0).(*models.Blog), args.Error(1)
}

func TestCreateBlogUsecase(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := NewMockRepository(ctrl)

    expectedBlog := models.Blog{
        Title:   "Test Blog",
        Content: "This is a test blog",
    }

    mockRepo.EXPECT().CreateBlog(context.Background(), models.Blog{
        Title:   "Test Blog",
        Content: "This is a test blog",
    }).Return(&expectedBlog, nil)

    usecase := NewUsecase(mockRepo)

    blog, err := usecase.CreateBlog(context.Background(), models.Blog{
        Title:   "Test Blog",
        Content: "This is a test blog",
    })

    assert.NoError(t, err)
    assert.Equal(t, expectedBlog, *blog)
    mockRepo.AssertExpectations(t)
}