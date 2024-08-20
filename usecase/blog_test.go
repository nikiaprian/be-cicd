package usecase_test

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    "codein/models"
    "codein/repository"
    "codein/tests"
    "codein/usecase"
)

func TestCreateComment(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := NewMockRepository(ctrl)

    expectedComment := &models.Comment{
        Content: "This is a test comment",
        UserID:  1,
        BlogID:   2,
    }

    mockRepo.EXPECT().CreateComment(context.Background(), expectedComment).Return(expectedComment, nil)

    usecase := usecase.NewUsecase(mockRepo)

    comment, err := usecase.CreateComment(context.Background(), models.CommentRequest{
        Content: "This is a test comment",
        UserID:  1,
        BlogID:   2,
    })

    assert.Nil(t, err)
    assert.Equal(t, expectedComment, comment)

    mockRepo.AssertExpectations(t)
}

func TestGetCommentsByBlogID(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := NewMockRepository(ctrl)

    expectedComments := []*models.Comment{
        {ID: 1, Content: "Comment 1"},
        {ID: 2, Content: "Comment 2"},
    }

    mockRepo.EXPECT().GetCommentsByBlogID(context.Background(), 1).Return(expectedComments, nil)

    usecase := usecase.NewUsecase(mockRepo)

    comments, err := usecase.GetCommentsByBlogID(context.Background(), 1)

    assert.Nil(t, err)
    assert.Equal(t, expectedComments, comments)

    mockRepo.AssertExpectations(t)
}