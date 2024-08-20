package repository

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"codein/models"
)

func TestCreateBlog(t *testing.T) {
	repo := NewRepository(nil) // Assume nil or mock DB for simplicity

	// Simulate a context and blog request
	ctx := context.TODO()
	blogReq := models.BlogRequest{
		Title: "Test Blog",
		Content: "This is a test blog content.",
	}
	userID := 1 // Assume an example user ID
	blogID := 0 // Assume an example blog ID

	err := repo.CreateBlog(ctx, blogReq, "Test Photo", userID)

	assert.Nil(t, err)
	assert.Equal(t, "Test Blog", blogReq.Title)
}
