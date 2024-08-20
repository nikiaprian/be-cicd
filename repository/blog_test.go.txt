package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"codein/models"
)

func TestCreateBlog(t *testing.T) {
	repo := NewRepository(nil) // Assume nil or mock DB for simplicity
	blog := models.Blog{Title: "Test Blog"}
	err := repo.CreateBlog(&blog)

	assert.Nil(t, err)
	assert.Equal(t, "Test Blog", blog.Title)
}
