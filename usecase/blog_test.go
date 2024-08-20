package usecase

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"codein/models"
	"codein/repository"
)

func TestGetAllBlog(t *testing.T) {
	mockRepo := repository.NewRepository(nil) // Assume nil or mock DB for simplicity
	uc := &Usecase{Repository: mockRepo}
	blogs, err := uc.GetAllBlog(nil)

	assert.Nil(t, err)
	assert.NotEmpty(t, blogs)
}
