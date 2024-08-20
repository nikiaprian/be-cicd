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

func TestCreateThread(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := NewMockRepository(ctrl)

    expectedThread := &models.ForumThread{
        Title:   "Test Thread",
        Content: "This is a test thread content",
        UserID:  1,
    }

    mockRepo.EXPECT().CreateThread(context.Background(), expectedThread).Return(expectedThread, nil)

    usecase := usecase.NewUsecase(mockRepo)

    thread, err := usecase.CreateThread(context.Background(), models.ForumThreadRequest{
        Title:   "Test Thread",
        Content: "This is a test thread content",
        UserID:  1,
    })

    assert.Nil(t, err)
    assert.Equal(t, expectedThread, thread)

    mockRepo.AssertExpectations(t)
}

func TestGetThreads(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := NewMockRepository(ctrl)

    expectedThreads := []*models.ForumThread{
        {ID: 1, Title: "Thread 1"},
        {ID: 2, Title: "Thread 2"},
    }

    mockRepo.EXPECT().GetThreads(context.Background()).Return(expectedThreads, nil)

    usecase := usecase.NewUsecase(mockRepo)

    threads, err := usecase.GetThreads(context.Background())

    assert.Nil(t, err)
    assert.Equal(t, expectedThreads, threads)

    mockRepo.AssertExpectations(t)
}