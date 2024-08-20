package repository

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"codein/models"
)

func TestCreateBlog(t *testing.T) {
	// Inisialisasi repository (anggap db nil atau mock DB untuk kesederhanaan)
	repo := NewRepository(nil)

	// Simulasi gin context dan blog request
	ctx, _ := gin.CreateTestContext(nil)
	blogReq := models.BlogRequest{
		Title:   "Test Blog",
		Content: "This is a test blog content.",
	}
	userID := 1 // Asumsikan user ID contoh

	// Panggil fungsi CreateBlog
	blog, err := repo.CreateBlog(ctx, blogReq, "Test Photo", userID)

	// Verifikasi hasil
	assert.Nil(t, err)
	assert.NotNil(t, blog)
	assert.Equal(t, "Test Blog", blog.Title)
}
